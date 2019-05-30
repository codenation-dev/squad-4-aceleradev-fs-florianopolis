package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/mholt/archiver"
)

type client struct {
	name         string
	isBankFunc   bool
	isPublicFunc bool
	income       float64
}

var clients []client

func main() {

	// TODO: fazer a base num BD SQL para começar a brincar
	// createDB()
	s, b := os.LookupEnv("DB_HOST")
	fmt.Println(s, b)

}

func check(command string, e error) {
	if e != nil {
		log.Fatalf("%q failed with %s\n", command, e.Error())
	}

}

func importClientsCSV() []client {
	// import from the file given by the exercise
	// returns []uf

	f, err := os.Open("clientes.csv")
	check("os.Open", err)
	defer f.Close()

	// Parse the file
	r := csv.NewReader(f)

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check("r.Read()", err)

		clients = append(clients, client{name: record[0]})
	}
	return clients
}

func downloadHTTPFile(path, filename string) (string, error) {
	// download the file with the list of SP public agents
	// returns the path to the file
	r, err := http.Get(path + filename + ".rar")
	check("http.Get", err)

	destPath := "file.rar"
	dest, err := os.Create(destPath)
	check("os.Create", err)
	defer dest.Close()

	_, err = io.Copy(dest, r.Body)
	check("io.Copy", err)

	return destPath, nil
}

func fetchPublicAgentsFile() {
	filename := "remuneracao_Marco_2019"
	path := "http://www.transparencia.sp.gov.br/PortalTransparencia-Report/historico/"
	destFolder := "extracted"
	compressedFile := "file.rar"

	if _, err := os.Stat(
		"/home/gui/dev/go_projects/src/codenation/banco-uati-presencial/file.rar",
	); err != nil {
		_, err := downloadHTTPFile(path, filename)
		check("downloadPublicAgentsFile", err)
	}

	if _, err := os.Stat(
		"/home/gui/dev/go_projects/src/codenation/banco-uati-presencial/extracted/Remuneracao_Marco_2019.txt",
	); err != nil {
		err := archiver.Unarchive(compressedFile, destFolder)
		check("archiver.Unarchive", err)
	}
}

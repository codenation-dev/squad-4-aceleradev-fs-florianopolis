package importing

import (
	"bufio"
	"codenation/squad-4-aceleradev-fs-florianopolis/pkg/entity"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/mholt/archiver"
)

func readCSV(path string, job func([]string) bool, sep rune, hasHeader bool) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open path %v: %v", path, err)
	}
	fmt.Println("CSV opened")
	defer f.Close()

	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = sep
	if hasHeader {
		r.Read()
	}
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Could not read csv file: %v", err)
		}
		fmt.Println("Reading CSV...")

		keepgoing := job(line)
		if keepgoing != true {
			break
		}
	}
}

// ImportClientesCSV imports the data from clientes.csv file
func ImportClientesCSV(path string) ([]entity.Customer, error) {
	var customers []entity.Customer

	job := func(row []string) bool {
		customers = append(customers, entity.Customer{
			Name: row[0],
		})
		return true
	}
	readCSV(path, job, ',', true)
	return customers, nil
}

// DownloadHTTPFile downloads the file with the list of SP public agents
// and returns the path to the downloaded file
func DownloadHTTPFile(path, filename string) (string, error) {
	r, err := http.Get(path + filename + ".rar")
	if err != nil {
		return "", err
	}
	destPath := "file.rar"
	dest, err := os.Create(destPath)
	if err != nil {
		return "", err
	}
	defer dest.Close()

	_, err = io.Copy(dest, r.Body)
	if err != nil {
		return "", err
	}

	return destPath, nil
}

// ImportPublicFunc import from the goverment site
func ImportPublicFunc() ([]entity.PublicFunc, error) {
	var indexName = 1
	var indexIncome = 3

	if _, err := os.Stat("file.rar"); os.IsNotExist(err) {
		err := fetchPublicAgentsFile()
		if err != nil {
			log.Fatal(err)
		}
	}

	publicFuncs := []entity.PublicFunc{}

	job := func(row []string) bool {
		n := row[indexName]
		ws := strings.Replace(row[indexIncome], ",", ".", 1)
		wf, err := strconv.ParseFloat(ws, 32)
		if err != nil {
			log.Fatal(err)
		}
		publicFuncs = append(publicFuncs, entity.PublicFunc{
			Name:  n,
			Wage:  float32(wf),
			Place: "São Paulo",
		})
		return true
	}

	readCSV("Remuneracao_Abril_2019.txt", job, ';', true)
	return publicFuncs, nil
}

func fetchPublicAgentsFile() error {
	filename := "Remuneracao_Abril_2019"
	path := "http://www.transparencia.sp.gov.br/PortalTransparencia-Report/historico/"
	destFolder := "."
	compressedFile := "file.rar"

	if _, err := os.Stat("file.rar"); err != nil {
		_, err := DownloadHTTPFile(path, filename)
		if err != nil {
			return err
		}
	}

	if _, err := os.Stat(filename + ".txt"); err != nil {
		err := archiver.Unarchive(compressedFile, destFolder)
		if err != nil {
			return err
		}
	}
	return nil
}

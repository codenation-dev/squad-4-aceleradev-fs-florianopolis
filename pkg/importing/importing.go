// Package importing implementa funções necessárias à importação dos
// dados tanto de clientes quanto de funcionários públicos.
package importing

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
	counter := 0
	for {
		counter++
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Could not read csv file: %v", err)
		}
		if counter%300000 == 0 {
			fmt.Println("Reading CSV...")
		}
		keepgoing := job(line)
		if keepgoing != true {
			break
		}
	}
}

// DownloadHTTPFile downloads the file with the list of SP public agents
// and returns the path to the downloaded file
func downloadHTTPFile(from, to string) error {
	r, err := http.Get(from)
	if err != nil {
		return err
	}
	dest, err := os.Create(to)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, r.Body)
	if err != nil {
		return err
	}

	return nil
}

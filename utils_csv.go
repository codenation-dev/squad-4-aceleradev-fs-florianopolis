package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
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

func importClientesCSV(path string) error {
	var accountHolders []accountHolder
	var errors []string

	job := func(row []string) bool {
		accountHolders = append(accountHolders, accountHolder{
			Name: row[0],
		})
		return true
	}
	readCSV(path, job, ',', true)

	for _, ah := range accountHolders {
		err := addAccountHolder(ah)
		if err != nil {
			log.Fatalf("could not add to db: %v", err)
		}
	}

	if errors != nil {
		return fmt.Errorf("Could not add: %v", errors)
	}
	return nil
}

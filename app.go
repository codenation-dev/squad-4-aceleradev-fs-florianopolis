package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/mholt/archiver"
)

const (
	extractFolder  = "extracted"
	downloadFolder = "downloaded"
)

func downloadHTTPFile(httpPath, localPath string) error {
	// download the file with the list of SP public agents

	r, err := http.Get(httpPath)
	dest, err := os.Create(localPath)
	defer dest.Close()
	_, err = io.Copy(dest, r.Body)
	if err != nil {
		return fmt.Errorf("could not download HTTPfile: %v", err)
	}
	return nil
}

func createFolderIfDoesntExist(expressions ...string) error {
	for _, expr := range expressions {
		if _, err := os.Stat(expr); os.IsNotExist(err) {
			if err := os.Mkdir(expr, 0777); err != nil {
				return fmt.Errorf("could not createFolderIfNotExist: %v", err)
			}
		}
	}
	return nil
}

func importSPFile() error {
	fileToImport := "Remuneracao_Marco_2019"
	pathToTextFile, err := downloadSPPublicFile(fileToImport, false)
	if err != nil {
		return fmt.Errorf("could not downloadSPPublicFile(): %v", err)
	}
	var employees []publicEmployee
	var indexName = 0
	var indexWage = 3

	job := func(row []string) bool {

		s := strings.Replace(row[indexWage], ",", ".", 1)
		f, err := strconv.ParseFloat(s, 32)
		if err != nil {
			panic(err)
		}
		employees = append(employees, publicEmployee{
			Name:  row[indexName],
			Wage:  float32(f),
			Local: "Estado de SP",
		})
		return true
	}

	err = readCSV(pathToTextFile, job, ';', true)
	if err != nil {
		return fmt.Errorf("could not readCSV(): %v", err)
	}

	for i, e := range employees {
		go createEmployee(e)
		fmt.Println(i, e.Name)
	}

	return nil
}

func downloadSPPublicFile(file string, forceDownload bool) (string, error) {
	fileToDownload := fmt.Sprintf("%s.rar", file)
	httpAddressSP := "http://www.transparencia.sp.gov.br/PortalTransparencia-Report/historico"

	createFolderIfDoesntExist(downloadFolder, extractFolder)

	downloadFrom := fmt.Sprintf("%s/%s", httpAddressSP, fileToDownload)
	downloadTo := fmt.Sprintf("%v/%v", downloadFolder, fileToDownload)
	_, err := os.Stat(downloadTo)
	if os.IsNotExist(err) {
		err = downloadHTTPFile(downloadFrom, downloadTo)
		if err != nil {
			return "", fmt.Errorf("could not download from %s: %v", downloadFrom, err)
		}
	} else if forceDownload {
		err := os.Remove(downloadTo)
		if err != nil {
			return "", fmt.Errorf("could not remove %s: %v", downloadTo, err)
		}
		err = downloadHTTPFile(downloadFrom, downloadTo)
		if err != nil {
			return "", fmt.Errorf("could not download from %s: %v", downloadFrom, err)
		}
	}

	extractTo := fmt.Sprintf("%s/%s.txt", extractFolder, file)
	_, err = os.Stat(extractTo)
	if os.IsNotExist(err) {
		err := archiver.Unarchive(downloadTo, extractFolder)
		if err != nil {
			return "", fmt.Errorf("could not extract %s: %v", extractTo, err)
		}
	} else if forceDownload {
		err := os.Remove(extractTo)
		if err != nil {
			return "", fmt.Errorf("could not remove %s: %v", extractTo, err)
		}
		err = archiver.Unarchive(downloadTo, extractFolder)
		if err != nil {
			return "", fmt.Errorf("could not extract %s: %v", extractTo, err)
		}
	}
	return extractTo, nil
}

func readCSV(path string, job func([]string) bool, sep rune, hasHeader bool) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Could not open path %v: %v", path, err)
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
			return fmt.Errorf("Could not read csv file: %v", err)
		}
		fmt.Println("Reading CSV...")

		keepgoing := job(line)
		if keepgoing != true {
			break
		}
	}
	return nil
}

func importClientesCSV(path string) error {
	var customers []customer

	job := func(row []string) bool {
		customers = append(customers, customer{
			Name:        row[0],
			Wage:        0.00,
			IsPublic:    0,
			SentWarning: "",
		})
		return true
	}
	readCSV(path, job, ',', true)

	for _, c := range customers {
		// ASK: É melhor abrir e fechar o DB aqui, abrindo somente uma vez,
		// acrescentando todos os customers de uma vez, ou é melhor (como estou fazendo),
		// deixar o openDB() debtro da função addCustomer?
		_, err := createCustomer(c)
		if err != nil {
			return fmt.Errorf("could not add to db: %v", err)
		}
	}
	return nil
}

// func addNewUser() {
// 	var u user
// 	u.Email = readStringInput("Digite o email do usuário: ")
// 	u.Password = readStringInput("Digite a senha: ")
// 	err := addUser(u)
// 	if err != nil {
// 		log.Printf("não foi possivel cadastrar usuário: %v", err)
// 	}
// 	again := readStringInput("\nCadastrar mais um? (s/n) ")
// 	if again == "s" {
// 		addNewUser()
// 	}
// }

// func makeMapAgents() map[string]string {
// 	var indexName int
// 	var indexIncome = 3

// 	if _, err := os.Stat(extractedFilePath); os.IsNotExist(err) {
// 		downloadSPPublicFile(false)
// 	}

// 	mapAgents := make(map[string]string)
// 	job := func(row []string) bool {
// 		r := strings.Replace(row[indexIncome], ",", ".", 1)
// 		mapAgents[row[indexName]] = r
// 		return true
// 	}

// 	readCSV(extractedFilePath, job, ';', true)
// 	return mapAgents
// }

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var originalFilename = "remuneracao_Marco_2019.rar"
var savedFilename = "Remuneracao_Marco_2019.txt"
var httpAddress = "http://www.transparencia.sp.gov.br/PortalTransparencia-Report/historico"
var extractFolder = "extracted"
var downloadFolder = "downloaded"
var downloadedFilePath = fmt.Sprintf("%v/%v", downloadFolder, originalFilename)
var httpPath = fmt.Sprintf("%v/%v", httpAddress, originalFilename)
var extractedFilePath = fmt.Sprintf("%v/%v", extractFolder, savedFilename)
var dbaccountHoldersPath = "db/accountHolders.csv"

const indexaccountHolderName int = 0
const indexisBankFunc int = 1
const indexisPublicFunc int = 2
const indexIncome int = 3

func main() {
	menuPrinc()
}

func menuPrinc() {

	fmt.Println("\nMenu inicial:\n*************")
	fmt.Println("S-\tSetup novo banco de dados")
	fmt.Println("I-\tImportar arquivo de clientes")
	fmt.Println("A-\tAtualizar, comparando com lista dos funcionários públicos")
	fmt.Println("C-\tCadastrar usuários para receber alertas")
	fmt.Println("L-\tListar todos os clientes")
	fmt.Println("U-\tListar usuários")
	fmt.Println("D-\tDashboard")

	choice := readStringInput("\nDigite a opção desejada: ")

	if choice == "I" {
		err := importClientesCSV("clientes.csv")
		if err != nil {
			log.Fatal(err)
		}
		menuPrinc()
	} else if choice == "U" {
		err := readUsers()
		if err != nil {
			fmt.Printf("Could not read db: %v", err)
		}
		menuPrinc()
	} else if choice == "L" {
		err := readAccountHolders()
		if err != nil {
			fmt.Printf("Could not read db: %v", err)
		}
		menuPrinc()
	} else if choice == "S" {
		err := setupDB()
		if err != nil {
			log.Printf("Could not setup a new db: %v", err)
		}
		menuPrinc()
	} else if choice == "C" {
		addNewUser()
		menuPrinc()
	} else {
		fmt.Println("\nFudeu!!!")
	}

}

func addNewUser() {
	var u user
	u.Email = readStringInput("Digite o email do usuário: ")
	u.Password = readStringInput("Digite a senha: ")
	err := addUser(u)
	if err != nil {
		log.Printf("não foi possivel cadastrar usuário: %v", err)
	}
	again := readStringInput("\nCadastrar mais um? (s/n) ")
	if again == "s" {
		addNewUser()
	}
}

func makeMapAgents() map[string]string {
	var indexName int
	var indexIncome = 3

	if _, err := os.Stat(extractedFilePath); os.IsNotExist(err) {
		downloadPublicAgentsFile(false)
	}

	mapAgents := make(map[string]string)
	job := func(row []string) bool {
		r := strings.Replace(row[indexIncome], ",", ".", 1)
		mapAgents[row[indexName]] = r
		return true
	}

	readCSV(extractedFilePath, job, ';', true)
	return mapAgents
}

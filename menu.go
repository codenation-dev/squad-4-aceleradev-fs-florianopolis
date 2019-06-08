package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func mainMenu() {

	fmt.Println("\nMenu inicial:\n*************")
	fmt.Println("S-\tSetup novo banco de dados")
	fmt.Println("T-\tListar todas as tabelas")
	fmt.Println("L-\tListar todos os clientes")
	fmt.Println("U-\tListar usuários")
	fmt.Println("I-\tImportar arquivo de clientes")
	fmt.Println("SP-\tImportar arquivo de funcionários de SP")
	fmt.Println()

	fmt.Println("A-\tAtualizar, comparando com lista dos funcionários públicos")
	fmt.Println("C-\tCadastrar usuários para receber alertas")
	fmt.Println("D-\tDashboard")

	choice, err := readStringInput("\nDigite a opção desejada: ")
	if err != nil {
		log.Fatal(err)
	}

	if choice == "I" || choice == "i" {
		err := importClientesCSV("clientes.csv")
		if err != nil {
			log.Fatal(err)
		}
		mainMenu()
	} else if choice == "T" || choice == "t" {
		t, err := listTables()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(t)
		fmt.Scanln()
		mainMenu()
	} else if choice == "SP" {
		err := importSPFile()
		if err != nil {
			log.Fatal(err)
		}
		mainMenu()
	} else if choice == "L" {
		// customers, err := readAllCustomers()
		// if err != nil {
		// 	fmt.Printf("Could not readAllCustomers: %v", err)
		// }
		// for _, c := range customers {
		// 	fmt.Println(c)
		// }
		mainMenu()
	} else if choice == "S" {
		err := createAllTables()
		if err != nil {
			log.Printf("Could not setup a new db: %v", err)
		}
		mainMenu()
	} else if choice == "C" {
		// addNewUser()
		mainMenu()
	} else {
		fmt.Println("\nFudeu!!!")
	}

}
func readStringInput(message string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	var choice string
	fmt.Print("\n", message)
	choice, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("could not readStringInput: %v", err)
	}
	return strings.Replace(choice, "\n", "", -1), nil
}

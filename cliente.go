package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type cliente struct {
	Name string
}

const (
	host     = "127.17.0.2"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "uati"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	clienteList := readCsvFile("clientes.csv")
	//fmt.Println(clienteList)
	//	insertClientList(db, clienteList)
	for id, cliente := range clienteList {
		println(fmt.Sprintf("INSERT INTO cliente(id_cliente, nome) VALUES(%d, '%s')", id+1, cliente.Name))
		rows, err := db.Query(fmt.Sprintf("INSERT INTO cliente(id_cliente, nome) VALUES(%d, '%s')", id+1, cliente.Name))
		if err != nil {
			panic(err)
		} else {
			fmt.Println(rows)
		}
	}
	defer db.Close()

}

/*
func insertClientList(db *DB, clienteList []cliente) {
	for id, cliente := range clienteList {
		rows, err := db.Query(fmt.Sprintf("INSERT INTO cliente(id, nome) VALUES(%d, '%s')", id, cliente.nome))
	}
}*/

func readCsvFile(filename string) []cliente {
	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	line, error := reader.Read()
	var clienteList []cliente
	for {
		line, error = reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		clienteList = append(clienteList, cliente{
			Name: line[0],
		})
	}
	defer csvFile.Close()
	//fmt.Println(jogadorList[:20])
	return clienteList
}

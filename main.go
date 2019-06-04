package main

import (
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "12345"
	DB_NAME     = "postgres"
)

type customer struct {
	ID          int32   `json:"id"`
	Name        string  `json:"name"`
	Wage        float32 `json:"wage"`
	IsPublic    int8    `json:"isPublic"`
	SentWarning string  `json:"sentWarning"` // todo: colocar timestamp
}

type publicEmployee struct {
	ID    int32   `json:"id"`
	Name  string  `json:"name"`
	Wage  float32 `json:"wage"`
	Local string  `json:"local"`
}

func main() {
	// mainMenu()
	// r, err := deleteAllRows("public_employees")
	// fmt.Println(r, err)
	// fmt.Println(importSPFile())
	// db := openDB()
	// defer db.Close()

	// createAllTables(db)
	// fmt.Println(listTables())
	// fmt.Println(dropTable("customers"))

	// answer, err := readEmployee("WHERE wage > 20000")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, a := range answer {
	// 	fmt.Println(a)
	// }
	// fmt.Println(len(answer))

	var e publicEmployee
	e.read("WHERE name LIKE '%SOUZA%' ORDER BY RANDOM() LIMIT 1")
	fmt.Println(e)

}

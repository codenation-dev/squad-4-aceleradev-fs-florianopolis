package postgres

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/backend/pkg/entity"

	"github.com/DATA-DOG/go-sqlmock"
)

var customerRows = sqlmock.NewRows([]string{
	"id", "name", "wage", "is_public", "sent_warning"}).
	AddRow(1, "test name", 1234.56, 1, "test warning")

// var twoRows = rows.AddRow(2, "test name 2", 123456.78, 0, "")

var mc = entity.Customer{ // mock customer
	ID:          1,
	Name:        "TEST NAME",
	Wage:        1234.56,
	IsPublic:    1,
	SentWarning: "TEST WARNING",
}

var userRows = sqlmock.NewRows([]string{
	"id", "login", "email", "pass"}).
	AddRow(1, "test login", "test@email.com", "1234")

var mu = entity.User{ // mock customer
	ID:    1,
	Login: "TEST login",
	Email: "test@email.com",
	Pass:  "1234",
}

var warningRows = sqlmock.NewRows([]string{
	"id", "dt", "message", "sent_to", "from_customer"}).
	AddRow(1, "test dt", "test message", "id user", "id customer")

var mw = entity.Warning{ // mock warning
	1, "test dt", "test message", "id user", "id customer",
}

var publicRows = sqlmock.NewRows([]string{
	"id", "name", "wage", "place"}).
	AddRow(1, "test name", 1234.56, "test place")

var mp = entity.PublicFunc{ // mock public_func
	1, "test name", 23456.78, "test place",
}

// TODO: Tem como fazer estes testes usando uma tt como essa abaixo?
// // func TestExec(t *testing.T){
// // 	var tt = []struct{
// // 		name string
// // 		query string
// // 		expResult []int64
// // 		funcToCall interface{}
// // 	}{
// // 		{"TestDeleteCustomerById", `DELETE FROM customers WHERE id=\$1`, []int64{1, 1}, interface{func}}
// // 	}

// // 	for _

// // }

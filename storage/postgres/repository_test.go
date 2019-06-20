package postgres

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/entity"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var rows = sqlmock.NewRows([]string{
	"id", "name", "wage", "is_public", "sent_warning"}).
	AddRow(1, "test name", 1234.56, 1, "test warning")
var twoRows = rows.AddRow(2, "test name 2", 123456.78, 0, "")

var mc = entity.Customer{ // mock customer
	ID:          1,
	Name:        "TEST NAME",
	Wage:        1234.56,
	IsPublic:    1,
	SentWarning: "TEST WARNING",
}

func TestDeleteCustomerById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err, fmt.Sprintf("error when opening the mock db connection: %v", err))
	defer db.Close()

	query := `DELETE FROM customers WHERE id=\$1`
	expResult := sqlmock.NewResult(1, 1)
	mock.ExpectExec(query).
		WithArgs(1).WillReturnResult(expResult)

	s := Storage{db}
	err = s.DeleteCustomerByID(1)
	assert.NoError(t, err, err)
}

func TestAddCustomer(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.Nil(t, err, fmt.Sprintf("error when opening the mock db connection: %v", err))
	defer db.Close()

	query := `INSERT INTO customers \(name, wage, is_public, sent_warning\)
	VALUES \(\$1, \$2, \$3, \$4\)`
	expResult := sqlmock.NewResult(1, 1)
	mock.ExpectExec(query).
		WithArgs(mc.Name, mc.Wage, mc.IsPublic, mc.SentWarning).
		WillReturnResult(expResult)

	s := Storage{db}
	res, err := s.AddCustomer(mc)
	assert.NoError(t, err, err)
	expLII, _ := expResult.LastInsertId()
	expRA, _ := expResult.RowsAffected()
	LII, _ := res.LastInsertId()
	RA, _ := res.RowsAffected()
	assert.Equal(t, expLII, LII)
	assert.Equal(t, expRA, RA)
}

func TestGetCustomerByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err, fmt.Sprintf("error when opening the mock db connection: %v", err))
	defer db.Close()

	id := int(1)
	mock.ExpectQuery("SELECT \\* FROM customers WHERE id=\\$1").
		WithArgs(id).
		WillReturnRows(rows)

	s := Storage{db}

	customer, err := s.GetCustomerByID(int(1))
	assert.NoError(t, err)
	assert.NotNil(t, customer)
}

func TestGetAllCustomers(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	s := Storage{db}
	query := `SELECT \* FROM customers`

	mock.ExpectQuery(query).WillReturnRows(twoRows)
	customers, err := s.GetAllCustomers()
	assert.NoError(t, err)
	assert.NotNil(t, customers)
	// assert.Equal(t, twoRows, customers) //TODO: tem como testar o conteudo das rows?
}

func TestGetCustomersByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	s := Storage{db}

	pattern := "test"
	query := `SELECT \* FROM customers WHERE name LIKE`
	mock.ExpectQuery(query).WillReturnRows(twoRows)

	customers, err := s.GetCustomerByName(pattern)
	assert.NoError(t, err)
	assert.NotNil(t, customers) //TODO: incluir teste para quando não retornar colunas
}

func TestUpdateCustomer(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	s := Storage{db}
	query := `UPDATE customers
			SET name=\$1, wage=\$2, is_public=\$3, sent_warning=\$4
			WHERE id=\$5`
	mock.ExpectExec(query).
		WithArgs(mc.Name, mc.Wage, mc.IsPublic, mc.SentWarning, mc.ID).
		WillReturnError(errors.New("could not update the customer"))
	// WillReturnResult(sqlmock.NewResult(1, 1))
	err = s.UpdateCustomer(mc)
	assert.Error(t, err, err)
}

package postgres

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/entity"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCustomerById(t *testing.T) {

}

func TestAddCustomer(t *testing.T) {

	mc := entity.Customer{ // mock customer
		ID:          1,
		Name:        "TEST NAME",
		Wage:        1234.56,
		IsPublic:    1,
		SentWarning: "TEST WARNING",
	}
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

	rows := sqlmock.NewRows([]string{
		"id", "name", "wage", "is_public", "sent_warning"}).
		AddRow(1, "test name", 1234.56, 1, "test warning")

	id := int(1)
	mock.ExpectQuery("SELECT \\* FROM customers WHERE id=\\$1").
		WithArgs(id).
		WillReturnRows(rows)

	s := Storage{db}

	customer, err := s.GetCustomerByID(int(1))
	assert.NoError(t, err)
	assert.NotNil(t, customer)
}

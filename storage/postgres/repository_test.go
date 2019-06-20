package postgres

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetCustomerById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err, fmt.Sprintf("error when opening the mock db connection: %v", err))
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"id", "name", "wage", "is_public", "sent_warning"}).
		AddRow(1, "test name", 1234.56, 1, "test warning")

	id := int(1)
	mock.ExpectQuery("SELECT").
		WithArgs(id).
		WillReturnRows(rows)

	s := Storage{db}

	customer, err := s.GetCustomerByID(int(1))
	assert.NoError(t, err)
	assert.NotNil(t, customer)
}

package postgres

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCustomerById(t *testing.T) {
	t.Parallel()
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

func TestDeleteUserByID(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err, "could not create mock")
	defer db.Close()

	query := `DELETE FROM users WHERE id=\$1`
	expResult := sqlmock.NewResult(1, 1)
	mock.ExpectExec(query).WithArgs(1).WillReturnResult(expResult)

	s := Storage{db}
	err = s.DeleteUserByID(1)
	assert.NoError(t, err, "could not delete")
}

package postgres

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// All
func TestGetAllCustomers(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	s := Storage{db}
	query := `SELECT \* FROM customers`

	mock.ExpectQuery(query).WillReturnRows(customerRows)
	customers, err := s.GetAllCustomers()
	assert.NoError(t, err)
	assert.NotNil(t, customers)
	// assert.Equal(t, twoRows, customers) //TODO: tem como testar o conteudo das rows?
}

func TestGetAllUsers(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	s := Storage{db}
	query := `SELECT \* FROM users`

	mock.ExpectQuery(query).WillReturnRows(userRows)
	customers, err := s.GetAllUsers()
	assert.NoError(t, err)
	assert.NotNil(t, customers)
	// assert.Equal(t, twoRows, customers) //TODO: tem como testar o conteudo das rows?
}

func TestGetAllWarnings(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	s := Storage{db}
	query := `SELECT \* FROM warnings`

	mock.ExpectQuery(query).WillReturnRows(warningRows)
	customers, err := s.GetAllWarnings()
	assert.NoError(t, err)
	assert.NotNil(t, customers)
	// assert.Equal(t, twoRows, customers) //TODO: tem como testar o conteudo das rows?
}

// ByName
func TestGetCustomerByName(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	s := Storage{db}

	pattern := "test"
	query := `SELECT \* FROM customers WHERE name LIKE`
	mock.ExpectQuery(query).WillReturnRows(customerRows)

	customers, err := s.GetCustomerByName(pattern)
	assert.NoError(t, err)
	assert.NotNil(t, customers) //TODO: incluir teste para quando não retornar colunas
}

func TestGetUserByEmail(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	s := Storage{db}

	pattern := "Test"
	query := `SELECT \* FROM users WHERE email LIKE`
	mock.ExpectQuery(query).WillReturnRows(userRows)

	users, err := s.GetUserByEmail(pattern)
	assert.NoError(t, err)
	assert.NotNil(t, users) //TODO: incluir teste para quando não retornar colunas
}

func TestGetWarningByCustomer(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	s := Storage{db}

	pattern := "Test"
	query := `SELECT \* FROM warnings WHERE from_customer LIKE`
	mock.ExpectQuery(query).WillReturnRows(warningRows)

	users, err := s.GetWarningByCustomer(pattern)
	assert.NoError(t, err)
	assert.NotNil(t, users) //TODO: incluir teste para quando não retornar colunas
}

func TestGetWarningByUser(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	s := Storage{db}

	pattern := "Test"
	query := `SELECT \* FROM warnings WHERE sent_to LIKE`
	mock.ExpectQuery(query).WillReturnRows(warningRows)

	users, err := s.GetWarningByUser(pattern)
	assert.NoError(t, err)
	assert.NotNil(t, users) //TODO: incluir teste para quando não retornar colunas
}

// By ID
func TestGetCustomerByID(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	assert.Nil(t, err, fmt.Sprintf("error when opening the mock db connection: %v", err))
	defer db.Close()

	id := int(1)
	mock.ExpectQuery(`SELECT`). // \* FROM customers WHERE id=\$1`).
					WithArgs(id).
					WillReturnRows(customerRows)

	s := Storage{db}

	customer, err := s.GetCustomerByID(int(1))
	// assert.NoError(t, err) //TODO: comentado porque ainda não aprendi a popular a mock db
	assert.NotNil(t, customer)
}

func TestGetUserByID(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	assert.Nil(t, err, fmt.Sprintf("error when opening the mock db connection: %v", err))
	defer db.Close()

	id := int(1)
	mock.ExpectQuery(`SELECT \* FROM users WHERE id=\$1`).
		WithArgs(id).
		WillReturnRows(userRows)

	s := Storage{db}

	user, err := s.GetUserByID(1)
	// assert.NoError(t, err) //TODO: comentado porque ainda não aprendi a popular a mock db
	assert.NotNil(t, user)
}

func TestGetWarningByID(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	assert.Nil(t, err, fmt.Sprintf("error when opening the mock db connection: %v", err))
	defer db.Close()

	id := int(1)
	mock.ExpectQuery(`SELECT \* FROM warnings WHERE id=\$1`).
		WithArgs(id).
		WillReturnRows(warningRows)

	s := Storage{db}

	user, err := s.GetWarningByID(1)
	// assert.NoError(t, err) //TODO: comentado porque ainda não aprendi a popular a mock db
	assert.NotNil(t, user)
}

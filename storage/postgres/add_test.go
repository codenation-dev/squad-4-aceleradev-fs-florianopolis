package postgres

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAddCustomer(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	assert.Nil(t, err, fmt.Sprintf("error when opening the mock db connection: %v", err))
	defer db.Close()

	query := `INSERT INTO customers \(name, wage, is_public, sent_warning\)
	VALUES \(\$1, \$2, \$3, \$4\)`
	expResult := sqlmock.NewResult(1, 1)
	mock.ExpectExec(query).
		WithArgs(mc.Name, mc.Wage, mc.IsPublic, mc.SentWarning).WillReturnError(nil).
		WillReturnResult(expResult)

	s := Storage{db}
	err = s.AddCustomer(mc)
	assert.NoError(t, err, err)
	// expLII, _ := expResult.LastInsertId()
	// expRA, _ := expResult.RowsAffected()
	// LII, _ := res.LastInsertId()
	// RA, _ := res.RowsAffected()
	// assert.Equal(t, expLII, LII)
	// assert.Equal(t, expRA, RA)
}

func TestAddUser(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	assert.Nil(t, err, fmt.Sprintf("error when opening the mock db connection: %v", err))
	defer db.Close()

	query := `INSERT INTO users \(login, email, pass\)
	VALUES \(\$1, \$2, \$3\)`
	expResult := sqlmock.NewResult(1, 1)
	mock.ExpectExec(query).
		WithArgs(mu.Login, mu.Email, mu.Pass).WillReturnError(nil).
		WillReturnResult(expResult)

	s := Storage{db}
	err = s.AddUser(mu)
	assert.NoError(t, err, err)
	// expLII, _ := expResult.LastInsertId()
	// expRA, _ := expResult.RowsAffected()
	// LII, _ := res.LastInsertId()
	// RA, _ := res.RowsAffected()
	// assert.Equal(t, expLII, LII)
	// assert.Equal(t, expRA, RA)
}

package postgres

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateCustomer(t *testing.T) {
	t.Parallel()
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

func TestUpdateUser(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	s := Storage{db}
	query := `UPDATE users
			SET login=\$1, email=\$2, pass=\$3
			WHERE id=\$4`
	mock.ExpectExec(query).
		WithArgs(mu.Login, mu.Email, mu.Pass, mu.ID).
		WillReturnError(errors.New("could not update the user"))
	// WillReturnResult(sqlmock.NewResult(1, 1))
	err = s.UpdateUser(mu)
	assert.Error(t, err, err)
}

func TestUpdateWaring(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	s := Storage{db}
	query := `UPDATE warnings
			SET dt=\$1, message=\$2, sent_to=\$3, from_customer=\$4
			WHERE id=\$5`
	mock.ExpectExec(query).
		WithArgs(mw.Dt, mw.Message, mw.SentTo, mw.FromCustomer, mw.ID).
		WillReturnError(errors.New("could not update the warning"))
	// WillReturnResult(sqlmock.NewResult(1, 1))
	err = s.UpdateWarning(mw)
	assert.Error(t, err, err)
}

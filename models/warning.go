package models

import (
	"fmt"
)

// Warning Models the warning
type Warning struct {
	WID          int32  `json:"wid"`
	Dt           string `json:"dt"`
	Message      string `json:"msg"`
	SentTo       string `json:"sent_to"`
	FromCustomer string `json:"from_customer"`
}

// NewWarning insert a new warning on the db
func NewWarning(wa Warning) (bool, error) {
	db := Connect()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return false, err
	}
	sql := `INSERT INTO warnings (dt, msg, sent_to, from_customer)
			VALUES ($1, $2, $3, $4)
			RETURNING WID`
	{
		stmt, err := tx.Prepare(sql)
		if err != nil {
			tx.Rollback()
			return false, err
		}
		defer stmt.Close()
		err = stmt.QueryRow(
			wa.Dt, wa.Message, wa.SentTo, wa.FromCustomer,
		).Scan(&wa.WID)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}

	return true, tx.Commit()
}

// GetWarnings returns a list of all bank's customers
func GetWarnings() ([]Warning, error) {
	db := Connect()
	defer db.Close()
	sql := "SELECT * FROM warnings"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer rows.Close()
	var warnings []Warning
	for rows.Next() {
		var wa Warning
		err := rows.Scan(&wa.WID, &wa.Dt, &wa.Message, &wa.SentTo, &wa.FromCustomer)
		if err != nil {
			return nil, err
		}
		warnings = append(warnings, wa)
	}
	return warnings, nil
}

// GetWarningsByUser returns the warnings
func GetWarningsByUser(u string) ([]Warning, error) {
	db := Connect()
	defer db.Close()
	sql := "SELECT * FROM warnings WHERE sent_to = $1"
	rows, err := db.Query(sql, u)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer rows.Close()
	var warnings []Warning
	for rows.Next() {
		var wa Warning
		err := rows.Scan(&wa.WID, &wa.Dt, &wa.Message, &wa.SentTo, &wa.FromCustomer)
		if err != nil {
			return nil, err
		}
		warnings = append(warnings, wa)
	}
	return warnings, nil
}

package models

// UpdateCustomer updates a customer
// func UpdateCustomer(sql, i interface) (int64, error) {
// 	db := Connect()
// 	defer db.Close()
// 	stmt, err := db.Prepare(sql)
// 	if err != nil {
// 		return 0, err
// 	}
// 	// TODO: tem como pegar dinamicamente os valores para colocar na Exec???
// 	rows, err := stmt.Exec(c.Name, c.Wage, c.IsPublic, c.SentWarning, c.CID)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return rows.RowsAffected()
// }

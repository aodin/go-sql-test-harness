package main

import (
	"database/sql"
)

// Good closes the transaction it creates
func Good(conn *sql.DB) error {
	tx, err := conn.Begin()
	defer tx.Rollback()
	rows, err := tx.Query("select id from harness_tests")
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

// Bad does not close the transaction it creates
func Bad(conn *sql.DB) error {
	tx, err := conn.Begin()
	// Do not close the transaction
	rows, err := tx.Query("select id from harness_tests")
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func main() {
	println("Do nothing")
}

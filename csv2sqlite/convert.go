package csv2sqlite

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// Function to infer data type
func inferDataType(value string) string {
	if _, err := strconv.Atoi(value); err == nil {
		return "INTEGER"
	}
	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return "REAL"
	}
	if matched, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, value); matched {
		return "DATE"
	}
	if value == "true" || value == "false" {
		return "BOOLEAN"
	}
	return "TEXT"
}

// Function to convert csv to sqlite3 database
func Convert(dbName string, tableName string, csvFilePath string) error {
	// Open the CSV file
	file, err := os.Open(csvFilePath)
	if err != nil {
		return fmt.Errorf("error opening csv file: %v", err)
	}
	defer file.Close()

	// Read the headers
	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		return fmt.Errorf("error reading csv file headers: %v", err)
	}

	// Read the first row to infer the data types
	firstRow, err := reader.Read()
	if err != nil {
		return fmt.Errorf("error reading csv file first row: %v", err)
	}

	// Infer the data types
	columnTypes := make([]string, len(headers))
	for i, value := range firstRow {
		columnTypes[i] = inferDataType(value)
	}

	// Open the database
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	defer db.Close()

	// Construct the CREATE TABLE query
	createTableQuery := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", tableName)
	for i, header := range headers {
		createTableQuery += fmt.Sprintf("%s %s", header, columnTypes[i])
		if i < len(headers)-1 {
			createTableQuery += ", "
		}
	}
	createTableQuery += ");"

	// Execute the CREATE TABLE query
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}

	// Prepare the INSERT query
	insertQuery := fmt.Sprintf("INSERT INTO %s VALUES (", tableName)
	for i := range headers {
		insertQuery += "?"
		if i < len(headers)-1 {
			insertQuery += ","
		}
	}
	insertQuery += ")"

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}
	defer tx.Rollback()

	// Prepare the INSERT statement
	stmt, err := tx.Prepare(insertQuery)
	if err != nil {
		return fmt.Errorf("error preparing insert statement: %v", err)
	}
	defer stmt.Close()

	// Insert the first row
	values := make([]interface{}, len(firstRow))
	for i, v := range firstRow {
		values[i] = v
	}
	_, err = stmt.Exec(values...)
	if err != nil {
		return fmt.Errorf("error inserting first row: %v", err)
	}

	// Insert remaining rows
	for {
		row, err := reader.Read()
		if err != nil {
			break // EOF or error
		}
		values := make([]interface{}, len(row))
		for i, v := range row {
			values[i] = v
		}
		_, err = stmt.Exec(values...)
		if err != nil {
			return fmt.Errorf("error inserting row: %v", err)
		}
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}

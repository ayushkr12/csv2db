package main

import (
	"fmt"
	"os"

	"github.com/ayushkr12/csv2db/csv2sqlite"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: csv2db <db_name> <table_name> <csv_file>")
		return
	}

	dbName := os.Args[1]
	tableName := os.Args[2]
	csvFilePath := os.Args[3]

	err := csv2sqlite.ConvertCSVToDB(dbName, tableName, csvFilePath)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("CSV successfully converted to SQLite database!")
	}
}

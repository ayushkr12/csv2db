# CSV to DB

A simple Go project to convert CSV files into SQLite databases.

## Installation

To install the CLI tool, use the following command:

```
go install github.com/ayushkr12/csv2db/cmd/csv2db@latest
```

This will install the `csv2db` command globally.

## Usage

You can use this project as either a package or a CLI tool.

### As a Go Package

You can use this as a Go package in your projects:

```go
package main

import (
    "github.com/ayushkr12/csv2db/csv2sqlite"
)

func main() {
    err := csv2sqlite.ConvertCSVToDB("mydb.db", "users", "data.csv")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("CSV successfully converted to SQLite!")
    }
}
```

### As a Command-Line Tool

To use it from the command line:

```
csv2db <db_name> <table_name> <csv_file>
```

Where:
- `<db_name>`: Name of the SQLite database.
- `<table_name>`: Name of the table in the SQLite database.
- `<csv_file>`: Path to the CSV file to be converted.

Example:
```
csv2db mydb.db users data.csv
```

This will create or append to `mydb.db`, creating a table `users` with columns inferred from the CSV file.

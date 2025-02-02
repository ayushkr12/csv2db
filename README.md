# CSV to DB

A simple Go project to convert CSV files into SQLite databases.

## Installation

To install the CLI tool, use the following command:

```sh
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
    err := csv2sqlite.Convert("mydb.db", "users", "data.csv")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("CSV successfully converted to SQLite!")
    }
}
```

### As a Command-Line Tool

To use it from the command line:

```sh
csv2db <output_db_path> <table_name> <csv_file>
```

Where:

- `<output_db_path>`: Path to the output sqlite database.
- `<table_name>`: Name of the table in the SQLite database.
- `<csv_file>`: Path to the CSV file to be converted.

Example:

```sh
csv2db mydb.db users data.csv
```

This will create or append to `mydb.db`, creating a table `users` with columns inferred from the CSV file.

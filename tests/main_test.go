package tests

import (
	"testing"

	"github.com/ayushkr12/csv2db/csv2sqlite"
)

func TestConvertCSVToDB(t *testing.T) {
	err := csv2sqlite.Convert("test.db", "test_table", "test.csv")
	if err != nil {
		t.Errorf("Error converting CSV to DB: %v", err)
	}
}

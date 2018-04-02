package csvhelper

import (
	"encoding/csv"
	"os"
)

// ReadFile reads CSV file and return CSV records.
func ReadFile(f string) ([][]string, error) {
	reader, err := os.Open(f)
	if err != nil {
		return [][]string{}, err
	}
	defer reader.Close()

	// Create a new CSV reader.
	r := csv.NewReader(reader)
	return r.ReadAll()
}

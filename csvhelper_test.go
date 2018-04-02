package csvhelper_test

import (
	"log"

	"github.com/northbright/csvhelper"
)

func ExampleReadFile() {
	f := "files/example.csv"

	records, err := csvhelper.ReadFile(f)
	if err != nil {
		log.Printf("ReadFile() error: %v", err)
		return
	}

	log.Printf("records: %v", records)
	// Output:
}

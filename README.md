# csvhelper

[![Build Status](https://travis-ci.org/northbright/csvhelper.svg?branch=master)](https://travis-ci.org/northbright/csvhelper)
[![GoDoc](https://godoc.org/github.com/northbright/csvhelper?status.svg)](https://godoc.org/github.com/northbright/csvhelper)

Package csvhelper provides helper functions for CSV files.

#### Example of Reading CSV File

    f := "files/example.csv"

    records, err := csvhelper.ReadFile(f)
    if err != nil {
            log.Printf("ReadFile() error: %v", err)
            return
    }

    log.Printf("records: %v", records)


#### Documentation
* [API References](https://godoc.org/github.com/northbright/csvhelper)

#### License
* [MIT License](./LICENSE)

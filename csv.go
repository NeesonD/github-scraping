package main

import (
	"github.com/gocarina/gocsv"
	"os"
)

func WriteCsv(file string, data interface{}) {
	clientsFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	if _, err := clientsFile.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}

	err = gocsv.MarshalFile(data, clientsFile) // Use this to save the CSV back to the file
	if err != nil {
		panic(err)
	}
}

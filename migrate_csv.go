package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	fileName string = "logins.csv"
)

func main() {
	inputFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Could not open file", err)
	}
	defer inputFile.Close()
	reader := csv.NewReader(inputFile)
	// reader.Comma = ','

	if _, err := reader.Read(); err != nil {
		log.Fatal("Error reading file", err)
	}

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Could not process file entry", err)
		}
		fmt.Println(line)
	}
}

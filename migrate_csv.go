package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

var (
	fileName    string = "logins.csv"
	outFileName string = "logins_processed.csv"
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

	outputFile, err := os.Create(outFileName)
	if err != nil {
		log.Fatal("Could not open file", err)
	}
	defer outputFile.Close()
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	if err = writer.Write([]string{"url", "username", "password"}); err != nil {
		log.Fatal("Could not write to output file", err)
	}

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Could not process file entry", err)
		}
		//fmt.Println(line)
		err = writer.Write([]string{line[0], line[1], line[2]})
	}
}

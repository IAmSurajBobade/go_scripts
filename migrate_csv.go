package main

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

var (
	fileName        string = "logins.csv"
	outFileName     string = "logins_processed.csv"
	errInvalidEntry error  = errors.New("invalid entry")
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
			log.Println("Could not process input", err)
			continue
		}
		//fmt.Println(line)
		if processedRecord, err := processLine(line); err == nil {
			err = writer.Write(processedRecord)
			if err != nil {
				log.Println("Could not write entry", err)
				continue
			}
		}
	}
}

func processLine(line []string) ([]string, error) {
	if len(line) < 3 {
		return nil, errInvalidEntry
	}
	if line[0] == "" || !strings.HasPrefix(line[0], "http") {
		return nil, errInvalidEntry
	}
	if line[1] == "" || line[2] == "" {
		return nil, errInvalidEntry
	}
	return []string{line[0], line[1], line[2]}, nil
}

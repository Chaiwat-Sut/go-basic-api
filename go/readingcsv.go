package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type student struct {
	Lastname  string
	Firstname string
	SSN       string
	Test1     float64
	Test2     float64
	Test3     float64
	Test4     float64
	Final     float64
	Grade     string
}

func main() {
	//Open file
	csvfile, err := os.Open("../homework/api/rawdata/grades.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(csvfile)
	r.LazyQuotes = true
	const expectedRecord int = 9
	lineCount := 0
	students := []student{}
	// skip first line

	for {
		record, err := r.Read()
		if lineCount == 0 {
			lineCount++
			continue
		}
		if err == io.EOF { // If read to the end of file
			break
		}
		if len(record) != expectedRecord {
			fmt.Printf("Warning skiping the line with the wrong number of field (%d instead of %d)\n", len(record), expectedRecord)
			continue
		}
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(record)
		test1, err := strconv.ParseFloat(strings.TrimSpace(record[3]), 64)
		if err != nil {
			fmt.Println(err)
		}
		test2, err := strconv.ParseFloat(strings.TrimSpace(record[4]), 64)
		if err != nil {
			fmt.Println(err)
		}
		test3, err := strconv.ParseFloat(strings.TrimSpace(record[5]), 64)
		if err != nil {
			fmt.Println(err)
		}
		test4, err := strconv.ParseFloat(strings.TrimSpace(record[6]), 64)
		if err != nil {
			fmt.Println(err)
		}
		final, err := strconv.ParseFloat(strings.TrimSpace(record[7]), 64)
		if err != nil {
			fmt.Println(err)
		}
		students = append(students, student{
			Lastname:  record[0],
			Firstname: record[1],
			SSN:       record[2],
			Test1:     test1,
			Test2:     test2,
			Test3:     test3,
			Test4:     test4,
			Final:     final,
			Grade:     record[8],
		})
		lineCount++
	}

	for _, student := range students {
		fmt.Println(student)
	}
}

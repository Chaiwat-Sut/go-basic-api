package readcsv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"studentmodule/models"
)

func ReadCsvFile(r *csv.Reader) []models.Student {
	students := []models.Student{}
	const expectedRecord int = 9
	lineCount := 0

	for {
		record, err := r.Read()
		if lineCount == 0 { // skip the first line
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
		students = append(students, models.Student{
			Lastname:  deleteQuotes(strings.TrimSpace(record[0])),
			Firstname: deleteQuotes(strings.TrimSpace(record[1])),
			SSN:       deleteQuotes(strings.TrimSpace(record[2])),
			Test1:     test1,
			Test2:     test2,
			Test3:     test3,
			Test4:     test4,
			Final:     final,
			Grade:     deleteQuotes(strings.TrimSpace(record[8])),
		})
		lineCount++
	}

	return students
}

func deleteQuotes(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}

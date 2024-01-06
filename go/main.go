package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"studentmodule/dbmanage"
	"studentmodule/jsonhandle"
	"studentmodule/readcsv"
	"studentmodule/typestruct"
)

func main() {
	csvfile, err := os.Open("../homework/api/rawdata/grades.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvfile.Close()

	r := csv.NewReader(csvfile)
	r.LazyQuotes = true
	students := readcsv.ReadCsvFile(r)

	writeJsonFile(students)
	database := dbmanage.InitDB()
	students = jsonhandle.ReadJsonFile("../homework/json/students.json").([]typestruct.Student)

	for _, student := range students {
		dbmanage.InsertStudentTable(student, database)
	}
}

func writeJsonFile(students []typestruct.Student) {
	file, err := os.Create("../homework/json/students.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(students)
	if err != nil {
		log.Fatal(err)
	}
}

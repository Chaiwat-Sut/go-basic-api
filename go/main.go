package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"studentmodule/dbmanage"
	"studentmodule/jsonhandle"
	"studentmodule/models"
	"studentmodule/readcsv"
	"studentmodule/restapi"
)

func main() {
	// Open Csv
	csvfile, err := os.Open("../homework/api/rawdata/grades.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvfile.Close()

	// Read Csv
	r := csv.NewReader(csvfile)
	r.LazyQuotes = true
	students := readcsv.ReadCsvFile(r)

	jsonhandle.WriteJsonFile(students)
	students = jsonhandle.ReadJsonFile("../homework/json/students.json").([]models.Student)

	for _, student := range students {
		dbmanage.InsertStudentTable(student)
	}

	restapi.SetupDB()
	restapi.SetupRoute("/api")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package jsonhandle

import (
	"encoding/json"
	"log"
	"os"
	"studentmodule/models"
)

func ReadJsonFile(path string) interface{} {
	var students []models.Student

	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(bytes, &students)
	if err != nil {
		log.Fatal(err)
	}

	return students
}

func WriteJsonFile(students []models.Student) {
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

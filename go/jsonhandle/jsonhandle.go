package jsonhandle

import (
	"encoding/json"
	"log"
	"os"
	"studentmodule/typestruct"
)

func ReadJsonFile(path string) interface{} {
	var students []typestruct.Student

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

package restapi

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"studentmodule/models"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

const studentPath = "student"

func SetupDB() {
	var err error
	Db, err = sql.Open("sqlite3", "../homework/database/database.db")
	if err != nil {
		log.Fatal(err)
	}
	if Db == nil {
		log.Println("Database is nil")
	}

	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(5)
	Db.SetMaxIdleConns(5)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", studentPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastname := urlPathSegments[len(urlPathSegments)-1]
	switch r.Method {
	case http.MethodGet:
		student, err := getStudentByLastname(lastname)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if student == nil {
			w.WriteHeader(http.StatusNotFound)
		}
		j, err := json.Marshal(student)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusBadRequest)
		}

		i, err := w.Write(j)
		if err != nil {
			log.Fatal(err, i)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getStudentByLastname(lastname string) (*models.Student, error) {
	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancle()
	if Db == nil {
		return nil, errors.New("database connection is nil")
	}
	row := Db.QueryRowContext(ctx, `SELECT * 
	FROM student 
	WHERE lastname = ?`, lastname)

	student := &models.Student{}
	err := row.Scan(
		&student.SSN,
		&student.Firstname,
		&student.Lastname,
		&student.Test1,
		&student.Test2,
		&student.Test3,
		&student.Test4,
		&student.Final,
		&student.Grade,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Println("Succesful GET method")
	return student, nil
}

func SetupRoute(apiBasePath string) {
	// studentHandle := http.HandlerFunc(handleRequest)
	// http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, studentPath), corsMiddleware(studentHandle))
	http.HandleFunc(fmt.Sprintf("%s/%s/", apiBasePath, studentPath), handleRequest)
}

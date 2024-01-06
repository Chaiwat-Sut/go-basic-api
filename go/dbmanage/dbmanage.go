package dbmanage

import (
	"database/sql"
	"log"
	"studentmodule/models"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func InitDB() {
	Db, _ = sql.Open("sqlite3", "../homework/database/database.db")
	statement, err := Db.Prepare(`CREATE TABLE IF NOT EXISTS student 
		(ssn TEXT PRIMARY KEY,firstname TEXT,lastname 	
		TEXT,test1 REAL,test2 REAL,test3 REAL,test4 
		REAL,final REAL,grade TEXT)`)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()

}

func InsertStudentTable(student models.Student) {
	statement, err := Db.Prepare(`INSERT INTO student 
		(ssn,firstname,lastname,test1,test2,test3,test4,final,grade) 
		VALUES (?,?,?,?,?,?,?,?,?)`)

	if err != nil {
		log.Fatal(err)
	}
	statement.Exec(student.SSN, student.Firstname, student.Lastname, student.Test1, student.Test2, student.Test3, student.Test4, student.Final, student.Grade)
}

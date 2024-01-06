package dbmanage

import (
	"database/sql"
	"studentmodule/typestruct"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	database, _ := sql.Open("sqlite3", "../homework/database/database.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS student (ssn TEXT PRIMARY KEY,firstname TEXT,lastname TEXT,test1 REAL,test2 REAL,test3 REAL,test4 REAL,final REAL,grade TEXT)")
	statement.Exec()

	return database
}

func InsertStudentTable(student typestruct.Student, database *sql.DB) {
	statement, _ := database.Prepare("INSERT INTO student (ssn,firstname,lastname,test1,test2,test3,test4,final,grade) VALUES (?,?,?,?,?,?,?,?,?)")
	statement.Exec(student.SSN, student.Firstname, student.Lastname, student.Test1, student.Test2, student.Test3, student.Test4, student.Final, student.Grade)
}

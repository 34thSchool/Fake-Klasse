package storage

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

var StudentsSchema = `CREATE TABLE IF NOT EXISTS students (
	name	TEXT,
	surname	TEXT
	);`
var selectStudents = `SELECT rowid, name, surname FROM students`
var insertStudent = `INSERT INTO students (name, surname) VALUES(?,?)`
var deleteAllStudents = `DELETE FROM students`
var deleteStudentByID = `DELETE FROM students WHERE rowid IN (?)`

type Student struct {
	Rowid   int
	Name    string
	Surname string
}

type Storage struct {
	db *sqlx.DB
}

var Singleton *Storage = &Storage{}// Singleton

func (storage *Storage) Init(path string) {
	// Creating and/or opening DB:
	db, err := sqlx.Open("sqlite", path)
	if err != nil {
		log.Fatal("failed to open SQLite DB: ", db)
	}
	storage.db = db

	// Creating tables if they don't exist:
	storage.CreateTable(StudentsSchema)
}
func (storage *Storage) CreateTable(schema string) {
	// Creating tables if they don't exist:
	result, err := storage.db.Exec(schema)
	if err != nil {
		log.Fatal("table creation failed. Query: ", schema, "\nError:", err)
	}
	if count, err := result.RowsAffected(); err != nil {
		log.Fatal(count, " rows affected.")
	}
}
func (storage *Storage) Close() { // To control flow in other files.
	storage.db.Close()
}
func (storage *Storage) AddStudent(name, surname string) {
	result, err := storage.db.Exec(insertStudent, name, surname)
	if err != nil {
		log.Fatal("table creation failed. Query: ", insertStudent, "\nError:", err)
	}
	if count, err := result.RowsAffected(); err != nil {
		log.Fatal(count, " rows affected.")
	}
}
func (storage *Storage) GetStudents() *[]Student { // Writes all the students from table to slice.

	var students []Student

	if err := storage.db.Select(&students, selectStudents); err != nil {
		log.Fatal("querying 'students' table failed. Query: ", selectStudents, "\nError: ", err)
	}

	return &students
}
func (storage *Storage) PrintStudents() {
	// Getting students:
	students := storage.GetStudents()

	// Printing:
	for _, student := range *students {
		fmt.Print(student.Rowid, ": ", student.Name, " ", student.Surname, "\n")
	}
}
func (storage *Storage) DeleteAllStudents() {
	result, err := storage.db.Exec(deleteAllStudents)
	if err != nil {
		log.Fatal("student deletion failed. Query: ", deleteAllStudents, "\nError:", err)
	}
	if count, err := result.RowsAffected(); err != nil {
		log.Fatal(count, " rows affected.")
	}
}
func (storage *Storage) DeleteStudent(id int) {
	result, err := storage.db.Exec(deleteStudentByID, id)
	if err != nil {
		log.Fatal("student deletion failed. Query: ", deleteStudentByID, "\nError:", err)
	}
	if count, err := result.RowsAffected(); err != nil {
		log.Fatal(count, " rows affected.")
	}
}


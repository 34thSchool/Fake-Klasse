package storage

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

// Queries:
// Students:
var studentsSchema = `CREATE TABLE IF NOT EXISTS students (name TEXT, surname TEXT, class TEXT, 
														FOREIGN KEY ("class") REFERENCES "classes"("name"));`
var selectStudents = `SELECT rowid, name, surname, class FROM students`
var insertStudent = `INSERT INTO students (name, surname, class) VALUES(?, ?, ?)`
var deleteAllStudents = `DELETE FROM students`
var deleteStudentByID = `DELETE FROM students WHERE rowid IN (?)`

// Classes:
var classesSchema = `CREATE TABLE IF NOT EXISTS classes (name TEXT UNIQUE);`
var selectAllClasses = `SELECT rowid, name FROM classes`
var insertClass = `INSERT OR IGNORE INTO classes (name) VALUES (?)`
var deleteAllClasses = `DELETE FROM classes`
var deleteClassByID = `DELETE FROM classes WHERE rowid IN (?)`
var linkClassToStudent = `SELECT name FROM classes INNER JOIN students ON classes.name = students.class;`

// Storage:
type Storage struct {
	db *sqlx.DB
}

var Singleton *Storage = &Storage{} // Singleton
func (storage *Storage) Init(path string) {
	// Creating and/or opening DB:
	db, err := sqlx.Open("sqlite", path)
	if err != nil {
		log.Fatal("failed to open SQLite DB: ", db)
	}
	storage.db = db

	// Creating tables if they don't exist:
	storage.CreateTable(studentsSchema)
	storage.CreateTable(classesSchema)
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

// Student:
type Student struct {
	Rowid   int
	Name    string
	Surname string
	Class   string
}

func (storage *Storage) AddStudent(name, surname string, className string) {

	// Inserting student into students table:
	result, err := storage.db.Exec(insertStudent, name, surname, className)
	if err != nil {
		log.Fatal("student insertion failed. Query: ", insertStudent, "\nError:", err)
	}
	if count, err := result.RowsAffected(); err != nil {
		log.Fatal(count, " rows affected.")
	}
	// Linking class to student:
	// Checking if class with this name doesn't exist already:
	rows, err := storage.db.Query("SELECT * FROM classes WHERE name = (?)", className)
	if err != nil {
		log.Fatal(err)
	}
	rows.Close()

	//storage.AddClass(class.Name)

	// Linking class to student:
	storage.db.Exec(linkClassToStudent)
}
func (storage *Storage) GetAllStudents() *[]Student { // Writes all the students from table to slice.

	var students []Student

	// if err := storage.db.Select(&students, selectStudents); err != nil {
	// 	log.Fatal("querying 'students' table failed. Query: ", selectStudents, "\nError: ", err)
	// }

	rows, err := storage.db.Query(selectStudents)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for i := 0; rows.Next(); i++ {
		students = append(students, Student{})
		err := rows.Scan(&students[i].Rowid, &students[i].Name, &students[i].Surname, &students[i].Class)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &students
}
func (storage *Storage) GetClassStudents(class Class) *[]Student {

	var students []Student

	rows, err := storage.db.Query(`SELECT * FROM students WHERE class = (?)`, class.Name)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for i := 0; rows.Next(); i++ {
		students = append(students, Student{})
		err := rows.Scan( /*&students[i].Rowid,*/ &students[i].Name, &students[i].Surname, &students[i].Class)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &students
}
func (storage *Storage) PrintClassStudents(class Class) {
	// Getting students:
	students := storage.GetClassStudents(class)

	// Printing:
	for _, student := range *students {
		fmt.Print(student.Name, " ", student.Surname, " ", student.Class, "\n")
	}
}
func (storage *Storage) PrintAllStudents() {
	// Getting students:
	students := storage.GetAllStudents()

	// Printing:
	for _, student := range *students {
		fmt.Print(student.Rowid, ": ", student.Name, " ", student.Surname, " ", student.Class, "\n")
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
func (storage *Storage) UpdateStudentClass(prevClassName, newClassName string) {
	update := `UPDATE students SET class = (?) WHERE class = (?)`
	_, err := storage.db.Exec(update, newClassName, prevClassName)
	if err != nil {
		log.Fatal("student update failed. Query: ", update, "\nError:", err)
	}
	// if count, err := result.RowsAffected(); err != nil {
	// 	log.Fatal(count, " rows affected.")
	// }
}

// func (storage *Storage) WipeStudentClass(prevClassName string){
// 	update := `UPDATE students SET class = NUL WHERE class = (?)`
// 	result, err := storage.db.Exec(update, prevClassName)
// 	if err != nil {
// 		log.Fatal("student update failed. Query: ", update, "\nError:", err)
// 	}
// 	if count, err := result.RowsAffected(); err != nil {
// 		log.Fatal(count, " rows affected.")
// 	}
// }
// Class:
type Class struct {
	Rowid int
	Name  string
}

func (storage *Storage) AddClass(className string) {
	// Adding class to classes table:
	result, err := storage.db.Exec(insertClass, className) //class.Name
	if err != nil {
		log.Fatal("class insertion failed. Query: ", insertClass, "\nError:", err)
	}
	if count, err := result.RowsAffected(); err != nil {
		log.Fatal(count, " rows affected.")
	}
}
func (storage *Storage) DeleteAllClasses() {
	result, err := storage.db.Exec(deleteAllClasses)
	if err != nil {
		log.Fatal("class deletion failed. Query: ", deleteAllClasses, "\nError:", err)
	}
	if count, err := result.RowsAffected(); err != nil {
		log.Fatal(count, " rows affected.")
	}
}
func (storage *Storage) DeleteClass(class Class) {

	//fmt.Println("Delete Class: ", class)
	result, err := storage.db.Exec(deleteClassByID, class.Rowid)
	if err != nil {
		log.Fatal(" DEL!!! class deletion failed. Query: ", deleteClassByID, "\nError:", err)
	}
	if count, err := result.RowsAffected(); err != nil {
		log.Fatal(count, " rows affected.")
	}

	// Annulating class field for all the students from this class:
	//storage.UpdateStudentClass(class.Name, "")//decided not to do it here

	//storage.Singleton.UpdateStudentClass((*storage.Singleton.GetAllClasses())[id].Name, "")
}
func (storage *Storage) GetAllClasses() *[]Class {
	var classes []Class

	rows, err := storage.db.Query(selectAllClasses)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for i := 0; rows.Next(); i++ {
		classes = append(classes, Class{})
		err := rows.Scan(&classes[i].Rowid, &classes[i].Name)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &classes
}
func (storage *Storage) GetClassByID(id int) Class {

	classes := storage.GetAllClasses()

	return (*classes)[id]
}
func (storage *Storage) PrintAllClasses() {

	classes := storage.GetAllClasses()
	for _, class := range *classes {
		fmt.Print(class, "\n")
	}
}

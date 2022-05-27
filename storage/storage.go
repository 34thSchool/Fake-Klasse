package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

// Queries:
const (
	// Students:
	studentsSchema = `CREATE TABLE IF NOT EXISTS students (name TEXT, surname TEXT, class TEXT, 
				FOREIGN KEY ("class") REFERENCES "classes"("name"));`
	selectStudents    = `SELECT rowid, name, surname, class FROM students`
	insertStudent     = `INSERT INTO students (name, surname, class) VALUES(?, ?, ?)`
	deleteAllStudents = `DELETE FROM students`
	deleteStudentByID = `DELETE FROM students WHERE rowid IN (?)`

	// Classes:
	classesSchema      = `CREATE TABLE IF NOT EXISTS classes (name TEXT UNIQUE);`
	selectAllClasses   = `SELECT rowid, name FROM classes`
	insertClass        = `INSERT OR IGNORE INTO classes (name) VALUES (?)`
	deleteAllClasses   = `DELETE FROM classes`
	deleteClassByID    = `DELETE FROM classes WHERE rowid IN (?)`
	linkClassToStudent = `SELECT name FROM classes INNER JOIN students ON classes.name = students.class;`
)

// Storage:
type Storage struct {
	db *sqlx.DB
}

func (s *Storage) Init(path string) error {
	// Creating and/or opening DB:
	db, err := sqlx.Open("sqlite", path)
	if err != nil {
		//log.Fatal("failed to open SQLite DB: ", db)
		return err
	}
	s.db = db

	// Creating tables if they don't exist:
	s.CreateTable(studentsSchema)
	s.CreateTable(classesSchema)

	return nil
}
func (s *Storage) CreateTable(schema string) error {
	// Creating tables if they don't exist:
	_, err := s.db.Exec(schema)
	return err
}
func (s *Storage) Close() { // To control flow in other files.
	s.db.Close()
}

// Student:
type Student struct {
	Rowid   int
	Name    string
	Surname string
	Class   string
}

func (s *Storage) AddStudent(name, surname string, className string) error {

	// Inserting student into students table:
	_, err := s.db.Exec(insertStudent, name, surname, className)
	if err != nil {
		//log.Fatal("student insertion failed. Query: ", insertStudent, "\nError:", err)
		return err
	}

	// Linking class to student:
	// Checking if class with this name doesn't exist already:
	rows, err := s.db.Query("SELECT * FROM classes WHERE name = (?)", className)
	if err != nil {
		//log.Fatal(err)
		return err
	}
	rows.Close()

	//s.AddClass(class.Name)

	// Linking class to student:
	_, err = s.db.Exec(linkClassToStudent)
	
	return err
}
func (s *Storage) GetAllStudents() ([]Student, error){ // Writes all the students from table to slice.

	var students []Student

	if err := s.db.Select(&students, selectStudents); err != nil {
		//log.Fatal("querying 'students' table failed. Query: ", selectStudents, "\nError: ", err)
		return nil, err
	}

	return students, nil
}
func (s *Storage) GetClassStudents(class Class) ([]Student, error) {

	var students []Student

	if err := s.db.Select(&students, `SELECT * FROM students WHERE class = (?)`, class.Name); err != nil{
		//log.Fatal(err)
		return nil, err
	}

	return students, nil
}
func (s *Storage) PrintClassStudents(class Class) error {
	// Getting students:
	students, err := s.GetClassStudents(class)
	if err != nil{
		return err
	}

	// Printing:
	for _, student := range students {
		fmt.Print(student.Name, " ", student.Surname, " ", student.Class, "\n")
	}

	return nil
}
func (s *Storage) PrintAllStudents() error {
	// Getting students:
	students, err := s.GetAllStudents()
	if err != nil{
		return err
	}
	// Printing:
	for _, student := range students {
		fmt.Print(student.Rowid, ": ", student.Name, " ", student.Surname, " ", student.Class, "\n")
	}

	return nil
}
func (s *Storage) DeleteAllStudents() error {
	_, err := s.db.Exec(deleteAllStudents)
	return err
}
func (s *Storage) DeleteStudent(id int) error {
	_, err := s.db.Exec(deleteStudentByID, id)
	return err
}

// All students:
func (s *Storage) UpdateAllClassStudentsClass(prevClassName, newClassName string) error {
	update := `UPDATE students SET class = (?) WHERE class = (?)`
	_, err := s.db.Exec(update, newClassName, prevClassName)
	return err
}

// Single student:
func (s *Storage) updateStudentClass(student Student, newClassName string) error {
	update := `UPDATE students SET class = (?) WHERE rowid = (?)`
	_, err := s.db.Exec(update, newClassName, student.Rowid)
	return err
}
func (s *Storage) updateStudentName(student Student, newName string) error {
	update := `UPDATE students SET name = (?) WHERE rowid = (?)`
	_, err := s.db.Exec(update, newName, student.Rowid)
	return err
}
func (s *Storage) updateStudentSurname(student Student, newSurname string) error {
	update := `UPDATE students SET surname = (?) WHERE rowid = (?)`
	_, err := s.db.Exec(update, newSurname, student.Rowid)
	return err
}
func (s *Storage) UpdateStudent(prevStudent Student, newStudent Student) error {
	var err error

	if newStudent.Name != "" {
		err = s.updateStudentName(prevStudent, newStudent.Name)
		if err != nil{return err}
	}
	if newStudent.Surname != "" {
		err = s.updateStudentSurname(prevStudent, newStudent.Surname)
		if err != nil{return err}
	}
	if newStudent.Class != "" {
		err = s.updateStudentClass(prevStudent, newStudent.Class)
		if err != nil{return err}
	}

	return nil
}

// Class:
type Class struct {
	Rowid int
	Name  string
}

func (s *Storage) AddClass(className string) error {
	// Adding class to classes table:
	_, err := s.db.Exec(insertClass, className)
	return err
}
func (s *Storage) DeleteAllClasses() error {
	_, err := s.db.Exec(deleteAllClasses)
	return err
}
func (s *Storage) DeleteClass(class Class) error {

	//fmt.Println("Delete Class: ", class)
	_, err := s.db.Exec(deleteClassByID, class.Rowid)
	return err

	// Annulating class field for all the students from this class:
	//s.UpdateStudentClass(class.Name, "")//decided not to do it here

	//s.S.UpdateStudentClass((*s.S.GetAllClasses())[id].Name, "")

}
func (s *Storage) GetAllClasses() ([]Class, error) {
	var classes []Class

	if err := s.db.Select(&classes, selectAllClasses); err != nil{
		return nil, err
	}

	return classes, nil
}
func (s *Storage) GetClassByIndex(index int) (*Class, error) {

	classes, err := s.GetAllClasses()
	if err != nil{return nil, err}
	return &(classes)[index], nil
}
func (s *Storage) PrintAllClasses() error{

	classes, err := s.GetAllClasses()
	if err != nil{return err}
	for _, class := range classes {
		fmt.Print(class, "\n")
	}
	return nil
}

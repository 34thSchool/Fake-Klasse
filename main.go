package main

import(
	"fake-klasse/storage"
)


func main() {
	
	storage := storage.Storage{}
	
	storage.Init("School.db")
	
	defer storage.Close()

	storage.DeleteAllStudents()
	
	
	storage.AddStudent("Ezitis", "Migla")

	storage.AddStudent("Grizzly", "Bear")

	storage.PrintStudents()	

	storage.DeleteStudent(2)

	

	storage.PrintStudents()	
}
package main

import (
	"fmt"

	"learnGo/Patterns/1.2_postgres/models"
)

func main() {

	st, err := models.NewStudentService()
	if err != nil {
		panic(err)
	}

	//create a new student if not already
	st.NewStudent(123, "Brad", 3.0)

	student, err := st.StudentByID(123)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Student %s has GPA %d", student.Name, student.GPA)

}

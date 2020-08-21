package models

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 32768
	user     = "docker"
	dbname   = "docker"
	password = "docker"
)

const (
	RecordNotFound = "RecordNotFound"
)

type Student struct {
	gorm.Model
	StudentID uint `gorm:"not null;unique_index"`
	Name      string
	GPA       uint
}

type StudentService struct {
	db *gorm.DB
}

func generateConnectionString() (connectionString string) {
	connectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return connectionString

}

func NewStudentService() (*StudentService, error) {

	db, err := gorm.Open("postgres", generateConnectionString())
	if err != nil {
		return nil, err
	}

	studentService := &StudentService{
		db: db,
	}

	//migrate student table
	db.AutoMigrate(&Student{})

	return studentService, nil

}

func (st *StudentService) StudentByID(id uint) (*Student, error) {

	var student Student

	//find Bobby
	err := st.db.Where("student_id = ?", id).First(&student).Error
	switch err {
	case nil:
		return &student, nil
	case gorm.ErrRecordNotFound:
		return nil, errors.New(RecordNotFound)
	default:
		return nil, err

	}

}

//NewStudent creates a new user
func (st *StudentService) NewStudent(studentID uint, name string, gpa uint) {

	//create user object
	st.db.Create(&Student{
		StudentID: studentID,
		Name:      name,
		GPA:       gpa,
	})

}

func (st *StudentService) Close() {
	st.db.Close()
}

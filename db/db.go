package db

import (
	"github.com/rs/zerolog/log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StudentHandler struct {
	DB *gorm.DB
}

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    string `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"registration"`
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to initialize SQLite %s", err.Error())
	}

	// Migrate the schema
	db.AutoMigrate(&Student{})

	return db
}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

func (s *StudentHandler) GetStudents() ([]Student, error) {
	var students []Student
	if err := s.DB.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (s *StudentHandler) CreateStudent(student Student) error {
	if result := s.DB.Create(&student); result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to create student")
		return result.Error
	}
	log.Info().Msg("Student created successfully")
	return nil
}

// func GetStudentbyID(id string) (*Student, error) {
// 	var student Student
// 	if err := db.Fist(&student, "id = ?", id).Error; err != nil {
// 		return err
// 	}
// 	return &student, nil
// }

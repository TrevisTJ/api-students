package schemas

import (
	"gorm.io/gorm"
	"time"

)

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

type StudentResponse struct {
	ID       int       `json:"id"`
	CreateAt time.Time `json:"createdAt"`
	UpdateAt time.Time `json:"updatedAt"`
	DeleteAt time.Time `json:"deletedAt"`
	Name     string    `json:"name"`
	CPF      int       `json:"cpf"`
	Email    string    `json:"email"`
	Age      int       `json:"age"`
	Active   bool      `json:"active"`
}

func NewResponse(students []Student) []StudentResponse {
	studentsResponse := []StudentResponse{}

	for _, student := range students {
		studentResponse := StudentResponse{
			ID:       int(student.ID),
			CreateAt: student.CreatedAt,
			UpdateAt: student.UpdatedAt,
			Name:     student.Name,
			Email:    student.Email,
			Age:      student.Age,
			CPF:      student.CPF,
			Active:   student.Active,
		}
		studentsResponse =append(studentsResponse, studentResponse)
	}
	return studentsResponse
}

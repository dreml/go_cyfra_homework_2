package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	Login string `json:"login"`
	Password string `json:"password"`
	Phone string `json:"phone"`
	BirthDate time.Time `json:"birthDate"`
}

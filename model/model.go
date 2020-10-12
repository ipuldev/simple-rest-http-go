package model


import (
	"time"
)

type User struct {
	Base 
	Username string `json:"username"`
	Password string `json:"password"`
}

type Base struct{
	ID uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}



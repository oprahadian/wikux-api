package database

import (
	"database/sql"
)

type UserLogin struct {
	Id int64 `json:"id" db:"id"`
	Password string `json:"password" db:"password"`
	Fullname string `json:"fullname" db:"fullname"`
	Email string `json:"email" db:"email"`
	IsActive bool `json:"isActive" db:"is_active"`
	CreatedDate string `json:"createdDate" db:"created_date"`
	UpdatedDate sql.NullString`json:"updatedDate" db:"updated_date"`
}


type UserPermission struct {
	UserId int64 `json:"userId" db:"user_id"`
	Permission string `json:"permission" db:"permission"`
}
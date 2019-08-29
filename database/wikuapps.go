package database

import (
	"database/sql"
)

type WikuappsSession struct {
	Id int64 `json:"id" db:"id"`
	WikuappsUserId int64 `json:"userId" db:"wikuapps_user_id"`
	SessionId string `json:"sessionId" db:"session_id"`
	CreateTime sql.NullString `json:"createTime" db:"create_time"`
	UpdateTime sql.NullString`json:"updateTime" db:"update_time"`
}

type WikuappsUser struct {
	Id int64 `json:"id" db:"id"`
	Fullname string `json:"fullname" db:"fullname"`
	Email string `json:"email" db:"email"`
	Phone string `json:"phone" db:"phone"`
	Generation int64`json:"generation" db:"generation"`
	GraduateYear int64 `json:"graduateYear" db:"graduate_year"`
	Domicile string `json:"domicile" db:"domicile"`
	Occupation string`json:"occupation" db:"occupation"`
	Password string  `json:"password" db:"password"`
	CreateTime sql.NullString `json:"createTime" db:"create_time"`
	UpdateTime sql.NullString`json:"updateTime" db:"update_time"`
}
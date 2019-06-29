package database

import (
	_ "database/sql"
	wikuxConfig "wikux-api/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Wordpress *sqlx.DB
	Wikufest *sqlx.DB
)

func InitDB() {
	Wordpress = sqlx.MustConnect("mysql", wikuxConfig.DatabaseWordpressUrl)
	Wikufest = sqlx.MustConnect("mysql", wikuxConfig.DatabaseWikufestUrl)

	Wordpress.SetMaxOpenConns(wikuxConfig.DatabaseMaxOpenConns)
	Wordpress.SetMaxIdleConns(wikuxConfig.DatabaseMaxIdleConns)

	Wikufest.SetMaxOpenConns(wikuxConfig.DatabaseMaxOpenConns)
	Wikufest.SetMaxIdleConns(wikuxConfig.DatabaseMaxIdleConns)
}
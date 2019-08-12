package config

import (
	"os"
	"fmt"
	"strconv"
)

var (
	DatabaseUser = os.Getenv("WIKUX_DB_USER")
	DatabasePassword = os.Getenv("WIKUX_DB_PASSWORD")
	DatabaseHost = os.Getenv("WIKUX_DB_HOST")
	DatabaseWordpressUrl = fmt.Sprintf("%s:%s@tcp(%s)/wikusama_wp_db?&autocommit=true", DatabaseUser, DatabasePassword, DatabaseHost)
	DatabaseWikufestUrl = fmt.Sprintf("%s:%s@tcp(%s)/wikufest_db?&autocommit=true", DatabaseUser, DatabasePassword, DatabaseHost)
	DatabaseMaxOpenConns, _ = strconv.Atoi(os.Getenv("WIKUX_DB_MAXOPENCONS"))
	DatabaseMaxIdleConns, _ = strconv.Atoi(os.Getenv("WIKUX_DB_MAXIDLECONS"))

	ServiceGmailSmtp = os.Getenv("WIKUX_SMTP_HOST")
	ServiceGmailUser = os.Getenv("WIKUX_SMTP_USER")
	ServiceGmailPassword = os.Getenv("WIKUX_SMTP_PASSWORD")

	MailgunDomain = os.Getenv("WIKUX_MAILGUN_DOMAIN")
	MailgunPrivateKey = os.Getenv("WIKUX_MAILGUN_PRIVATE_KEY")

	HttpPort = os.Getenv("WIKUX_HTTP_PORT")
)
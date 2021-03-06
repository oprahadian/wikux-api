package handler

import (
	"log"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"fmt"
    "net/mail"
	"net/smtp"
    "crypto/tls"
	"wikux-api/config"
)

type SingleSendEmail struct {
	To string `json:"to"`
	Subject string `json:"subject"`
	Body string `json:"body"`
}

func ServiceSendEmail(c *gin.Context){
	var se SingleSendEmail
	var err error

	if err = c.BindJSON(&se);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": err,
		})

		return
	}

	go func(se SingleSendEmail) {
		from := mail.Address{"Publikasi Wikusama", config.ServiceGmailUser}
		to   := mail.Address{"", se.To}
		subj := se.Subject
		body := se.Body
	
		headers := make(map[string]string)
		headers["From"] = from.String()
		headers["To"] = to.String()
		headers["Subject"] = subj
		headers["MIME-Version"] = "1.0"
		headers["Content-Type"] = "text/html; charset=\"utf-8\""
		headers["Content-Transfer-Encoding"] = "base64"
	
		message := ""
		for k,v := range headers {
			message += fmt.Sprintf("%s: %s\r\n", k, v)
		}
		message += "\r\n" + body
	
		servername := config.ServiceGmailSmtp
	
		host, _, _ := net.SplitHostPort(servername)
	
		auth := smtp.PlainAuth("", config.ServiceGmailUser, config.ServiceGmailPassword, host)
	
		tlsconfig := &tls.Config {
			InsecureSkipVerify: true,
			ServerName: host,
		}
	
		conn, err := tls.Dial("tcp", servername, tlsconfig)
		if err != nil {
			log.Print(err)
		}
	
		c, err := smtp.NewClient(conn, host)
		if err != nil {
			log.Print(err)
		}
	
		if err = c.Auth(auth); err != nil {
			log.Print(err)
		}
	
		if err = c.Mail(from.Address); err != nil {
			log.Print(err)
		}
	
		if err = c.Rcpt(to.Address); err != nil {
			log.Print(err)
		}
	
		w, err := c.Data()
		if err != nil {
			log.Print(err)
		}
	
		_, err = w.Write([]byte(message))
		if err != nil {
			log.Print(err)
		}
	
		err = w.Close()
		if err != nil {
			log.Print(err)
		}
	
		c.Quit()
	}(se)
	
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": nil,
		"data": SingleSendEmail{To: se.To, Subject: se.Subject},
	})
}
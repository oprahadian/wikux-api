package common

import (
	"log"
	"net"
	"fmt"
	"sync"
	"strings"
    "net/mail"
	"net/smtp"
    "crypto/tls"
	"wikux-api/config"
)

type CommonSingleSendEmail struct {
	To string
	Cc []string
	Bcc []string
	Subject string
	Body string	
}

func CommonSendEmailConcurrent(se CommonSingleSendEmail, wg *sync.WaitGroup){
	CommonSendEmail(se)
	wg.Done()
}

func CommonSendEmail(se CommonSingleSendEmail){
	from := mail.Address{"Publikasi Wikusama", config.ServiceGmailUser}
	to   := mail.Address{"", se.To}
	subj := se.Subject
	body := se.Body

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	
	if len(se.Cc) > 0 {
		headers["Cc"] = strings.Join(se.Cc, ",") 
	}
	
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

	for _, v := range se.Cc {
		if err = c.Rcpt(v); err != nil {
			log.Print(err)
		}
	}

	for _, v := range se.Bcc {
		if err = c.Rcpt(v); err != nil {
			log.Print(err)
		}
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
}
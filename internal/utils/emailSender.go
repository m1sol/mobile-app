package utils

import (
	"fmt"
	"net/smtp"
)

type EmailSender struct {
	SMTPServer string
	Port       string
	Username   string
	Password   string
}

func NewEmailSender(smtpServer string, port string, username, password string) *EmailSender {
	return &EmailSender{
		SMTPServer: smtpServer,
		Port:       port,
		Username:   username,
		Password:   password,
	}
}

func (es *EmailSender) SendEmail(to, subject, body string) error {
	from := es.Username
	auth := smtp.PlainAuth("", es.Username, es.Password, es.SMTPServer)

	msg := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", from, to, subject, body))

	addr := fmt.Sprintf("%s:%s", es.SMTPServer, es.Port)
	return smtp.SendMail(addr, auth, from, []string{to}, msg)
}

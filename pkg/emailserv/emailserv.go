package emailserv

import (
	"fmt"
	"net/smtp"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
)

// Send the email to the user with the body received as a paramater
func Send(email entity.Email) error {
	from := "banco.uati.squad.4@gmail.com"
	pass := "bancouatisquad4"

	to := email.To
	subject := email.Subject
	body := email.Message

	// TODO: Deixar esta mensagem mais bonitinha
	msg := fmt.Sprintf("From: "+from+"\n"+
		"To: "+to+"\n"+
		"Subject: "+subject+"\n",
		body)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		panic(err)
	}

	return nil
}

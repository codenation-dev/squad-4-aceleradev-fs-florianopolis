package emailserv

import (
	"fmt"
	"net/smtp"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/entity"
)

// Send the email to the user with the body received as a paramater
func Send(body interface{}, user entity.User) error {
	from := "banco.uati.squad.4@gmail.com"
	pass := "bancouatisquad4"

	to := user.Email

	subject := "Novo Cliente"

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

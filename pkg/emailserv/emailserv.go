package emailserv

import (
	"fmt"
	"net/smtp"
)

// Send the email to the user with the body received as a paramater
func Send(msg interface{}, to ...string) error {
	from := "banco.uati.squad.4@gmail.com"
	pass := "bancouatisquad4"

	strTo := fmt.Sprintf("%v", to)
	subject := "Email automático Banco UAti"
	body := fmt.Sprintf("%v", msg)

	payload := fmt.Sprintf("From: "+from+"\n"+
		"To: "+strTo+"\n"+
		"Subject: "+subject+"\n\n",
		body)

	// payload := fmt.Sprintf(
	// 	`From: %s\n
	// 	 To: %v\n
	// 	 Subject: Email automático Banco UAti\n`, from, to, msg)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, to, []byte(payload))

	if err != nil {
		return err
	}

	return nil
}

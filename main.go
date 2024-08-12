package gomail

import (
	"github.com/go-mail/mail"
)

func main() {

	m := mail.NewMessage()

	m.SetHeader("From", "john.doe@gmail.com")

	m.SetHeader("To", "kate.doe@example.com", "noah.doe@example.com")

	m.SetAddressHeader("Cc", "oliver.doe@example.com", "Oliver")

	m.SetHeader("Subject", "Hello!")

	m.SetBody("text/html", "Hello <b>Kate</b> and <i>Noah</i>!")

	m.Attach("lolcat.jpg")

	d := mail.NewDialer("smtp.gmail.com", 587, "john.doe@gmail.com", "123456")

	// Send the email to Kate, Noah and Oliver.

	if err := d.DialAndSend(m); err != nil {

		panic(err)

	}

}

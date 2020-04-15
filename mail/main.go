package main

import (
	"log"

	"github.com/alexcesaro/mail/gomail"
)

func main() {
	msg := gomail.NewMessage()
	msg.SetAddressHeader("From", "postgre.sql8@msa.hinet.net", "Bunko")
	msg.SetHeader("To", "hitomitanaka666@gmail.com")
	msg.AddHeader("To", "bunko666@gmail.com")
	msg.SetHeader("Subject", "Hello!")
	msg.SetBody("text/plain", "Hello Bunko!")
	msg.AddAlternative("text/html", "Hello <b>Bunko</b>!")
	// if err := msg.Attach("p1.jpg"); err != nil {
	// 	log.Println(err)
	// 	return
	// }

	m := gomail.NewMailer("msa.hinet.net", "postgre.sql8", "yourpasswd", 25)
	if err := m.Send(msg); err != nil {
		log.Println(err)
	}
}

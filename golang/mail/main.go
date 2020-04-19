package main

import (
	"fmt"
	"net/smtp"
)

// 以下 variable 可參考 Gmail 的 smtp 設定說明
var (
	host     = "smtp.gmail.com:587"
	username = "ltscyt0717@gmail.com"
	password = "LTS0717cyt"
)

func main() {
	auth := smtp.PlainAuth(host, username, password, "smtp.gmail.com")

	to := []string{"gfes980615@yahoo.com.tw"}
	msg := []byte(
		"Subject: This is a test mail!\r\n" +
			"From: gfes980615@yahoo.com.tw\r\n" +
			`Content-Type: multipart/mixed; boundary="qwertyuio"` + "\r\n" +
			"\r\n" +
			"--qwertyuio\r\n" +
			"This is the body of email.\r\n" +
			"\r\n" +
			"--qwertyuio--\r\n",
	)
	err := smtp.SendMail(
		host,
		auth,
		username,
		to,
		msg,
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("success !")
}

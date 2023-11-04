package main

import (
	"fmt"

	 gomail "gopkg.in/gomail.v2"

)

func main() {
	abc := gomail.NewMessage()

	abc.SetHeader("From", "ajengnikita14@gmail.com")
	abc.SetHeader("To", "ajengnktaa15@gmail.com")
	abc.SetHeader("Subject", "HAI INI ADA YANG BARU")
	abc.SetBody("text/plain", "THIS IS BODY")

	a := gomail.NewDialer("smtp.gmail.com", 587, "ajengnikita14@gmail.com", "fird jdwa rujm xlyq")

	if err := a.DialAndSend(abc); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

// func sendMailSimple() {
// 	auth := smtp.PlainAuth(
// 		"",
// 		"ajengnikita14@gmail.com",
// 		"ajeng1056",
// 		"smtp.gmail.com",
// 	)

// 	// pw := "fird jdwa rujm xlyq"

// 	msg := "Subject: My special subject\nThis is the body of the email."

// 	err := smtp.SendMail(
// 		"smtp.gmail.com:587",
// 		auth,
// 		"ajengnikita14@gmail.com",
// 		[]string{"ajengnikita14@gmail.com"},
// 		[]byte(msg),
// 	)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func main() {
// 	sendMailSimple()
// }

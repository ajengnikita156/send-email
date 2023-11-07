package main

import (
	"fmt"

	 gomail "gopkg.in/gomail.v2"

)

func main() {
	send := gomail.NewMessage()

	send.SetHeader("From", "ajengnikita14@gmail.com")
	send.SetHeader("To", "ajengnktaa15@gmail.com")
	send.SetHeader("Subject", "HAI KAMU")
	send.SetBody("text/plain", "THIS IS BODY jbajbancahieaknab jbjaja")

	a := gomail.NewDialer("smtp.gmail.com", 587, "ajengnikita14@gmail.com", "fird jdwa rujm xlyq")
	fmt.Println("Sudah Terkirim!")
	if err := a.DialAndSend(send); err != nil {
		
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

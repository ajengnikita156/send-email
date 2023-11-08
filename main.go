package main

import (
	"fmt"

	 gomail "gopkg.in/gomail.v2"

)

func main() {
	send := gomail.NewMessage()

	send.SetHeader("From", "ajengnikita14@gmail.com")
	send.SetHeader("To", "ajengnktaa15@gmail.com")
	send.SetHeader("Subject", "selamat...")
	send.SetBody("text/html", 
	`<!DOCTYPE html>
	<html>
	<head>
		<title>Job Acceptance</title>
		<style>
			body {
				font-family: Arial, sans-serif;
			}
	
			.container {
				max-width: 400px;
				margin: 0 auto;
				padding: 20px;
			}
	
			h2 {
				color: #007BFF;
			}
	
			p {
				margin-top: 10px;
			}
	
			.message {
				background-color: #f0f0f0;
				border: 1px solid #ccc;
				padding: 10px;
				margin-top: 20px;
			}
	
			.signature {
				margin-top: 20px;
				font-weight: bold;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h2>Job Acceptance</h2>
			<p>Dear AJENG NIKITA ANGGRAENI ,</p>
			<p>We are pleased to inform you that you have been accepted for the position of <strong>JOB TITTLE</strong> at our company. Your employment will begin on <strong>10-11-2023</strong>.</p>
	
			<div class="message">
				<p>Congratulations on your new role, and we look forward to having you as part of our team!</p>
			</div>
	
			<p class="signature">Sincerely,<br>PT. PERUSAHAN</p>
		</div>
	</body>
	</html>`)

	a := gomail.NewDialer("smtp.gmail.com", 587, "ajengnikita14@gmail.com", "fird jdwa rujm xlyq")
	fmt.Println("Sudah Terkirim!")
	if err := a.DialAndSend(send); err != nil {
		
		fmt.Println(err)
		panic(err)
	}
}

// 	// pw := "fird jdwa rujm xlyq"
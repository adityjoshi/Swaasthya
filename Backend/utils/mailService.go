package utils

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func OtpRegistration(to, otp string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", "aditya3.collegeboard@gmail.com")
	message.SetHeader("To", to, "aditya30joshi@gmail.com")
	message.SetHeader("Subject", "Otp"+otp)

	// Construct the email body with dynamic complaint details
	body := fmt.Sprintf("Dear student, thank you for logging in to the hostel ease. If it's not you reach out to us asap through hostelvit@gmail.com\n\n")
	body += fmt.Sprintf("*Hostel Team*\n")
	body += fmt.Sprintf("*Block 4*")
	message.SetBody("text/plain", body)

	//message.Attach("/home/Alex/lolcat.jpg")

	// Initialize SMTP dialer
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "aditya3.collegeboard@gmail.com", "ehnxaubjqelkotks") // Update with your SMTP server details

	// Send email
	if err := dialer.DialAndSend(message); err != nil {
		panic(err)
	}

	fmt.Println("Email sent successfully!")
	return nil
}

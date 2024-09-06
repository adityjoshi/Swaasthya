package utils

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func OtpRegistration(to, otp string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", "aditya3.collegeboard@gmail.com")
	message.SetHeader("To", to, "aditya30joshi@gmail.com")
	message.SetHeader("Subject", "Otp Verification")

	htmlBody := `
    <html>
    <body>
        <h1>Your OTP Code</h1>
        <p>Dear User,</p>
        <p>Your One-Time Password (OTP) is <strong>` + otp + `</strong>.</p>
        <p>Please use this OTP to complete your verification.</p>
        <p>If you did not request this OTP, please ignore this email.</p>
        <p>Best regards,<br>Swaasthya</p>
    </body>
    </html>
    `
	body := fmt.Sprintf("htmlBody")
	body += fmt.Sprintf("*Best regards*\n")
	body += fmt.Sprintf("*Team Swaasthaya*")
	message.SetBody("text/html", htmlBody)

	//message.Attach("/home/Alex/lolcat.jpg")

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "aditya3.collegeboard@gmail.com", "ehnxaubjqelkotks") // Update with your SMTP server details

	// Send email
	if err := dialer.DialAndSend(message); err != nil {
		panic(err)
	}

	fmt.Println("Email sent successfully!")
	return nil
}
func SendAppointmentEmail(patientEmail, doctorName, appointmentDate, appointmentTime, bookingTime string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", "aditya3.collegeboard@gmail.com")
	message.SetHeader("To", patientEmail)
	message.SetHeader("Subject", "Appointment Confirmation")

	htmlBody := `
    <html>
    <body>
        <h1>Appointment Confirmation</h1>
        <p>Dear Patient,</p>
        <p>Your appointment with Dr. ` + doctorName + ` has been successfully booked.</p>
        <p><strong>Appointment Date:</strong> ` + appointmentDate + `</p>
        <p><strong>Appointment Time:</strong> ` + appointmentTime + `</p>
        <p><strong>Booking Time:</strong> ` + bookingTime + `</p>
        <p>If you have any questions, please contact us.</p>
        <p>Best regards,<br>Swaasthya</p>
    </body>
    </html>
    `
	message.SetBody("text/html", htmlBody)

	// Initialize SMTP dialer
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "aditya3.collegeboard@gmail.com", "ehnxaubjqelkotks") // Update with your SMTP server details

	// Send email
	if err := dialer.DialAndSend(message); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Println("Email sent successfully!")
	return nil
}

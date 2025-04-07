package services

import (
	"fmt"
	"os"

	"strconv"

	"gopkg.in/gomail.v2"
)

func SendOTPEmail(email string, otp string) error {
	// SMTP server configuration
	smtpHost := os.Getenv("SMTP_HOST")     // Replace with your SMTP server
	smtpPort := os.Getenv("SMTP_PORT")     // Replace with your SMTP port
	smtpUser := os.Getenv("SMTP_USER")     // Replace with your email
	smtpPass := os.Getenv("SMTP_PASSWORD") // Replace with your email password
	smtpFrom := os.Getenv("SMTP_FROM")     // Replace with your email password

	// Create a new email message
	message := gomail.NewMessage()
	message.SetHeader("From", smtpFrom)
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Your OTP Code")
	message.SetBody("text/plain", fmt.Sprintf("Your OTP code is: %s", otp))

	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		return fmt.Errorf("invalid SMTP port: %w", err)
	}
	dialer := gomail.NewDialer(smtpHost, port, smtpUser, smtpPass)

	// Send the email
	if err := dialer.DialAndSend(message); err != nil {
		return fmt.Errorf("failed to send OTP email: %w", err)
	}

	return nil
}

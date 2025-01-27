package notification

import (
	"crypto/tls"
	"fmt"
	"strconv"

	"github.com/hudaputrasantosa/auth-users-api/internal/config"
	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

func SendEmailNotification(recipient *RecipientInformation) error {
	senderEmail := config.Config("MAILSENDER_SENDER_EMAIL")
	server := config.Config("MAILSENDER_SMTP_SERVER")
	port, _ := strconv.Atoi(config.Config("MAILSENDER_SMTP_PORT"))
	username := config.Config("MAILSENDER_SMTP_USERNAME")
	password := config.Config("MAILSENDER_SMTP_PASSWORD")

	initMessage := gomail.NewMessage()
	// Set Header Message
	initMessage.SetHeader("From", senderEmail)
	// initMessage.SetHeader("To", "abc@gmail.com", "xyz@gmail.com", "123@gmail.com")  // Multiple recipients
	initMessage.SetHeader("To", recipient.Email)
	initMessage.SetHeader("Subject", "Notification with Gomail")

	// initMessage.SetBody("text/plain", "Test Email message")
	initMessage.AddAlternative("text/html", `
        <html>
            <body>
                <h1>Thanks for register!</h1>
                <p>Hello!This is otp for verification.</p>
                <p><b>12345</b></p>
            </body>
        </html>
    `)

	// Add attachments
	// initMessage.Attach("/invoice#1.pdf")

	// Embed image and set email body to reference the embedded image
	// initMessage.Embed("/path/to/image.jpg")

	// Setup SMTP Dialer
	dialer := gomail.NewDialer(server, port, username, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email
	if err := dialer.DialAndSend(initMessage); err != nil {
		logger.Error("Failed to send email", zap.Error(err))
		return err
	}

	fmt.Println("Email sent successfully!")
	return nil
}

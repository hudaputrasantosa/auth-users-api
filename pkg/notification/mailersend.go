package notification

import (
	"context"
	"time"

	"github.com/hudaputrasantosa/auth-users-api/internal/config"
	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"
	mailersend "github.com/mailersend/mailersend-go"
	"go.uber.org/zap"
)

type RecipientInformation struct {
	Name            string
	Email           string
	MessageTemplate string
}

func MailersendNotification(recipient *RecipientInformation) (any, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	key := config.Config("MAILSENDER_API_KEY")
	senderName := config.Config("MAILSENDER_SENDER_NAME")
	senderEmail := config.Config("MAILSENDER_SENDER_EMAIL")

	from := mailersend.From{
		Name:  senderName,
		Email: senderEmail,
	}

	recipients := []mailersend.Recipient{
		{
			Name:  recipient.Name,
			Email: recipient.Email,
		},
	}

	// attachment : []mailersend.Attachment{
	// 	{
	// 		Content: "",
	// 		Filename: "",
	// 	},
	// }

	clientMailerSend := mailersend.NewMailersend(key)
	subject := "Test email"
	text := "This is the plain text version of the email."
	html := "<h1>Hello {{name}},</h1><p>Thanks for register! this is OTP 12345</p>"

	message := clientMailerSend.Email.NewMessage()
	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)
	// message.Attachments(attachment)

	res, err := clientMailerSend.Email.Send(ctx, message)
	if err != nil {
		logger.Error("Failed to send email", zap.Error(err))
		return nil, err
	}

	logger.Info("Success to send email")
	return res, nil
}

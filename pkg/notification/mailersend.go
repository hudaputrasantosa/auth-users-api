package notification

import (
	"context"
	"time"

	"github.com/hudaputrasantosa/auth-users-api/internal/config"
	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"
	"github.com/hudaputrasantosa/auth-users-api/pkg/utils/templates"
	mailersend "github.com/mailersend/mailersend-go"
	"go.uber.org/zap"
)

type RecipientInformation struct {
	Name  string
	Email string
}

func MailersendNotification(recipient *RecipientInformation, data *templates.DataBodyInformation) (any, error) {
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

	clientMailerSend := mailersend.NewMailersend(key)
	template, err := templates.ParseFile(data)
	if err != nil {
		logger.Error("Failed to parse template", zap.Error(err))
		return nil, err
	}

	// text := "This is the plain text version of the email."
	message := clientMailerSend.Email.NewMessage()
	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(template.Subject)
	message.SetHTML(template.Body)
	// message.SetText(text)

	res, err := clientMailerSend.Email.Send(ctx, message)
	if err != nil {
		logger.Error("Failed to send email", zap.Error(err))
		return nil, err
	}

	logger.Info("Success to send email")
	return res, nil
}

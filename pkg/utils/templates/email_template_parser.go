package templates

import (
	"bytes"
	"fmt"
	"html/template"
)

type DataEmailTemplate struct {
	Subject string
	Body    string
}

type DataBodyInformation struct {
	Name            string
	Email           string
	Otp             string
	Attachment      any
	MessageTemplate string
}

func ParseFile(data *DataBodyInformation) (*DataEmailTemplate, error) {
	var body bytes.Buffer
	var subject string

	// parse email template from html
	templ, err := template.New("email").Parse(data.MessageTemplate)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return nil, err
	}

	// Execute the template with dynamic data
	switch data.MessageTemplate {
	case Otp_template:
		subject = "OTP Register Verification"
		err = templ.Execute(&body, &DataBodyInformation{Name: data.Name, Otp: data.Otp})
	case Reset_password_template:
		subject = "OTP Reset Password"
		err = templ.Execute(&body, &DataBodyInformation{Name: data.Name, Otp: data.Otp})
	default:
		return nil, fmt.Errorf("template not found")
	}

	if err != nil {
		fmt.Println("error executing template:", err)
		return nil, err
	}

	return &DataEmailTemplate{
		Subject: subject,
		Body:    body.String(),
	}, nil
}

package configs

import (
	"bytes"
	"html/template"

	"github.com/spf13/viper"
	"github.com/wneessen/go-mail"
)

type MailerService struct {
	ViperConfig *viper.Viper
	Sender      string
}

func NewMailerService(viperConfig *viper.Viper) *MailerService {
	return &MailerService{ViperConfig: viperConfig, Sender: viperConfig.GetString("EMAIL_USERNAME")}
}

type EmailPayload struct {
	Title     string
	Recipient string
	Body      string
	Sender    string
}

// SendEmail sends an email with the provided data

func (mailerService *MailerService) SendEmail(recipientEmail string, subjectEmail string, templateFile string, data EmailPayload) error {
	// Parse the HTML template

	templateFilePath, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}
	// Render the template with the data

	var body bytes.Buffer

	if err := templateFilePath.Execute(&body, data); err != nil {
		return err
	}
	// Create a new email message
	mailerMessage := mail.NewMsg()
	err = mailerMessage.From(mailerService.ViperConfig.GetString("EMAIL_USERNAME"))
	if err != nil {
		return err
	} // Sender email

	err = mailerMessage.To(recipientEmail)

	if err != nil {
	} // Recipient email

	mailerMessage.Subject(subjectEmail)                           // Subject
	mailerMessage.SetBodyString(mail.TypeTextHTML, body.String()) // Email body (HTML)

	// Set up SMTP client
	client, err := mail.NewClient("smtp.gmail.com",
		mail.WithPort(587),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(mailerService.ViperConfig.GetString("EMAIL_USERNAME")),
		mail.WithPassword(mailerService.ViperConfig.GetString("EMAIL_PASSWORD")), // Use App Password if using Gmail
		mail.WithTLSPolicy(mail.DefaultTLSPolicy),
	)

	if err != nil {

		return err
	}

	// Send the email
	if err := client.DialAndSend(mailerMessage); err != nil {

		return err
	}

	return nil
}

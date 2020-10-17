package commons

import (
	"email/config"

	"github.com/adityak368/swissknife/email"
	"github.com/adityak368/swissknife/email/knifemailer"
)

var mailer email.Mailer

// Mailer returns a new emailer
func Mailer() email.Mailer {
	if mailer == nil {
		mailer = knifemailer.New(email.MailerConfig{
			Host:     config.EmailHost,
			Port:     config.EmailPort,
			Username: config.EmailUsername,
			Password: config.EmailPassword,
		})
	}

	return mailer
}

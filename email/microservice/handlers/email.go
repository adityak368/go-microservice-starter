package handlers

import (
	"context"
	"email/commons"
	"email/config"

	"commons/proto/email"

	"github.com/adityak368/swissknife/logger"
)

type Email struct{}

func (e *Email) SendEmail(ctx context.Context, req *email.SendEmailRequest, rsp *email.SendEmailResponse) error {
	mailer := commons.Mailer()
	mailer.SendMail(config.FromEmailID, req.To, req.Subject, req.Body)
	logger.Info.Printf("Sent Email to %s", req.To)
	return nil
}

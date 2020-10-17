package subscribers

import (
	"context"
	"email/commons"
	"email/config"

	"commons/proto/auth"

	"github.com/adityak368/swissknife/logger"
)

type UserCreated struct{}

func (a *UserCreated) OnUserCreated(ctx context.Context, msg *auth.UserCreated) error {
	mailer := commons.Mailer()
	mailer.SendMail(config.FromEmailID, msg.UserData.Email, "Welcome", "Welcome")
	logger.Info.Printf("Sent Welcome Email to %s", msg.UserData.Email)
	return nil
}

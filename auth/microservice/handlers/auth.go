package handlers

import (
	"auth/internal"
	"auth/internal/models"
	"context"

	"commons/proto/auth"
)

type Auth struct{}

func (e *Auth) GetUserByEmail(ctx context.Context, req *auth.GetUserByEmailRequest, rsp *auth.User) error {

	email := req.Email
	res, err := internal.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	user := res.Data().(*models.User)

	rsp = user.ToProtobuf()

	return nil
}

func (e *Auth) GetUserById(ctx context.Context, req *auth.GetUserByIdRequest, rsp *auth.User) error {

	userID := req.UserId
	res, err := internal.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	user := res.Data().(*models.User)

	rsp = user.ToProtobuf()

	return nil
}

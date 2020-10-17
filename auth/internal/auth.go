package internal

import (
	"auth/commons"
	"auth/commons/jwt"
	"auth/config"
	"auth/db"
	"auth/internal/models"
	"auth/microservice/publishers"
	"commons/proto/auth"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/adityak368/swissknife/logger"
	"github.com/adityak368/swissknife/response"
	"github.com/adityak368/swissknife/validation/playground"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(c context.Context, r *LoginRequest) (response.Result, error) {

	if err := playground.ValidateStruct(r); err != nil {
		return nil, response.NewError(http.StatusBadRequest, err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.DbTimeout)*time.Second)
	defer cancel()

	user := &models.User{}
	if err := db.GetMongoConnection().Collection("User").FindOne(ctx, bson.D{{"email", r.Email}}).Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, response.NewError(http.StatusBadRequest, "EmailNotFound")
		}
		logger.Error.Println(err)
		return nil, response.NewError(http.StatusInternalServerError, "ErrorAtServer")
	}

	if user.Authenticate(r.Password) {

		t, err := jwt.GenerateAuthToken(user, config.UserSecret)
		if err != nil {
			logger.Error.Println(err)
			return nil, response.NewError(http.StatusInternalServerError, "ErrorAtServer")
		}
		return &response.ExecResult{Result: commons.Token{AuthToken: t}}, nil
	}

	return nil, response.NewError(http.StatusBadRequest, "InvalidEmailPassword")
}

func Signup(c context.Context, r *SignUpRequest) (response.Result, error) {

	if err := playground.ValidateStruct(r); err != nil {
		return nil, response.NewError(http.StatusBadRequest, err.Error())
	}

	user := &models.User{
		ID:        primitive.NewObjectID(),
		Email:     r.Email,
		Name:      r.Name,
		Password:  r.Password,
		CreatedAt: time.Now(),
	}

	user.HashPassword()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.DbTimeout)*time.Second)
	defer cancel()

	if _, err := db.GetMongoConnection().Collection("User").InsertOne(ctx, user); err != nil {
		switch ex := err.(type) {
		case mongo.WriteException:
			for _, writeErr := range ex.WriteErrors {
				if writeErr.Code == 11000 {
					return nil, response.NewError(http.StatusBadRequest, "EmailAlreadyExists")
				}
			}
		default:
		}
		logger.Error.Println(err)
		return nil, response.NewError(http.StatusInternalServerError, "ErrorAtServer")
	}

	userData := user.ToProtobuf()
	// chain context to maintain trace
	publishers.GetPublishers().UserCreated.Publish(c, &auth.UserCreated{UserData: userData})

	return &response.ExecResult{MessageID: "UserCreated"}, nil
}

func EditProfile(c context.Context, r *EditProfileRequest, user *models.User) (response.Result, error) {

	if err := playground.ValidateStruct(r); err != nil {
		return nil, response.NewError(http.StatusBadRequest, err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.DbTimeout)*time.Second)
	defer cancel()

	fmt.Println(user)

	update := bson.M{"$set": bson.M{
		"name":     r.Name,
		"headline": r.Headline,
		"contact":  r.Contact,
	}}

	_, err := db.GetMongoConnection().Collection("User").UpdateOne(ctx, bson.D{{"_id", user.ID}}, update)
	if err != nil {
		logger.Error.Println(err)
		return nil, response.NewError(http.StatusInternalServerError, "ErrorAtServer")
	}

	user.Name = r.Name
	user.Headline = r.Headline
	user.Contact = r.Contact
	userData := user.ToProtobuf()
	// chain context to maintain trace
	publishers.GetPublishers().UserUpdated.Publish(c, &auth.UserUpdated{UserData: userData})

	// Generate encoded token and send it as response.
	t, err := jwt.GenerateAuthToken(user, config.UserSecret)
	if err != nil {
		logger.Error.Println(err)
		return nil, response.NewError(http.StatusInternalServerError, "ErrorAtServer")
	}
	return &response.ExecResult{Result: commons.Token{AuthToken: t}}, nil
}

func UploadAvatar(c context.Context, r *UploadAvatarRequest, user *models.User) (response.Result, error) {

	if err := playground.ValidateStruct(r); err != nil {
		return nil, response.NewError(http.StatusBadRequest, err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.DbTimeout)*time.Second)
	defer cancel()

	update := bson.M{"$set": bson.M{
		"avatarUrl": r.AvatarURL,
	}}

	_, err := db.GetMongoConnection().Collection("User").UpdateOne(ctx, bson.D{{"_id", user.ID}}, update)
	if err != nil {
		logger.Error.Println(err)
		return nil, response.NewError(http.StatusInternalServerError, "ErrorAtServer")
	}

	user.AvatarURL = r.AvatarURL

	// Generate encoded token and send it as response.
	t, err := jwt.GenerateAuthToken(user, config.UserSecret)
	if err != nil {
		logger.Error.Println(err)
		return nil, response.NewError(http.StatusInternalServerError, "ErrorAtServer")
	}
	return &response.ExecResult{Result: commons.Token{AuthToken: t}}, nil
}

func GetUserByEmail(c context.Context, email string) (response.Result, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.DbTimeout)*time.Second)
	defer cancel()

	user := &models.User{}
	if err := db.GetMongoConnection().Collection("User").FindOne(ctx, bson.D{{"email", email}}).Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, response.NewError(http.StatusBadRequest, "EmailNotFound")
		}
		logger.Error.Println(err)
		return nil, response.NewError(http.StatusInternalServerError, "ErrorAtServer")
	}
	return &response.ExecResult{Result: user}, nil
}

func GetUserByID(c context.Context, userID string) (response.Result, error) {

	ID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, response.NewError(http.StatusBadRequest, "InvalidUserID")
	}

	user := &models.User{}
	if err := db.GetMongoConnection().Collection("User").FindOne(c, bson.D{{"_id", ID}}).Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, response.NewError(http.StatusBadRequest, "UserNotFound")
		}
		logger.Error.Println(err)
		return nil, response.NewError(http.StatusInternalServerError, "ErrorAtServer")
	}
	return &response.ExecResult{Result: user}, nil
}

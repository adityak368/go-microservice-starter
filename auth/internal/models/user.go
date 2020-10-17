package models

import (
	"time"

	"commons/proto/auth"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// User -- User Model
type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Headline  string             `bson:"headline,omitempty" json:"headline,omitempty"`
	Contact   string             `bson:"contact,omitempty" json:"contact,omitempty"`
	Email     string             `bson:"email" json:"email" validate:"required,email"`
	AvatarURL string             `bson:"avatarUrl,omitempty" json:"avatarUrl,omitempty"`
	Password  string             `bson:"password" json:"-" validate:"required,min=8,max=16"`
	CreatedAt time.Time          `bson:"createdAt" json:"-"`
}

//HashPassword -- Hash the password and save it
func (u *User) HashPassword() error {

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Authenticate -- Authenticate User with the entered Password
func (u *User) Authenticate(password string) bool {

	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// ToProtobuf Serializes to protobuf
func (u *User) ToProtobuf() *auth.User {
	return &auth.User{
		ID:        u.ID.Hex(),
		Name:      u.Name,
		Headline:  u.Headline,
		Contact:   u.Contact,
		Email:     u.Email,
		AvatarURL: u.AvatarURL,
		CreatedAt: u.CreatedAt.Unix(),
	}
}

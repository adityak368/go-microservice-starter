package internal

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignUpRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=16"`
	Name     string `json:"name" validate:"required"`
}

type EditProfileRequest struct {
	Name     string `json:"name"`
	Headline string `json:"headline"`
	Contact  string `json:"contact"`
}

type UploadAvatarRequest struct {
	AvatarURL string `json:"avatarUrl" validate:"required"`
}

package commons

type Token struct {
	AuthToken string `json:"authToken" validate:"required"`
}

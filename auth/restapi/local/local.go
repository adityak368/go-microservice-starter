package local

import (
	"auth/config"
	"auth/internal"
	"bufio"
	"math"
	"net/http"

	"auth/internal/models"

	authMiddleware "auth/commons/middleware"

	"github.com/adityak368/swissknife/logger"
	"github.com/adityak368/swissknife/objectstore/s3"
	"github.com/adityak368/swissknife/response"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Login(c echo.Context) error {
	r := new(internal.LoginRequest)
	if err := c.Bind(r); err != nil {
		return response.NewError(http.StatusBadRequest, "InvalidParams")
	}

	result, err := internal.Login(c.Request().Context(), r)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result.Data())
}

func Signup(c echo.Context) error {
	r := new(internal.SignUpRequest)
	if err := c.Bind(r); err != nil {
		return response.NewError(http.StatusBadRequest, "InvalidParams")
	}

	result, err := internal.Signup(c.Request().Context(), r)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result.Message())
}

func EditProfile(c echo.Context) error {

	r := new(internal.EditProfileRequest)
	if err := c.Bind(r); err != nil {
		return response.NewError(http.StatusBadRequest, "InvalidParams")
	}

	user := c.Get("user").(*models.User)

	result, err := internal.EditProfile(c.Request().Context(), r, user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result.Message())
}

func UploadAvatar(c echo.Context) error {

	avatar, err := c.FormFile("avatar")
	if err != nil {
		return response.NewError(http.StatusBadRequest, "InvalidParams")
	}
	if avatar.Size > 1024*1000 {
		return response.NewError(http.StatusBadRequest, "AvatarMaxSizeExceeded")
	}

	src, err := avatar.Open()
	defer src.Close()
	if err != nil {
		logger.Error.Println(err)
		return response.NewError(http.StatusBadRequest, "InvalidAvatar")
	}

	// Check for jpeg and png images
	bufReader := bufio.NewReader(src)
	buffer, err := bufReader.Peek(int(math.Min(512, float64(avatar.Size))))
	contentType := http.DetectContentType(buffer)
	if err != nil || !(contentType == "image/jpeg" || contentType == "image/png") {
		logger.Error.Println("Upload File Type: " + contentType)
		logger.Error.Println(err)
		return response.NewError(http.StatusBadRequest, "InvalidImage")
	}

	user := c.Get("user").(*models.User)

	uploadResult, err := s3.NewStore().AddImage(config.S3AvatarBucket, user.ID.Hex(), bufReader)
	if err != nil {
		logger.Error.Println(err)
		return response.NewError(http.StatusBadRequest, "CouldNotSaveAvatar")
	}

	r := &internal.UploadAvatarRequest{
		AvatarURL: uploadResult.Location,
	}

	result, err := internal.UploadAvatar(c.Request().Context(), r, user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result.Message())
}

// InitLocalSignIn -- Initialize The Local Sign In Route
func InitLocalSignIn(e *echo.Group) {
	// Login Route
	e.POST("/login", Login)
	// SignUp Route
	e.POST("/signup", Signup)
	// Profile Route
	editProfile := e.Group("/profile")
	editProfile.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(config.UserSecret),
		SigningMethod: "HS512",
	}))
	editProfile.Use(authMiddleware.EchoUserContextMiddleware())
	// Avatar Upload Route
	editProfile.POST("/upload", UploadAvatar)
	// Edit Profile Route
	editProfile.PUT("", EditProfile)
}

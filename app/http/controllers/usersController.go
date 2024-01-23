package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	response "github.com/rifansyah/go-crud/app/helpers"
	"github.com/rifansyah/go-crud/app/http/requests"
	"github.com/rifansyah/go-crud/app/initializers"
	"github.com/rifansyah/go-crud/app/models"
	"golang.org/x/crypto/bcrypt"
)

type UsersController struct{}

func NewUsersController() *UsersController {
	return &UsersController{}
}

func (ctl *UsersController) SignUp(c *gin.Context) {
	var b requests.Credentials
	if err := c.ShouldBind(&b); err != nil {
		response.ResponseError(c, http.StatusBadRequest, err, "Failed to read body")
		return
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(b.Password), 10)

	if err != nil {
		response.ResponseError(c, http.StatusInternalServerError, err, "Failed to hash the password")
		return
	}

	user := models.User{Email: b.Email, Password: string(passHash)}
	if result := initializers.DB.Create(&user); result.Error != nil {
		response.ResponseError(c, http.StatusInternalServerError, result.Error, "Failed to create the user")
		return
	} 

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (ctl *UsersController) Login(c *gin.Context) {
	var b requests.Credentials
	if err := c.ShouldBind(&b); err != nil {
		response.ResponseError(c, http.StatusBadRequest, err, "Failed to read body")
		return
	}

	var user models.User
	if result := initializers.DB.Where("email = ?", b.Email).First(&user); result.Error != nil {
		response.ResponseError(c, http.StatusBadRequest, result.Error, "User not found")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(b.Password)); err != nil {
		response.ResponseError(c, http.StatusBadRequest, err, "Password mismatch")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 5).Unix(),
	})

	signedToken, err := token.SignedString([]byte(os.Getenv("TOKEN_KEY")))
	if err != nil {
		response.ResponseError(c, http.StatusInternalServerError, err, "Failed to generate token")
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", signedToken, 3600 * 24 * 5, "", "", true, true)

	c.JSON(http.StatusOK, gin.H{
		"token": signedToken,
	})
}

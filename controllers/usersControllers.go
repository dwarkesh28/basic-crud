package controllers

import (
	"net/http"
	"os"
	"time"

	"go-crud/helper"
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Sugnup(c *gin.Context) {
	// Get data from request body
	var body struct {
		Email    string
		Password string
	}

	err := c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	// Create user
	user := models.User{
		Email:    body.Email,
		Password: string(hashedPassword),
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Create user",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}

func Login(c *gin.Context) {
	// Get data from body
	var body struct {
		Email    string
		Password string
	}

	err := c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Lookup requested user
	var user models.User
	// initializers.DB.First(&user, "email = ?", body.Email)
	initializers.DB.Table("users").Select("id", "email", "password").Where("email = ?", body.Email).Scan(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email",
		})
		return
	}
	// Compare sent in pass with saved user pass hash
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	valid, msg := helper.VerifyPassword(user.Password, body.Password)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": msg,
		})
		return
	}

	// Generate jwt token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.UserID,
		"type": "refresh",
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.UserID,
		"type": "refresh",
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"refreshToken": refreshTokenString,
		"accessToken":  accessTokenString,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

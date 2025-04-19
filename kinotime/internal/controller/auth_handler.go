package controller

import (
	"log/slog"
	"net/http"
	"time"

	"kinotime/internal/models"
	"kinotime/internal/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HandleLogin(c *gin.Context, userRepo *repository.UserRepository) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	storedPassword, exists := userRepo.AuthenticateUser(c, user.Username, user.Password)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * time.Duration(JwtExp)).Unix(),
	})

	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
	})
}

func HandleRegister(c *gin.Context, userRepo *repository.UserRepository) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		slog.Error(err.Error())
		return
	}

	if _, exists := userRepo.GetUserByUsername(c, user.Username); exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	err := userRepo.CreateUser(c, user.Username, user.Password)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully registered",
	})
}

func HandleProfile(c *gin.Context, userRepo *repository.UserRepository) {
	username, _ := c.Get("username")
	password, _ := userRepo.GetUserByUsername(c, username.(string))
	c.JSON(http.StatusOK, gin.H{"username": username, "password": password})
}

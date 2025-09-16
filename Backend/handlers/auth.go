package handlers

import (
	"net/http"
	"os"
	"time"

	"authdemo/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterPayload struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

type LoginPayload struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthHandler struct {
	store UserStore
}

type UserStore interface {
	FindByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) error
}

func NewAuthHandler(store UserStore) *AuthHandler {
	return &AuthHandler{store: store}
}

func getSecret() []byte {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev-secret-change-me"
	}
	return []byte(jwtSecret)
}

func (h *AuthHandler) Register(c *gin.Context) {
	var p RegisterPayload
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(p.Username) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username must be at least 3 characters"})
		return
	}
	if len(p.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password must be at least 6 characters"})
		return
	}

	if _, err := h.store.FindByUsername(p.Username); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
		return
	}

	if p.Password != p.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password and confirmPassword do not match"})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	u := &models.User{Username: p.Username, Password: string(hash)}

	if err := h.store.CreateUser(u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "registered"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var p LoginPayload
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.store.FindByUsername(p.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": p.Username,
		"exp": time.Now().Add(2 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})
	ss, _ := token.SignedString(getSecret())

	c.JSON(http.StatusOK, gin.H{"token": ss, "username": u.Username})
}

func Me(c *gin.Context) {
	u := c.GetString("username")
	c.JSON(http.StatusOK, gin.H{"username": u})
}

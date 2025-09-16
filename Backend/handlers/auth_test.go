package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"authdemo/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func performRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

type MockUserStore struct {
	users map[string]*models.User
}

func NewMockUserStore() *MockUserStore {
	return &MockUserStore{users: make(map[string]*models.User)}
}

func (m *MockUserStore) FindByUsername(username string) (*models.User, error) {
	if u, ok := m.users[username]; ok {
		return u, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *MockUserStore) CreateUser(user *models.User) error {
	if _, exists := m.users[user.Username]; exists {
		return gorm.ErrRegistered
	}
	m.users[user.Username] = user
	return nil
}

func TestRegister_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	store := NewMockUserStore()
	h := NewAuthHandler(store)

	r := gin.Default()
	r.POST("/register", h.Register)

	payload := RegisterPayload{
		Username:        "user1",
		Password:        "123456",
		ConfirmPassword: "123456",
	}
	body, _ := json.Marshal(payload)

	w := performRequest(r, "POST", "/register", body)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "registered")
}

func TestRegister_UsernameTooShort(t *testing.T) {
	gin.SetMode(gin.TestMode)
	store := NewMockUserStore()
	h := NewAuthHandler(store)

	r := gin.Default()
	r.POST("/register", h.Register)

	payload := RegisterPayload{
		Username:        "ab",
		Password:        "123456",
		ConfirmPassword: "123456",
	}
	body, _ := json.Marshal(payload)

	w := performRequest(r, "POST", "/register", body)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "username must be at least 3 characters")
}

func TestRegister_PasswordTooShort(t *testing.T) {
	gin.SetMode(gin.TestMode)
	store := NewMockUserStore()
	h := NewAuthHandler(store)

	r := gin.Default()
	r.POST("/register", h.Register)

	payload := RegisterPayload{
		Username:        "validuser",
		Password:        "123",
		ConfirmPassword: "123",
	}
	body, _ := json.Marshal(payload)

	w := performRequest(r, "POST", "/register", body)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "password must be at least 6 characters")
}

func TestRegister_DuplicateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	store := NewMockUserStore()
	h := NewAuthHandler(store)

	// seed user
	store.users["user1"] = &models.User{Username: "user1", Password: "hashed"}

	r := gin.Default()
	r.POST("/register", h.Register)

	payload := RegisterPayload{
		Username:        "user1",
		Password:        "123456",
		ConfirmPassword: "123456",
	}
	body, _ := json.Marshal(payload)

	w := performRequest(r, "POST", "/register", body)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Contains(t, w.Body.String(), "username already exists")
}

func TestRegister_PasswordMismatch(t *testing.T) {
	gin.SetMode(gin.TestMode)
	store := NewMockUserStore()
	h := NewAuthHandler(store)

	r := gin.Default()
	r.POST("/register", h.Register)

	payload := RegisterPayload{
		Username:        "user2",
		Password:        "123456",
		ConfirmPassword: "654321",
	}
	body, _ := json.Marshal(payload)

	w := performRequest(r, "POST", "/register", body)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "password and confirmPassword do not match")
}

func TestLogin_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	store := NewMockUserStore()
	h := NewAuthHandler(store)

	// seed user
	hash, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	store.users["user1"] = &models.User{Username: "user1", Password: string(hash)}

	r := gin.Default()
	r.POST("/login", h.Login)

	payload := LoginPayload{
		Username: "user1",
		Password: "123456",
	}
	body, _ := json.Marshal(payload)

	w := performRequest(r, "POST", "/login", body)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}

func TestLogin_WrongPassword(t *testing.T) {
	gin.SetMode(gin.TestMode)
	store := NewMockUserStore()
	h := NewAuthHandler(store)

	// seed user
	hash, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	store.users["user1"] = &models.User{Username: "user1", Password: string(hash)}

	r := gin.Default()
	r.POST("/login", h.Login)

	payload := LoginPayload{
		Username: "user1",
		Password: "wrongpass",
	}
	body, _ := json.Marshal(payload)

	w := performRequest(r, "POST", "/login", body)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "invalid username or password")
}

func TestLogin_UserNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	store := NewMockUserStore()
	h := NewAuthHandler(store)

	r := gin.Default()
	r.POST("/login", h.Login)

	payload := LoginPayload{
		Username: "nouser",
		Password: "123456",
	}
	body, _ := json.Marshal(payload)

	w := performRequest(r, "POST", "/login", body)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "invalid username or password")
}

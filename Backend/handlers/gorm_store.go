package handlers

import (
	"authdemo/models"

	"gorm.io/gorm"
)

type GormUserStore struct {
	DB *gorm.DB
}

func (s *GormUserStore) FindByUsername(username string) (*models.User, error) {
	var u models.User
	if err := s.DB.Where("username = ?", username).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *GormUserStore) CreateUser(user *models.User) error {
	return s.DB.Create(user).Error
}

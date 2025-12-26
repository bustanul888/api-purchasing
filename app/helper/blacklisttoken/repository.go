package blacklisttoken

import (
	"task-be/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	Create(blacklisttoken model.BlackListToken) error
	GetByToken(token string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(blacklisttoken model.BlackListToken) error {
	return r.db.Create(&blacklisttoken).Error
}

func (r *repository) GetByToken(token string) error {
	var blacklisttoken model.BlackListToken
	return r.db.Where("token = ?", token).First(&blacklisttoken).Error
}
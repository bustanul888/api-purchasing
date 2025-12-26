package user

import (
	"task-be/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	create(data model.User) error
	update(id string, newUserName string, role string) error
	updatePassword(id string, newPass string) error
	delete(id string) error
	getAll() []userResponse
	GetById(id string) userResponse
	GetByUsername(username string) userResponse
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) create(data model.User) error{
	return r.db.Create(&data).Error
}

func (r *repository) update(id string, newUserName string, role string) error{
	return r.db.Where("id = ?",id).Updates(map[string]interface{}{
		"user_name":newUserName,
		"role":role,
		}).Error
}

func (r *repository) updatePassword(id string, newPass string) error{
	return r.db.Where("id = ?",id).Updates(map[string]interface{}{
		"password":newPass,
		}).Error
}

func (r *repository) delete(id string) error{
	return r.db.Where("id = ?",id).Delete(&model.User{}).Error
}

func (r *repository) getAll() []userResponse{
	var res []userResponse
	r.db.Find(&res)
	return res
}

func (r *repository) GetById(id string) userResponse{
	var res userResponse
	r.db.Where("id = ?",id).Find(&res)
	return res
}

func (r *repository) GetByUsername(username string) userResponse {
	var res userResponse
	r.db.Where("user_name = ?", username).Find(&res)
	return res
}

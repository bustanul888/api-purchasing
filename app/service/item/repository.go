package item

import (
	"task-be/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	create(data model.Item) error
	update(id string, Name string, stock uint16, price uint64) error
	UpdateStock(tx *gorm.DB,id string, stock uint16) error
	delete(id string) error
	getAll() []itemResponse
	GetById(id string) itemResponse
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) create(data model.Item) error{
	return r.db.Create(&data).Error
}

func (r *repository) update(id string, Name string, stock uint16,price uint64) error{
	return r.db.Where("id = ?",id).Model(&model.Item{}).Updates(map[string]interface{}{
		"name":Name,
		"stock":stock,
		"price":price,
		}).Error
}

func (r *repository) UpdateStock(tx *gorm.DB,id string, stock uint16) error{
	return tx.Debug().Where("id = ?",id).Model(&model.Item{}).Updates(map[string]interface{}{
		"stock":stock,
		}).Error
}

func (r *repository) delete(id string) error{
	return r.db.Where("id = ?",id).Delete(&model.Item{}).Error
}

func (r *repository) getAll() []itemResponse{
	var res []itemResponse
	r.db.Find(&res)
	return res
}

func (r *repository) GetById(id string) itemResponse{
	var res itemResponse
	r.db.Where("id = ?",id).Find(&res)
	return res
}
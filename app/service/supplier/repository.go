package supplier

import (
	"fmt"
	"task-be/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	create(data model.Supplier) error
	update(id string, Name,Email,address string) error
	delete(id string) error
	getAll() []supplierResponse
	getById(id string) supplierResponse
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) create(data model.Supplier) error{
	fmt.Println("data",data)
	return r.db.Create(&data).Error
}

func (r *repository) update(id string, Name string, Email string,address string) error{
	return r.db.Where("id = ?",id).Updates(map[string]interface{}{
		"name":Name,
		"email":Email,
		"address":address,
		}).Error
}

func (r *repository) delete(id string) error{
	return r.db.Where("id = ?",id).Delete(&model.Supplier{}).Error
}

func (r *repository) getAll() []supplierResponse{
	var res []supplierResponse
	r.db.Find(&res)
	return res
}

func (r *repository) getById(id string) supplierResponse{
	var res supplierResponse
	r.db.Where("id = ?",id).Find(&res)
	return res
}
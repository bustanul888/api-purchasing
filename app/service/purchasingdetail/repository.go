package purchasingdetail

import (
	"task-be/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	Create(tx *gorm.DB,data model.PurchasingDetail) error
	update(id string, Name string, stock uint16, price uint64) error
	delete(id string) error
	Delete(tx *gorm.DB,purchasingId string) error
	// getAll() []itemResponse
	// getById(id string) itemResponse
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(tx *gorm.DB,data model.PurchasingDetail) error{
	return tx.Create(&data).Error
}

func (r *repository) update(id string, Name string, stock uint16,price uint64) error{
	return r.db.Where("id = ?",id).Updates(map[string]interface{}{
		"name":Name,
		"stock":stock,
		"price":price,
		}).Error
}

func (r *repository) Delete(tx *gorm.DB,purchasingId string) error{
	return tx.Where("purchasing_id = ?",purchasingId).Delete(&model.PurchasingDetail{}).Error
}

func (r *repository) delete(id string) error{
	return r.db.Where("id = ?",id).Delete(&model.PurchasingDetail{}).Error
}

// func (r *repository) getAll() []itemResponse{
// 	var res []itemResponse
// 	r.db.Find(&res)
// 	return res
// }

// func (r *repository) getById(id string) itemResponse{
// 	var res itemResponse
// 	r.db.Where("id = ?",id).Find(&res)
// 	return res
// }
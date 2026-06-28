package purchasing

import (
	"task-be/app/model"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	pool() *gorm.DB
	create(tx *gorm.DB,data model.Purchasing) (model.Purchasing,error)
	updateTotal(tx *gorm.DB,id string, price uint64) error
	delete(tx *gorm.DB,id string) error
	getAll() []purchasingResponse
	getById(id string) purchasingResponse
	getDashboard(start, end time.Time) []purchasingResponse
	update(id string,supplierId string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) pool() *gorm.DB{
	return r.db
}

func (r *repository) create(tx *gorm.DB,data model.Purchasing) (model.Purchasing,error){
	err:=tx.Create(&data).Error
	return data,err
}

func (r *repository) updateTotal(tx *gorm.DB,id string, price uint64) error{
	return tx.Where("id = ?",id).Model(&model.Purchasing{}).Updates(map[string]interface{}{
		"grand_total":price,
		}).Error
}

func (r *repository) delete(tx *gorm.DB,id string) error{
	return tx.Where("id = ?",id).Delete(&model.Purchasing{}).Error
}

func (r *repository) getAll() []purchasingResponse{
	var res []purchasingResponse
	r.db.Joins("Supplier").Joins("User").Preload("PurchasingDetails.Item").Find(&res)
	return res
}

func (r *repository) getById(id string) purchasingResponse{
	var res purchasingResponse
	r.db.Joins("Supplier").Joins("User").Preload("PurchasingDetails.Item").Where("id = ?",id).Find(&res)
	return res
}

func (r *repository) update(id string,supplierId string) error{
	return r.db.Where("id = ?",id).Updates(map[string]interface{}{
		"supplier_id":supplierId,
		}).Error
}

func (r *repository) getDashboard(start time.Time, end time.Time) []purchasingResponse {
	var res []purchasingResponse
	r.db.Joins("Supplier").Joins("User").Preload("PurchasingDetails.Item").
		Where("date >= ? AND date <= ?", start, end).
		Find(&res)
	return res
}
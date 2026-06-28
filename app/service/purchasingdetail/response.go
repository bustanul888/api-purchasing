package purchasingdetail

import "task-be/app/model"

type PurchasingDetailResponse struct{
	ID			string	`json:"ID"`
	ItemId		string	`json:"-"`
	PurchasingId	string	`json:"-"`
	Quantity	uint16	`json:"quantity"`
	Subtotal	uint64	`json:"sub_total"`
	Item		itemResponse	`gorm:"foreignKey:ItemId;references:ID" json:"item"`
	model.DateTime
}

func (PurchasingDetailResponse) TableName() string{
	return "purchasing_details"
}

type itemResponse struct{
	ID			string	`json:"ID"`
	Name		string	`json:"name"`
	model.DateTime
}

func (itemResponse) TableName() string{
	return "items"
}

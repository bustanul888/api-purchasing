package model

type PurchasingDetail struct{
	BaseModel
	PurchasingId	string	`gorm:"type:varchar(50)"`
	ItemId 			string	`gorm:"type:varchar(50)"`
	Quantity 		uint16	`gorm:"type:int"`
	Subtotal 		uint32	`gorm:"type:int"`
	Purchasing 		Purchasing	`gorm:"foreignKey:PurchasingId;references:ID"`
	Item 			Item		`gorm:"foreignKey:ItemId;references:ID"`
	DateTime
}

func (PurchasingDetail) TableName()string{
	return "purchasing_details"
}
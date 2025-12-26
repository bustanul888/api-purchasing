package model

import "time"

type Purchasing struct{
	BaseModel
	Date 		time.Time 	`gorm:"type:datetime"`
	SupplierId 	string		`gorm:"type:varchar(50)"`
	UserId 		string		`gorm:"type:varchar(50)"`
	GrandTotal 	*uint64		`gorm:"type:int"`
	Supplier 	Supplier	`gorm:"foreignKey:SupplierId;references:ID"`
	User 		User		`gorm:"foreignKey:UserId;references:ID"`
	DateTime
}

func (Purchasing) TableName() string{
	return "purchasings"
}
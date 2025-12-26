package purchasing

import (
	"task-be/app/model"
	"time"
)

type purchasingResponse struct{
	model.BaseModel
	Date 		time.Time	`json:"date"`
	SupplierId 	string	`json:"-"`
	UserId		string	`json:"-"`
	GrandTotal  uint64	`json:"Total"`
	Supplier 	supplierResponse	`gorm:"foreignKey:SupplierId;references:ID" json:"supplier"`
	User 		userResponse	`gorm:"foreignKey:UserId;references:ID" json:"user"`
	PurchasingDetails []purchasingDetailResponse	`gorm:"foreignKey:PurchasingId;references:ID" json:"purchasing_details"`
	model.DateTime
}

func (purchasingResponse) TableName() string{
	return "purchasings"
}

type supplierResponse struct{
	ID 		string		`json:"ID"`
	Name 	string		`json:"name"`
	Email 	string		`json:"email"`
	Address string		`json:"address"`
	model.DateTime
}

func (supplierResponse) TableName() string{
	return "suppliers"
}

type userResponse struct{
	ID 			string		`json:"ID"`
	UserName 	string		`json:"user_name"`
	Role 		string		`json:"role"`
	model.DateTime
}

func (userResponse) TableName() string{
	return "users"
}

type purchasingDetailResponse struct{
	ID			string	`json:"ID"`
	ItemId		string	`json:"-"`
	PurchasingId	string	`json:"-"`
	Quantity	uint16	`json:"quantity"`
	Subtotal	uint64	`json:"sub_total"`
	Item		itemResponse	`gorm:"foreignKey:ItemId;references:ID" json:"item"`
	model.DateTime
}

func (purchasingDetailResponse) TableName() string{
	return "purchasing_details"
}

type itemResponse struct{
	ID			string	`json:"ID"`
	Name		string	`json:"name"`
	Stock		uint16	`json:"stock"`
	Price		uint64	`json:"price"`
	model.DateTime
}

func (itemResponse) TableName() string{
	return "items"
}

type responseItemDashboard struct{
	Date		time.Time	`json:"date"`
	Name		string	`json:"name"`
	Stock		uint16	`json:"stock"`
	Price		uint64	`json:"price"`
	GrandTotal	uint64	`json:"grand_total"`
}

type responseDashboard struct{
	TotalPurchasing uint64	`json:"total_purchasing"`
	TotalItem uint64	`json:"total_item"`
	TotalStock uint64	`json:"total_stock"`
	Purchasing []responseItemDashboard	`json:"purchasing"`
}
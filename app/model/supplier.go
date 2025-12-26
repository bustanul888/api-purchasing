package model

type Supplier struct{
	BaseModel
	Name 		string	`gorm:"type:varchar(100)"`
	Email 		string	`gorm:"type:varchar(50)"`
	Address 	string	`gorm:"type:varchar(255)"`
	DateTime
}

func (Supplier) TableName() string{
	return "suppliers"
}
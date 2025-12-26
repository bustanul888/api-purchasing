package model

type Item struct{
	BaseModel
	Name 	string	`gorm:"type:varchar(100)"`
	Stock 	uint16	`gorm:"Type:int"`
	Price 	uint64	`gorm:"type:int"`
	DateTime
}

func (Item) TableName() string{
	return "items"
}
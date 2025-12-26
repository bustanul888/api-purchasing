package item

import "task-be/app/model"

type itemResponse struct{
	model.BaseModel
	Name 	string	`json:"name"`
	Stock 	uint16	`json:"stock"`
	Price	uint64	`json:"price"`
	model.DateTime
}

func (itemResponse) TableName() string{
	return "items"
}
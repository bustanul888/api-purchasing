package supplier

import "task-be/app/model"

type supplierResponse struct{
	model.BaseModel
	Name 	string	`json:"name"`
	Email 	string	`json:"email"`
	Address		string	`json:"address"`
	model.DateTime
}

func (supplierResponse) TableName() string{
	return "suppliers"
}
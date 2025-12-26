package supplier

type supplierRequest struct{
	Name string		`json:"name" validate:"required"`
	Email string		`json:"email" validate:"required"`
	Address string			`json:"address" validate:"required"`
}


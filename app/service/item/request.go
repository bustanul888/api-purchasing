package item

type itemRequest struct{
	Name string		`json:"name" validate:"required"`
	Stock uint16		`json:"stock" validate:"required"`
	Price uint64			`json:"price" validate:"required"`
}


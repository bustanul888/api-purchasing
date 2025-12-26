package item

type itemRequest struct{
	Name string		`json:"name" binding:"required"`
	Stock uint16		`json:"stock" binding:"required"`
	Price uint64			`json:"price" binding:"required"`
}


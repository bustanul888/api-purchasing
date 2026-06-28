package purchasing

type purchasingDetail struct{
	ItemId 	string `json:"item_id" validate:"required"`
	Quantity uint16 `json:"quantity" validate:"required"`
}

type purchasingRequest struct{
	SupplierId 			string				`json:"supplier_id" validate:"required"`
	PurchasingDetail 	[]purchasingDetail	`json:"purchasing_detail"`
}

type updatePurchasingRequest struct{
	SupplierId string `json:"supplier_id" validate:"required"`
}
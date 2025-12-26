package purchasing

type purchasingDetail struct{
	ItemId 	string `json:"item_id" binding:"required"`
	Quantity uint16 `json:"quantity" binding:"required"`
}

type purchasingRequest struct{
	SupplierId 			string				`json:"supplier_id" binding:"required"`
	PurchasingDetail 	[]purchasingDetail	`json:"purchasing_detail"`
}

type updatePurchasingRequest struct{
	SupplierId string `json:"supplier_id" binding:"required"`
}
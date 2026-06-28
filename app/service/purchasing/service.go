package purchasing

import (
	"fmt"
	"task-be/app/helper"
	"task-be/app/model"
	"task-be/app/service/item"
	"task-be/app/service/purchasingdetail"
	"time"

	"gorm.io/gorm"
)

type Service interface {
	create(userId string,req purchasingRequest) error
	getAll() []purchasingResponse
	getById(id string) purchasingResponse
	update(id string,req updatePurchasingRequest) error
	delete(id string) error
}

type service_ struct {
	repository Repository
	purchDetailRepo purchasingdetail.Repository
	itemRepo item.Repository
}

func NewService(repository Repository,purchDetailRepo purchasingdetail.Repository,itemRepo item.Repository) *service_ {
	return &service_{repository,purchDetailRepo,itemRepo}
}

func (s *service_) create(userId string,req purchasingRequest) error{
	return s.repository.pool().Transaction(func(tx *gorm.DB) error {
		fmt.Println("userId",userId)
		presentTime := time.Now().UTC()
		data,err := s.repository.create(tx,model.Purchasing{
			Date: presentTime,
			SupplierId: req.SupplierId,
			UserId: userId,
		})
		if err!=nil{
			fmt.Println("---------------------")
			fmt.Println(err)
			return err
		}

		grandTotal,err := s.createDetail(tx,data.ID,req.PurchasingDetail)
		if err!=nil {
			fmt.Println(err)
			return err
		}

		err = s.repository.updateTotal(tx,data.ID,grandTotal)
		if err!=nil {
			fmt.Println(err)
			return err
		}

		return nil
	})
}

func (s *service_) createDetail(tx *gorm.DB,purchasingId string,request []purchasingDetail) (uint64,error){
	var grandTotal uint64
	var err error
	for _,req := range request{
		itemData := s.itemRepo.GetById(req.ItemId)
		subTotal := itemData.Price * uint64(req.Quantity)
		fmt.Println("subTotal",subTotal)
		newStock := itemData.Stock + uint16(req.Quantity)
		err = s.itemRepo.UpdateStock(tx,req.ItemId,newStock)
		if err!=nil {
			fmt.Println(err)
			return 0,err
		}
		err = s.purchDetailRepo.Create(tx,model.PurchasingDetail{
			PurchasingId: purchasingId,
			ItemId: req.ItemId,
			Quantity: req.Quantity,
			Subtotal: uint32(subTotal),
		})
		if err!=nil {
			fmt.Println(err)
			return 0,err
		}
		grandTotal += subTotal
	}
	return grandTotal,err
}
func (s *service_) getAll() []purchasingResponse{
	return s.repository.getAll()
}

func (s *service_) getById(id string) purchasingResponse{
	return s.repository.getById(id)
}

func (s *service_) getDetailById(id string) []purchasingdetail.PurchasingDetailResponse{
	return s.purchDetailRepo.GetByPurchasingId(id)
}

func (s *service_) update(id string,req updatePurchasingRequest) error{
	return s.repository.update(id,req.SupplierId)
}

func (s *service_) delete(id string) error{
	return s.repository.pool().Transaction(func(tx *gorm.DB) error {
		err := s.repository.delete(tx,id)
		if err!=nil{
			return err
		}
		err = s.purchDetailRepo.Delete(tx,id)
		if err!=nil{
			return err
		}
		return nil
	})
}

func (s *service_) dashboard(startDate, endDate *string) (responseDashboard, error){
	var res responseDashboard
	var err error
	
	// Parse/Setup date range filters
	var start, end time.Time
	if startDate != nil && *startDate != "" {
		if t, err := helper.StringToDate(*startDate); err == nil {
			start = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
		} else {
			return responseDashboard{}, err
		}
	} else {
		start = helper.UtcTime().AddDate(0, 0, -30)
		start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, time.UTC)
	}

	if endDate != nil && *endDate != "" {
		if t, err := helper.StringToDate(*endDate); err == nil {
			end = time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, time.UTC)
		} else {
			return responseDashboard{}, err
		}
	} else {
		end = helper.UtcTime()
		end = time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 999999999, time.UTC)
	}

	// Fetch filtered purchasing records from repository
	data := s.repository.getDashboard(start, end)

	// Initialize stats & slices
	res.TotalPurchasing = 0
	res.TotalItem = 0
	res.TotalStock = 0
	res.Purchasing = []responseItemDashboard{}
	
	itemMap := make(map[string]bool)
	for _, v := range data {
		res.TotalPurchasing += v.GrandTotal
		for _, r := range v.PurchasingDetails {
			res.TotalStock += uint64(r.Quantity)
			if _, exists := itemMap[r.Item.ID]; !exists {
				res.TotalItem++
				itemMap[r.Item.ID] = true
			}
			res.Purchasing = append(res.Purchasing, responseItemDashboard{
				Date:       v.Date,
				Name:       r.Item.Name,
				Stock:      r.Quantity,
				Price:      uint64(r.Item.Price),
				GrandTotal: r.Subtotal,
			})
		}
	}
	return res, err
}
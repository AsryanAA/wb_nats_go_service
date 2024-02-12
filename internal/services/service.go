package services

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"wb_nats_go_service/internal/models"
)

var Cache = make(map[string]*models.Order)

func SaveInDB(db *gorm.DB, o models.Order) error {
	var item models.Item
	var order models.OrderForDB
	var orderItems []models.Item

	db.Where("order_uid = ?", o.OrderUID).Find(&order)
	if order.OrderUID != "" {
		fmt.Println("Сообщение не сохранено в БД запись уже существует")
		return errors.New("Запись с таким id уже существует")
	}
	orderItems = o.Items

	order.OrderUID = o.OrderUID
	order.TrackNumber = o.TrackNumber
	order.Entry = o.Entry
	order.Name = o.Name
	order.Phone = o.Phone
	order.Zip = o.Zip
	order.City = o.City
	order.Address = o.Address
	order.Region = o.Region
	order.Email = o.Email
	order.Transaction = o.Transaction
	order.RequestID = o.RequestID
	order.Currency = o.Currency
	order.Provider = o.Provider
	order.Amount = o.Amount
	order.PaymentDt = o.PaymentDt
	order.Bank = o.Bank
	order.DeliveryCost = o.DeliveryCost
	order.GoodsTotal = o.GoodsTotal
	order.CustomFee = o.CustomFee
	order.Locale = o.Locale
	order.InternalSignature = o.InternalSignature
	order.CustomerID = o.CustomerID
	order.DeliveryService = o.DeliveryService
	order.ShardKey = o.ShardKey
	order.SmID = o.SmID
	order.DateCreated = o.DateCreated
	order.OofShard = o.OofShard

	if len(orderItems) > 0 {
		for _, val := range orderItems {
			item.OrderId = order.OrderUID
			item.ChrtID = val.ChrtID
			item.TrackNumber = val.TrackNumber
			item.Price = val.Price
			item.Rid = val.Rid
			item.Name = val.Name
			item.Sale = val.Sale
			item.Size = val.Size
			item.TotalPrice = val.TotalPrice
			item.NmID = val.NmID
			item.Brand = val.Brand
			item.Status = val.Status
			db.Create(item)
		}
	} else {
		db.Create(order)
	}
	db.Create(order)
	return nil
}

func SaveInCache(o models.Order) error {
	Cache[o.OrderUID] = &o
	fmt.Println("Количество записей в хэш", len(Cache))
	return nil
}

func ReadInCache(id string) *models.Order {
	val, ok := Cache[id]
	if ok {
		return val
	}
	return &models.Order{}
}

func LoadInCacheInDB(db *gorm.DB) error {
	fmt.Println("In cache before load", len(Cache))
	var orders []models.OrderForDB
	var items []models.Item
	var orderCache models.Order
	db.Find(&orders)
	if len(orders) > 0 {
		for _, o := range orders {
			orderCache.OrderUID = o.OrderUID
			orderCache.TrackNumber = o.TrackNumber
			orderCache.Entry = o.Entry
			orderCache.Name = o.Name
			orderCache.Phone = o.Phone
			orderCache.Zip = o.Zip
			orderCache.City = o.City
			orderCache.Address = o.Address
			orderCache.Region = o.Region
			orderCache.Email = o.Email
			orderCache.Transaction = o.Transaction
			orderCache.RequestID = o.RequestID
			orderCache.Currency = o.Currency
			orderCache.Provider = o.Provider
			orderCache.Amount = o.Amount
			orderCache.PaymentDt = o.PaymentDt
			orderCache.Bank = o.Bank
			orderCache.DeliveryCost = o.DeliveryCost
			orderCache.GoodsTotal = o.GoodsTotal
			orderCache.CustomFee = o.CustomFee
			orderCache.Locale = o.Locale
			orderCache.InternalSignature = o.InternalSignature
			orderCache.CustomerID = o.CustomerID
			orderCache.DeliveryService = o.DeliveryService
			orderCache.ShardKey = o.ShardKey
			orderCache.SmID = o.SmID
			orderCache.DateCreated = o.DateCreated
			orderCache.OofShard = o.OofShard

			db.Where("order_id = ?", o.OrderUID).Find(&items)
			if len(items) > 0 {
				orderCache.Items = items
			}
			Cache[o.OrderUID] = &orderCache
		}
	}
	fmt.Println("In cache after load", len(Cache))
	return nil
}

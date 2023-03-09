package datastore

import (
	"draft/model"
)

type MemCache struct {
	buyerMap map[string]*model.Buyer
	shopMap  map[string]*model.Shop
}

// В рамках задачи нет смысла использовать реальную базу,
// если бы это был реальный проект, вероятно я бы использовал
// интерфейс, под который можно подогнать несколько датасторов
// от структур в памяти и файлов до баз данных и связей с другими сервисами
func NewDatastore() *MemCache {
	return &MemCache{
		buyerMap: make(map[string]*model.Buyer),
		shopMap:  make(map[string]*model.Shop),
	}
}

// Ну а дальше идут просто круды

func (m *MemCache) SaveBuyer(buyer *model.Buyer) error {
	m.buyerMap[buyer.LastName] = buyer
	return nil
}

func (m *MemCache) SaveShop(shop *model.Shop) error {
	m.shopMap[shop.Name] = shop
	return nil
}

func (m *MemCache) GetBuyer(lastName string) *model.Buyer {
	buyer, ok := m.buyerMap[lastName]
	if !ok {
		return nil
	}
	return buyer
}

func (m *MemCache) GetShop(name string) *model.Shop {
	shop, ok := m.shopMap[name]
	if !ok {
		return nil
	}
	return shop
}

func (m *MemCache) DeleteBuyer(buyer *model.Buyer) error {
	delete(m.buyerMap, buyer.LastName)
	return nil
}

func (m *MemCache) DeleteShop(shop *model.Shop) error {
	delete(m.shopMap, shop.Name)
	return nil
}

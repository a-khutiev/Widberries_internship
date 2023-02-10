package cache

import (
	"context"
	"log"
	"sync"

	"github.com/a-khutiev/Widberries_internship/internal/db"
	"github.com/a-khutiev/Widberries_internship/internal/model"
)

type Cache struct {
	database *db.Database
	orders   map[string]*model.Order
	mutex    sync.RWMutex
}

func New(ctx context.Context, database *db.Database) (*Cache, error) {
	orders, err := database.GetOrders(ctx)
	if err != nil {
		return nil, err
	}

	ordersMap := make(map[string]*model.Order, len(orders))
	for _, order := range orders {
		ordersMap[order.OrderUid] = order
	}

	return &Cache{
		database: database,
		orders:   ordersMap,
		mutex:    sync.RWMutex{},
	}, err
}

func (cache *Cache) GetOrder(ctx context.Context, id string) (*model.Order, error) {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	order, ok := cache.orders[id]
	if ok {
		log.Println("got order from cache")
		return order, nil
	}

	order, err := cache.database.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (cache *Cache) CreateOrder(ctx context.Context, order *model.Order) error {
	cache.mutex.RLock()

	_, ok := cache.orders[order.OrderUid]
	if ok {
		cache.mutex.RUnlock()
		return nil
	}

	cache.mutex.RUnlock()
	err := cache.database.CreateOrder(ctx, order)
	if err != nil {
		return err
	}

	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	cache.orders[order.OrderUid] = order

	return nil
}

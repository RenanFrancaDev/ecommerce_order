package mongo

import (
	"context"
	"ecommerce_order/internal/application/ports"
	"ecommerce_order/internal/domain/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(client *mongo.Client, dbName, collectionName string) ports.OrderRepository {
	return &OrderRepository{
		collection: client.Database(dbName).Collection(collectionName),
	}
}

func (r *OrderRepository) Save(ctx context.Context, order *entity.Order) error {
	_, err := r.collection.InsertOne(ctx, order)
	return err
}

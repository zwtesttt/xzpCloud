package adapters

import (
	"context"
	"github.com/zwtesttt/xzpCloud/internal/order/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionName = "idc_order"

// Order
type Order struct {
	Id          primitive.ObjectID `bson:"id"`
	UserId      string             `bson:"user_id"`
	TotalAmount float64            `bson:"total_amount"`
	Status      int                `bson:"status"`
	Items       []*Item            `bson:"items"`
	CreatedAt   int64              `bson:"created_at"`
	UpdatedAt   int64              `bson:"updated_at"`
	DeletedAt   int64              `bson:"deleted_at"`
}

type Item struct {
	ProductId string  `bson:"product_id"`
	Quantity  int     `bson:"quantity"`
	Price     float64 `bson:"price"`
}

func (i *Item) ToItem() *domain.Item {
	return domain.NewOrderItem(i.ProductId, i.Quantity, i.Price)
}

func ToItems(items []*Item) []*domain.Item {
	result := make([]*domain.Item, 0)
	for _, item := range items {
		result = append(result, item.ToItem())
	}
	return result
}

func (o *Order) ToOrder() *domain.Order {
	return domain.NewOrder(
		o.Id.Hex(),
		o.UserId,
		o.TotalAmount,
		domain.OrderStatus(o.Status),
		ToItems(o.Items),
		o.CreatedAt,
		o.UpdatedAt,
		o.DeletedAt,
	)
}

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	return &OrderRepository{
		collection: db.Collection(collectionName),
	}
}

func (o *OrderRepository) Insert(ctx context.Context, order *domain.Order) error {
	return nil
}

func (o *OrderRepository) FindOne(ctx context.Context, id string) (*domain.Order, error) {
	return nil, nil
}

package adapters

import (
	"context"
	"github.com/zwtesttt/xzpCloud/internal/order/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
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
	items := make([]*Item, 0)
	for _, item := range order.Items() {
		items = append(items, &Item{
			ProductId: item.ProductId(),
			Quantity:  item.Quantity(),
			Price:     item.Price(),
		})
	}

	_, err := o.collection.InsertOne(ctx, &Order{
		Id:          primitive.NewObjectID(),
		UserId:      order.UserId(),
		TotalAmount: order.TotalAmount(),
		Status:      int(order.Status()),
		Items:       items,
		CreatedAt:   time.Now().UnixMilli(),
		UpdatedAt:   time.Now().UnixMilli(),
		DeletedAt:   order.DeletedAt(),
	})
	return err
}

func (o *OrderRepository) FindOne(ctx context.Context, id string) (*domain.Order, error) {
	filter := bson.M{
		"deleted_at": 0,
	}

	if id != "" {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		filter["id"] = objId
	}

	var o1 Order

	err := o.collection.FindOne(ctx, filter).Decode(&o1)
	if err != nil {
		return nil, err
	}

	return o1.ToOrder(), nil
}

func (o *OrderRepository) Update(ctx context.Context, order *domain.Order) error {
	orderId, err := primitive.ObjectIDFromHex(order.Id())
	if err != nil {
		return err
	}

	filter := bson.M{
		"deleted_at": 0,
		"id":         orderId,
	}

	items := make([]*Item, 0)
	for _, item := range order.Items() {
		items = append(items, &Item{
			ProductId: item.ProductId(),
			Quantity:  item.Quantity(),
			Price:     item.Price(),
		})
	}

	_, err = o.collection.UpdateOne(ctx, filter, bson.M{
		"$set": bson.M{
			"user_id":      order.UserId(),
			"total_amount": order.TotalAmount(),
			"status":       int(order.Status()),
			"items":        items,
			"updated_at":   time.Now().UnixMilli(),
		},
	})
	return err
}

func (o *OrderRepository) Find(ctx context.Context, opt *domain.FindOptions) ([]*domain.Order, error) {
	filter := bson.M{
		"deleted_at": 0,
	}
	if opt.UpdatedAt > 0 {
		filter["updated_at"] = bson.M{
			"$gt": opt.UpdatedAt,
		}
	}

	findopts := options.Find().SetSort(
		bson.D{
			{"updated_at", -1}, //倒序
		},
	).SetLimit(int64(opt.Size))

	cursor, err := o.collection.Find(ctx, filter, findopts)
	if err != nil {
		return nil, err
	}

	var res []*Order
	err = cursor.All(ctx, &res)
	if err != nil {
		return nil, err
	}

	var dres []*domain.Order
	for _, v := range res {
		dres = append(dres, v.ToOrder())
	}

	return dres, nil
}

package adapters

import (
	"context"
	"github.com/zwtesttt/xzpCloud/internal/product/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const collectionName = "idc_product"

type Product struct {
	Id          primitive.ObjectID `bson:"id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Price       float64            `bson:"price"`
	Stock       int                `bson:"stock"`
	CreatedAt   int64              `bson:"created_at"`
	UpdatedAt   int64              `bson:"updated_at"`
	DeletedAt   int64              `bson:"deleted_at"`
}

func (p *Product) ToProduct() *domain.Product {
	return domain.NewProduct(
		p.Id.Hex(),
		p.Name,
		p.Description,
		p.Price,
		p.Stock,
		p.CreatedAt,
		p.UpdatedAt,
		p.DeletedAt,
	)
}

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{
		collection: db.Collection(collectionName),
	}
}

func (p *ProductRepository) FindOne(ctx context.Context, id string) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProductRepository) Insert(ctx context.Context, order *domain.Product) error {
	//调整成毫秒级别
	thisTime := time.Now().UnixMilli()

	o := &Product{
		Id:          primitive.NewObjectID(),
		Name:        order.Name(),
		Description: order.Description(),
		Price:       order.Price(),
		Stock:       order.Stock(),
		CreatedAt:   thisTime,
		UpdatedAt:   thisTime,
		DeletedAt:   order.DeletedAt(),
	}
	_, err := p.collection.InsertOne(ctx, o)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepository) Find(ctx context.Context, opt *domain.FindOptions) ([]*domain.Product, error) {
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
	).SetLimit(opt.Size)

	cursor, err := p.collection.Find(ctx, filter, findopts)
	if err != nil {
		return nil, err
	}

	var res []*Product
	err = cursor.All(ctx, &res)
	if err != nil {
		return nil, err
	}

	var dres []*domain.Product
	for _, v := range res {
		dres = append(dres, v.ToProduct())
	}

	return dres, nil
}

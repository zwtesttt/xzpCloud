package adapters

import (
	"context"
	"github.com/zwtesttt/xzpCloud/internal/product/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	//TODO implement me
	panic("implement me")
}

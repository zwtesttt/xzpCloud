package adapters

import (
	"context"
	"github.com/zwtesttt/xzpCloud/internal/policy/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionName = "idc_action_url"

type ActionUrl struct {
	Sid  string `bson:"sid"`
	Name string `bson:"name"`
	Url  string `bson:"url"`
}

func (a *ActionUrl) ToActionUrl() *domain.ActionUrl {
	return domain.NewActionUrl(a.Url, a.Name, a.Sid)
}

type ActionUrlRepository struct {
	collection *mongo.Collection
}

func NewActionUrlRepository(db *mongo.Database) *ActionUrlRepository {
	return &ActionUrlRepository{
		collection: db.Collection(collectionName),
	}
}

func (a *ActionUrlRepository) FindOne(ctx context.Context, sid string) (*domain.ActionUrl, error) {
	return nil, nil
}

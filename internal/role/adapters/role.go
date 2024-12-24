package adapters

import (
	"context"
	"github.com/zwtesttt/xzpCloud/internal/role/domain"
	"github.com/zwtesttt/xzpCloud/pkg/role"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionName = "idc_role"

type Role struct {
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	Policies    []*Policy `bson:"policies"`
	Type        int       `bson:"type"`
	CreatedAt   int64     `bson:"created_at"`
	UpdatedAt   int64     `bson:"updated_at"`
}

type Policy struct {
	Sid    string `bson:"sid"`
	Effect string `bson:"effect"`
}

func (r *Role) ToRole() *domain.Role {
	return domain.NewRole(r.Name, r.Description, ToPolicies(r.Policies), role.Type(r.Type), r.CreatedAt, r.UpdatedAt)
}

func (p *Policy) ToPolicy() *domain.Policy {
	return domain.NewPolicy(p.Sid, p.Effect)
}

func ToPolicies(policies []*Policy) []*domain.Policy {
	var result []*domain.Policy
	for _, policy := range policies {
		result = append(result, policy.ToPolicy())
	}
	return result
}

type RoleRepository struct {
	collection *mongo.Collection
}

func NewRoleRepository(db *mongo.Database) *RoleRepository {
	return &RoleRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *RoleRepository) FindOne(ctx context.Context, t role.Type) (*domain.Role, error) {
	//TODO implement me
	panic("implement me")
}

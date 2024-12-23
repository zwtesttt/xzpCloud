package adapters

import (
	"context"
	"github.com/zwtesttt/xzpCloud/internal/vm/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionName = "idc_vm"

type Vm struct {
	Id           primitive.ObjectID `bson:"id"`
	Name         string             `bson:"name"`
	Type         string             `bson:"type"`
	Status       int                `bson:"status"`
	UserId       string             `bson:"user_id"`
	CreatedAt    int64              `bson:"created_at"`
	UpdatedAt    int64              `bson:"updated_at"`
	ExpirationAt int64              `bson:"expiration_at"`
}

func (v *Vm) ToVm() *domain.Vm {
	return domain.NewVm(v.Id.Hex(), v.Name, domain.VmStatus(v.Status), v.UserId, v.CreatedAt, v.UpdatedAt, v.ExpirationAt)
}

type VmRepository struct {
	collection *mongo.Collection
}

func NewVmRepository(db *mongo.Database) *VmRepository {
	return &VmRepository{
		collection: db.Collection(collectionName),
	}
}

func (v *VmRepository) Insert(ctx context.Context, vm *domain.Vm) error {
	//TODO implement me
	panic("implement me")
}

func (v *VmRepository) FindOne(ctx context.Context, id string) (*domain.Vm, error) {
	//TODO implement me
	panic("implement me")
}

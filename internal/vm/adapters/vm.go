package adapters

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/zwtesttt/xzpCloud/internal/vm/domain"
)

var collectionName = "idc_vm"

type Vm struct {
	Id           primitive.ObjectID `bson:"id"`
	Name         string             `bson:"name"`
	Status       int                `bson:"status"`
	UserId       string             `bson:"user_id"`
	Ip           string             `bson:"ip"`
	Config       *VmConfig          `bson:"config"`
	CreatedAt    int64              `bson:"created_at"`
	UpdatedAt    int64              `bson:"updated_at"`
	ExpirationAt int64              `bson:"expiration_at"`
	DeletedAt    int64              `bson:"deleted_at"`
}

type VmConfig struct {
	Cpu    int    `bson:"cpu"`
	Disk   string `bson:"disk"`
	Memory string `bson:"memory"`
}

func (v *Vm) ToVm() *domain.Vm {
	return domain.NewVm(v.Id.Hex(), v.Name, domain.VmStatus(v.Status), v.UserId, v.Ip, v.ToVmConfig(), v.CreatedAt, v.UpdatedAt, v.ExpirationAt, v.DeletedAt)
}

func (v *Vm) ToVmConfig() *domain.VmConfig {
	return domain.NewVmConfig(v.Config.Cpu, v.Config.Disk, v.Config.Memory)
}

type VmRepository struct {
	collection *mongo.Collection
}

func NewVmRepository(db *mongo.Database) domain.Repository {
	return &VmRepository{
		collection: db.Collection(collectionName),
	}
}

func (v *VmRepository) Insert(ctx context.Context, vm *domain.Vm) (string, error) {
	id := primitive.NewObjectID()
	_, err := v.collection.InsertOne(ctx, &Vm{
		Id:     id,
		Name:   vm.Name(),
		Status: int(vm.Status()),
		UserId: vm.UserId(),
		Ip:     vm.Ip(),
		Config: &VmConfig{
			Cpu:    vm.Config().Cpu(),
			Disk:   vm.Config().Disk(),
			Memory: vm.Config().Memory(),
		},
		CreatedAt:    time.Now().UnixMilli(),
		UpdatedAt:    time.Now().UnixMilli(),
		ExpirationAt: vm.ExpirationAt(),
	})
	if err != nil {
		return "", err
	}
	return id.Hex(), nil
}

func (v *VmRepository) FindOne(ctx context.Context, id string) (*domain.Vm, error) {
	vid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"id": vid}

	v1 := &Vm{}
	err = v.collection.FindOne(ctx, filter).Decode(v1)
	if err != nil {
		return nil, err
	}
	return v1.ToVm(), nil
}

func (v *VmRepository) Delete(ctx context.Context, id string) error {
	vid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"id": vid}
	_, err = v.collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"deleted_at": time.Now().UnixMilli()}})
	if err != nil {
		return err
	}
	return nil
}

func (v *VmRepository) Find(ctx context.Context, opts *domain.VmFindOptions) ([]*domain.Vm, error) {
	filter := bson.M{
		"deleted_at": bson.M{
			"$eq": 0,
		},
	}

	if opts.Name != "" {
		filter["name"] = bson.M{
			"$eq": opts.Name,
		}
	}

	if opts.UserId != "" {
		filter["user_id"] = bson.M{
			"$eq": opts.UserId,
		}
	}

	cursor, err := v.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	res := make([]*Vm, 0)
	err = cursor.All(ctx, &res)
	if err != nil {
		return nil, err
	}

	var dres []*domain.Vm
	for _, v := range res {
		dres = append(dres, v.ToVm())
	}

	return dres, nil
}

func (v *VmRepository) Update(ctx context.Context, vm *domain.Vm) error {
	vid, err := primitive.ObjectIDFromHex(vm.Id())
	if err != nil {
		return err
	}

	filter := bson.M{"id": vid}
	update := bson.M{
		"$set": bson.M{
			"name":          vm.Name(),
			"status":        int(vm.Status()),
			"user_id":       vm.UserId(),
			"ip":            vm.Ip(),
			"updated_at":    time.Now().UnixMilli(),
			"expiration_at": vm.ExpirationAt(),
			"config": &VmConfig{
				Cpu:    vm.Config().Cpu(),
				Disk:   vm.Config().Disk(),
				Memory: vm.Config().Memory(),
			},
		},
	}

	_, err = v.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

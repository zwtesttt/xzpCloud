package adapters

import (
	"context"

	"github.com/zwtesttt/xzpCloud/pkg/role"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/zwtesttt/xzpCloud/internal/user/domain"
)

var collectionName = "idc_user"

// User
type User struct {
	Id        primitive.ObjectID `bson:"id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	Avatar    string             `bson:"avatar"`
	RoleId    int                `bson:"role_id"`
	CreatedAt int64              `bson:"created_at"`
	UpdatedAt int64              `bson:"updated_at"`
	DeletedAt int64              `bson:"deleted_at"`
}

func (u *User) ToUser() *domain.User {
	return domain.NewUser(
		u.Id.Hex(),
		u.Name,
		u.Email,
		u.Password,
		u.Avatar,
		role.Type(u.RoleId),
		u.CreatedAt,
		u.UpdatedAt,
		u.DeletedAt,
	)
}

func NewUser(u *domain.User) *User {
	uid, _ := primitive.ObjectIDFromHex(u.Id())
	return &User{
		Id:        uid,
		Name:      u.Name(),
		Email:     u.Email(),
		Password:  u.Password(),
		Avatar:    u.Avatar(),
		RoleId:    int(u.RoleId()),
		CreatedAt: u.CreatedAt(),
		UpdatedAt: u.UpdatedAt(),
		DeletedAt: u.DeletedAt(),
	}
}

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{collection: db.Collection(collectionName)}
}

func (u *UserRepository) Insert(ctx context.Context, user *domain.User) error {

	_, err := u.collection.InsertOne(ctx, &User{
		Id:        primitive.NewObjectID(),
		Name:      user.Name(),
		Email:     user.Email(),
		Password:  user.Password(),
		Avatar:    user.Avatar(),
		RoleId:    int(user.RoleId()),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
		DeletedAt: user.DeletedAt(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) FindOne(ctx context.Context, id string) (*domain.User, error) {
	filter := bson.M{
		"deleted_at": bson.M{
			"$eq": 0,
		},
	}
	if id != "" {
		objId, _ := primitive.ObjectIDFromHex(id)
		filter["id"] = objId
	}

	var u1 *User
	err := u.collection.FindOne(ctx, filter).Decode(u1)
	if err != nil {
		return nil, err
	}
	return u1.ToUser(), nil
}

func (u *UserRepository) FindOneByEmail(ctx context.Context, email string) (*domain.User, error) {
	filter := bson.M{
		"deleted_at": bson.M{
			"$eq": 0,
		},
	}
	if email != "" {
		filter["email"] = email
	}

	var user User
	err := u.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user.ToUser(), nil
}

func (u *UserRepository) Find(ctx context.Context, opts domain.FindOptions) ([]*domain.User, error) {
	filter := bson.M{
		"deleted_at": bson.M{
			"$eq": 0,
		},
	}

	if opts.Email != "" {
		filter["email"] = bson.M{
			"$eq": opts.Email,
		}
	}

	cursor, err := u.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	res := make([]*User, 0)
	err = cursor.All(ctx, &res)
	if err != nil {
		return nil, err
	}

	dres := make([]*domain.User, 0)
	for _, v := range res {
		dres = append(dres, v.ToUser())
	}

	return dres, nil
}

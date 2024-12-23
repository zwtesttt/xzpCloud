package app

import (
	"context"
	"errors"
	"github.com/zwtesttt/xzpCloud/internal/user/domain"
	"github.com/zwtesttt/xzpCloud/pkg/api"
	"github.com/zwtesttt/xzpCloud/pkg/role"
	"go.mongodb.org/mongo-driver/mongo"
)

type RegisterInput struct {
	Email    string
	Name     string
	Password string
}

type RegisterHandler struct {
	userRepo domain.Repository
}

func NewRegisterHandler(userRepo domain.Repository) *RegisterHandler {
	return &RegisterHandler{userRepo: userRepo}
}

func (h *RegisterHandler) Handle(ctx context.Context, input *RegisterInput) error {
	_, err := h.userRepo.FindOneByEmail(ctx, input.Email)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return err
		}
	}
	if err == nil {
		return api.FindCodeError(api.UserExists)
	}

	return h.userRepo.Insert(ctx, domain.NewUser("", input.Name, input.Email, input.Password, "", role.RoleUser, 0, 0, 0))
}

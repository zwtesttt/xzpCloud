package app

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zwtesttt/xzpCloud/internal/user/domain"
	"github.com/zwtesttt/xzpCloud/pkg/api"
	ijwt "github.com/zwtesttt/xzpCloud/pkg/jwt"
)

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	Token  string
	UserId string
	Name   string
	Avatar string
}

type LoginHandler struct {
	userRepo domain.Repository
}

func NewLoginHandler(userRepo domain.Repository) *LoginHandler {
	return &LoginHandler{userRepo: userRepo}
}

func (h *LoginHandler) Handle(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	find, err := h.userRepo.Find(ctx, domain.FindOptions{Email: input.Email})
	if err != nil {
		return nil, err
	}

	if len(find) == 0 {
		return nil, api.FindCodeError(api.UserLoginPasswordIncorrect)
	}
	user := find[0]

	if user.Password() != input.Password {
		return nil, api.FindCodeError(api.UserLoginPasswordIncorrect)
	}

	cli := jwt.MapClaims{
		"user_id": user.Id(),
		"role_id": user.RoleId(),
	}

	//TODO 把token存入缓存
	token, err := ijwt.GenerateJWT(cli)
	res := &LoginOutput{
		Token:  token,
		UserId: user.Id(),
		Name:   user.Name(),
		Avatar: user.Avatar(),
	}
	return res, nil
}

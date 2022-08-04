package service

import (
	"awahids/go-opinia/model/web"
	"context"
)

type UserService interface {
	SingUp(ctx context.Context, request web.UserSignupRequest) web.UserResponse
	SingIn(ctx context.Context, request web.UserSigninRequest) (web.UserResponse, error)
}

package service

import (
	"awahids/go-opinia/helpers"
	"awahids/go-opinia/model/domain"
	"awahids/go-opinia/model/web"
	"awahids/go-opinia/repository"
	"context"
	"database/sql"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func (service *UserServiceImpl) Singup(ctx context.Context, request web.UserSignupRequest) web.UserResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	user := domain.User{
		Email:    request.Email,
		Password: domain.HashPassword(request.Password),
		Fullname: request.Fullname,
		Phone:    request.Phone,
	}

	user = service.UserRepository.Signup(ctx, tx, user)

	return helpers.ToUserResponse(user)
}

func (service *UserServiceImpl) Signin(ctx context.Context, request web.UserSigninRequest) web.UserResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	user, err := service.UserRepository.findOne(ctx, tx, request.Email)
	helpers.PanicIfError(err)

	if err := domain.CheckPassword(request.Password, user.Password); err != nil {
		panic(err)
	}

	return helpers.ToUserResponse(user)
}

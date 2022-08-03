package repository

import (
	"awahids/go-opinia/model/domain"
	"context"
	"database/sql"
)

type UserRepository interface {
	SignUp(ctx context.Context, tx sql.Tx, user domain.User) domain.User
	SignIn(ctx context.Context, tx sql.Tx, email string, password string) domain.User
}

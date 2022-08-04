package repository

import (
	"awahids/go-opinia/model/domain"
	"context"
	"database/sql"
)

type UserRepository interface {
	Signup(ctx context.Context, tx *sql.Tx, user domain.User)
	Signin(ctx context.Context, tx *sql.Tx, email string, password string) (domain.User, error)
}

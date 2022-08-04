package repository

import (
	"awahids/go-opinia/helpers"
	"awahids/go-opinia/model/domain"
	"context"
	"database/sql"
	"errors"
)

type UserRepositoryImpl struct {
}

func (repository *UserRepositoryImpl) Signup(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users (email, password, fullname, phone) Values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.Email, user.Password, user.Fullname, user.Phone)
	helpers.PanicIfError(err)

	id, err := result.LastInsertId()
	helpers.PanicIfError(err)

	user.Id = int(id)
	return user
}

func (repository *UserRepositoryImpl) Signin(ctx context.Context, tx *sql.Tx, email string, password string) (domain.User, error) {
	SQL := "SELECT id, email, password, fullname, phone FROM users WHERE email = ? AND password = ?"
	rows, err := tx.QueryContext(ctx, SQL, email, password)
	helpers.PanicIfError(err)

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Fullname, &user.Phone)
		helpers.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

package domain

import "context"

type User struct {
	id       string `db:"id"`
	email    string `db:"email"`
	password string `db:"password"`
}

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (User, error)
}
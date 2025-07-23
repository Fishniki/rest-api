package domain

import (
	"context"
)

type User struct {
	Id       string `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Role     string `db:"role"`
	Password string `db:"password"`
}

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (User, error)
	Save(ctx context.Context, req *User) error
	FindById(ctx context.Context, id string) (User, error)
}

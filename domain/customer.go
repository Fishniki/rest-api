package domain

import (
	"context"
	"database/sql"
)

type Customer struct {
	ID        string `DB:"id"`
	Code      string `DB:"code"`
	Name      string `DB:"name"`
	CreatedAt sql.NullTime `DB:"created_at"`
	UpdatedAt sql.NullTime `DB:"updated_at"`
	DeletedAt sql.NullTime `DB:"deleted_at"`
}

type CustomerRepository interface {
	FindAll(ctx context.Context) ([]Customer, error)
	FindByID(ctx context.Context, id string) (Customer, error)
	Save(ctx context.Context, c *Customer) error
	Update(ctx context.Context, c *Customer) error
	Delete(ctx context.Context, id string) error
}
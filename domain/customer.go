package domain

import (
	"context"
	"database/sql"
	"rest-api/dto"
)

type Customer struct {
	ID        string `db:"customer_id"`
	Code      string `db:"code"`
	Name      string `db:"name"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type CustomerRepository interface {
	FindAll(ctx context.Context) ([]Customer, error)
	FindByID(ctx context.Context, customer_id string) (Customer, error)
	Save(ctx context.Context, c *Customer) error
	Update(ctx context.Context, c *Customer) error
	Delete(ctx context.Context, customer_id string) error
}

type CustomerService interface {
	Index(ctx context.Context) ([]dto.CustomerData, error)
	Create(ctx context.Context, req dto.CreateCustomerRequest) error
	Update(ctx context.Context, req dto.UpdateCustomerRequest) error
	Delete(ctx context.Context, customer_id string) error
	Show(ctx context.Context, customer_id string) (dto.CustomerData, error)
}



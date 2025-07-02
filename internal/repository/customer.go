package repository

import (
	"context"
	"rest-api/domain"
)

type customerRepository struct {
}

func NewCustomer() domain.CustomerRepository {
	return &customerRepository{}
}

func (cr customerRepository) FindAll(ctx context.Context) ([]domain.Customer, error){
	panic("implement me")
}

func (cr customerRepository) FindByID(ctx context.Context, id string) (domain.Customer, error){
	panic("implement me")
}

func (cr customerRepository) Save(ctx context.Context, c *domain.Customer) error {
	panic("implement me")
}

func (cr customerRepository) Update(ctx context.Context, c *domain.Customer) error {
	panic("implement me")
}

func (cr customerRepository) Delete(ctx context.Context, id string) error {
	panic("implement me")
}


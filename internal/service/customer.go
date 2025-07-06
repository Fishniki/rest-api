package service

import (
	"context"
	"database/sql"
	"errors"
	"rest-api/domain"
	"rest-api/dto"
	"time"

	"github.com/google/uuid"
)

type CustomerService struct {
	customerRepository domain.CustomerRepository
}


func NewCustomer(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &CustomerService{
		customerRepository: customerRepository,
	}
}

// Index implements domain.CustomerService.
func (c CustomerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := c.customerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var customerData []dto.CustomerData
	for _, v := range customers {
		customerData = append(customerData, dto.CustomerData{
			ID: v.ID,
			Code: v.Code,
			Name: v.Name,
		})
	}

	return  customerData, nil
}


func (c CustomerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error  {
	customer := domain.Customer{
		ID: uuid.NewString(),
		Code: req.Code,
		Name: req.Name,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}

	return  c.customerRepository.Save(ctx, &customer)
	
}

func (c CustomerService) Update(ctx context.Context, req dto.UpdateCustomerRequest) error {
	persisted, err := c.customerRepository.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}

	if persisted.ID == "" {
		return errors.New("data customer tidak di temukan")
	}

	persisted.Code = req.Code
	persisted.Name = req.Name
	persisted.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}

	return  c.customerRepository.Update(ctx, &persisted)
}

func (c CustomerService) Delete(ctx context.Context, customer_id string) error {

	exist, err := c.customerRepository.FindByID(ctx, customer_id)
	if err != nil {
		return  err
	}

	if exist.ID == "" {
		 return  errors.New("data customer tidak di temukan")
	}

	return  c.customerRepository.Delete(ctx, customer_id)

}

func (c CustomerService) Show(ctx context.Context, customer_id string) (dto.CustomerData, error) {

	persisted, err := c.customerRepository.FindByID(ctx, customer_id)
	if err != nil {
		return  dto.CustomerData{}, err
	}

	if persisted.ID == "" {
		return  dto.CustomerData{}, errors.New("data customer tidak di temukan")
	}

	return dto.CustomerData{
		ID: persisted.ID,
		Code: persisted.Code,
		Name: persisted.Name,
	}, nil

}
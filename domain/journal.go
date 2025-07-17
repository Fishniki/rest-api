package domain

import (
	"context"
	"database/sql"
	"rest-api/dto"
)



const (
	JournalStatusInProgres = "IN_PROGRESS"
	JournalStatusCompled = "COMPLETED"
)


type Journal struct {
	Id         string `json:"id"`
	BookId     string `json:"book_id"`
	StockCode  string `json:"stock_code"`
	CustomerId string `json:"customer_id"`
	Status     string `json:"status"`
	BorrowedAt sql.NullTime `json:"borrowed_at"`
	ReturnedAt sql.NullTime `json:"returned_at"`
}

type JournalSearch struct{
	CustomerId string
	Status string
}

type JournalRepository interface {
	Find(ctx context.Context, se JournalSearch) ([]Journal, error)
	FindById(ctx context.Context, id string) (Journal, error)
	Save(ctx context.Context, journal *Journal) error
	Update(ctx context.Context, journal *Journal) error
}


type JournalService interface {
	Index(ctx context.Context, se JournalSearch) ([]dto.JournalData, error)
	Create(ctx context.Context, req dto.CreateJournalRequest) error
	Return(ctx context.Context, req dto.ReturnJournalRequest) error
}
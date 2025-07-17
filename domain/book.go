package domain

import (
	"context"
	"database/sql"
	"rest-api/dto"
)

type Book struct {
	Id          string         `db:"id"`
	Title       string         `db:"title"`
	Isbn        string         `db:"isbn"`
	Description string         `db:"description"`
	CoverId     sql.NullString `db:"cover_id"`
	CreatedAt   sql.NullTime   `db:"created_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`
	DeletedAt   sql.NullTime   `db:"deleted_at"`
}

type BookRepository interface {
	FindAll(ctx context.Context) ([]Book, error)
	FindByID(ctx context.Context, id string) (Book, error)
	FindByIDs(ctx context.Context, id []string) ([]Book, error)
	Save(ctx context.Context, b *Book) error
	Update(ctx context.Context, b *Book) error
	Delete(ctx context.Context, id string) error
}

type BookService interface {
	Index(ctx context.Context) ([]dto.BookData, error)
	Show(ctx context.Context, id string) (dto.BookShowData, error)
	Create(ctx context.Context, req dto.CreateBookRequest) error
	Update(ctx context.Context, req dto.UpdateBookRequest) error
	Delete(ctx context.Context, id string) error
}

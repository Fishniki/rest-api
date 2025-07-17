package repository

import (
	"context"
	"database/sql"
	"rest-api/domain"

	"github.com/doug-martin/goqu/v9"
)

type BookStockRepository struct {
	db *goqu.Database
}

func NewBookStock(con *sql.DB) domain.BookStockRepository {
	return &BookStockRepository{
		db: goqu.New("default", con),
	}
}

func (b BookStockRepository) FindByBookId(ctx context.Context, id string) (result []domain.BookStock, err error) {
	
	dataset := b.db.From("book_stocks").Where(goqu.C("book_id").Eq(id))
	err = dataset.ScanStructsContext(ctx, &result)
	return

}

// FindByBookIdAndCode implements domain.BookStockRepository.
func (b BookStockRepository) FindByBookIdAndCode(ctx context.Context, id string, code string) (result domain.BookStock, err error) {
	
	dataset := b.db.From("book_stocks").Where(
		goqu.C("book_id").Eq(id), 
		goqu.C("code").Eq(code),
	)
	_, err = dataset.ScanStructContext(ctx, &result)
	return

}

// DeleteByBookId implements domain.BookStockRepository.
func (b BookStockRepository) DeleteByBookId(ctx context.Context, id string) error {
	
	executor := b.db.Delete("book_stocks").Where(
		goqu.C("book_id").Eq(id),
	).Executor()

	_, err := executor.ExecContext(ctx)
	return  err

}

// DeleteByCodes implements domain.BookStockRepository.
func (b BookStockRepository) DeleteByCodes(ctx context.Context, codes []string) error {
	
	executor := b.db.Delete("book_stocks").Where(
		goqu.C("code").In(codes),
	).Executor()

	_, err := executor.ExecContext(ctx)
	return  err

}

// FindByBookId implements domain.BookStockRepository.

// Save implements domain.BookStockRepository.
func (b BookStockRepository) Save(ctx context.Context, data []domain.BookStock) error {
	
	executor := b.db.Insert("book_stocks").Rows(data).Executor()
	_, err := executor.ExecContext(ctx)
	return err

}

// Update implements domain.BookStockRepository.
func (b BookStockRepository) Update(ctx context.Context, stock *domain.BookStock) error {
	
	executor := b.db.Update("book_stocks").Where(
		goqu.C("code").Eq(stock.Code),
	).Set(stock).Executor()

	_, err := executor.ExecContext(ctx)
	return  err

}


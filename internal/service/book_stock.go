package service

import (
	"context"
	"errors"
	"rest-api/domain"
	"rest-api/dto"
)

type BookStockService struct {
	BookRepository      domain.BookRepository
	BookStockRepository domain.BookStockRepository
}


func NewBookStock(
	bookRepository domain.BookRepository,
	bookStockRepository domain.BookStockRepository,
) domain.BookStockService {

	return &BookStockService{
		BookRepository:      bookRepository,
		BookStockRepository: bookStockRepository,
	}

}




// Create implements domain.BookStockService.
func (b BookStockService) Create(ctx context.Context, req dto.CreateBookStockRequest) error {
	
	book, err := b.BookRepository.FindByID(ctx, req.BookId)
	if err != nil {
		return err
	}

	if book.Id == "" {
		return errors.New("data buku tidak di temukan")
	}

	stocks := make([]domain.BookStock, 0)
	for _, v := range req.Codes{

		if v == "" {
			return errors.New("Field Kode Buku Tidak Boleh Kosong")
		}

		stocks = append(stocks, domain.BookStock{
			Code: v,
			BookId: req.BookId,
			Status: domain.BookStockStatusAvailable,
		})
	}

	return b.BookStockRepository.Save(ctx, stocks)

}

// Delete implements domain.BookStockService.
func (b *BookStockService) Delete(ctx context.Context, req dto.DeleteBookStockRequest) error {
	
	return b.BookStockRepository.DeleteByCodes(ctx, req.Codes)

}


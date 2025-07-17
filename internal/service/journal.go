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

type journalService struct {
	journalRepository   domain.JournalRepository
	bookRepository      domain.BookRepository
	bookStockRepository domain.BookStockRepository
	customerRepository  domain.CustomerRepository
}

func NewJournal(journalRepository domain.JournalRepository,
	bookRepository domain.BookRepository,
	bookStockRepository domain.BookStockRepository,
	customerRepository domain.CustomerRepository) domain.JournalService {
	return &journalService{
		journalRepository:   journalRepository,
		bookRepository:      bookRepository,
		bookStockRepository: bookStockRepository,
		customerRepository: customerRepository,
	}
}


// Create implements domain.JournalService.
func (j journalService) Create(ctx context.Context, req dto.CreateJournalRequest) error {
	
	book, err := j.bookRepository.FindByID(ctx, req.BookId)
	if err != nil {
		return err
	}

	if book.Id == "" {
		return errors.New("data buku tidak di temukan")
	}

	stock, err := j.bookStockRepository.FindByBookIdAndCode(ctx, book.Id, req.BookStock)
	if err != nil {
		return err
	}

	if stock.Code == "" {
		return errors.New("data buku tidak di temukan")
	}

	if stock.Status != domain.BookStockStatusAvailable{
		return errors.New("stock buku sudah di pinjam sebelumnya")
	}

	journal := domain.Journal{
		Id: uuid.NewString(),
		BookId: req.BookId,
		StockCode: req.BookStock,
		CustomerId: req.CustomerId,
		Status: domain.JournalStatusInProgres,
		BorrowedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}

	err = j.journalRepository.Save(ctx, &journal)
	if err != nil {
		return err
	}

	stock.Status = domain.BookStockStatusBorrowed
	stock.BorrowerAt = journal.BorrowedAt
	stock.BorrowerId = sql.NullString{Valid: true, String: journal.CustomerId}
	
	return  j.bookStockRepository.Update(ctx, &stock)

}

// Index implements domain.JournalService.
func (j journalService) Index(ctx context.Context, se domain.JournalSearch) ([]dto.JournalData, error) {
	
	journals, err :=j.journalRepository.Find(ctx, se)
	if err != nil {
		return nil, err
	}

	customerId := make([]string, 0)
	bookId := make([]string, 0)

	for _, v :=	range journals {
		customerId = append(customerId, v.CustomerId)
		bookId = append(bookId, v.BookId)
	}

	customers := make(map[string] domain.Customer)
	if len(customerId) > 0 {
		customerDb, _ :=j.customerRepository.FindByIDs(ctx, customerId)
		for _, v := range customerDb {
			customers[v.ID] = v
		}
	}

	books := make(map[string] domain.Book)
	if len(bookId) > 0 {
		bookDb, _ := j.bookRepository.FindByIDs(ctx, bookId)
		for _, v := range bookDb {
			books[v.Id] = v
		}
	}

	result := make([]dto.JournalData, 0)
	for _, v := range journals {

		book := dto.BookData{}
		if v2, e := books[v.BookId]; e {
			book = dto.BookData{
				Id: v2.Id,
				Isbn: v2.Isbn,
				Title: v2.Title,
				Description: v2.Description,
			}
		}

		customer := dto.CustomerData{}
		if v2, e := customers[v.CustomerId]; e {
			customer = dto.CustomerData{
				ID: v2.ID,
				Code: v2.Code,
				Name: v2.Name,
			}
		}

		result = append(result, dto.JournalData{
			Id: v.Id,
			BookStock: v.StockCode,
			Book: book,
			Customer: customer,

			BorrowedAt: v.BorrowedAt.Time.String(),
			ReturnedAt: v.ReturnedAt.Time.String(),

		})
	}

	return result, nil

}

// Return implements domain.JournalService.
func (j journalService) Return(ctx context.Context, req dto.ReturnJournalRequest) error {
	
	journal, err := j.journalRepository.FindById(ctx, req.JournalId)
	if err != nil {
		return  err
	}

	if journal.Id == "" {
		return errors.New("data journal tidak di temukan")
	}

	stock, err := j.bookStockRepository.FindByBookIdAndCode(ctx, journal.BookId, journal.StockCode)
	if err != nil {
		return err
	}

	if stock.Code != "" {
		stock.Status = domain.BookStockStatusAvailable
		stock.BorrowerId = sql.NullString{Valid: false}
		stock.BorrowerAt = sql.NullTime{Valid: false}
		err = j.bookStockRepository.Update(ctx, &stock)
		if err != nil {
			return err
		}
	}

	journal.Status = domain.JournalStatusCompled
	journal.ReturnedAt = sql.NullTime{Valid: true, Time: time.Now()}

	return j.journalRepository.Update(ctx, &journal)

}

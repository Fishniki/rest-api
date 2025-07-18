package service

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"path"
	"rest-api/domain"
	"rest-api/dto"
	"rest-api/internal/config"
	"time"

	"github.com/google/uuid"
)

type BookService struct {
	cnf                 *config.Config
	BookRepository      domain.BookRepository
	BookStockRepository domain.BookStockRepository
	MediaRepository     domain.MediaRepository
}

func NewBook(
	cnf *config.Config,
	bookRepository domain.BookRepository,
	bookStockRepository domain.BookStockRepository,
	mediaRepository domain.MediaRepository,
) domain.BookService {

	return &BookService{
		cnf:                 cnf,
		BookRepository:      bookRepository,
		BookStockRepository: bookStockRepository,
		MediaRepository:     mediaRepository,
	}

}

// Create implements domain.BookService.
func (b BookService) Create(ctx context.Context, req dto.CreateBookRequest) error {

	coverId := sql.NullString{Valid: false, String: req.CoverId}
	if req.CoverId != "" {
		coverId.Valid = true
	}

	book := domain.Book{
		Id:          uuid.NewString(),
		Title:       req.Title,
		Isbn:        req.Isbn,
		Description: req.Description,
		CoverId:     coverId,
		CreatedAt:   sql.NullTime{Valid: true, Time: time.Now()},
	}

	return b.BookRepository.Save(ctx, &book)

}

// Delete implements domain.BookService.
func (b BookService) Delete(ctx context.Context, id string) error {

	persisted, err := b.BookRepository.FindByID(ctx, id)

	if err != nil {
		return err
	}

	if persisted.Id == "" {
		return errors.New("data buku tidak di temukan")
	}

	err = b.BookRepository.Delete(ctx, persisted.Id)
	if err != nil {
		return err
	}

	return b.BookStockRepository.DeleteByBookId(ctx, persisted.Id)

}

// Index implements domain.BookService.
func (b BookService) Index(ctx context.Context) ([]dto.BookData, error) {

	result, err := b.BookRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	coverId := make([]string, 0)
	for _, v := range result {
		if v.CoverId.Valid {
			coverId = append(coverId, v.CoverId.String)
		}
	}

	covers := make(map[string]string)
	if len(coverId) > 0 {
		coverDb, _ := b.MediaRepository.FindByIds(ctx, coverId)
		for _, v := range coverDb {
			covers[v.Id] = path.Join(b.cnf.Server.Asset, v.Path)
		}
		
			log.Println("coverId:", coverId)
			log.Println("coverDb:", coverDb)
	}


	var data []dto.BookData
	for _, v := range result {
		var coverUrl string
		if v2, e := covers[v.CoverId.String]; e {
			coverUrl = v2
		}
		data = append(data, dto.BookData{
			Id:          v.Id,
			Title:       v.Title,
			Isbn:        v.Isbn,
			CoverUrl:    coverUrl,
			Description: v.Description,
		})
	}

	return data, nil

}

// Show implements domain.BookService.
func (b BookService) Show(ctx context.Context, id string) (dto.BookShowData, error) {

	data, err := b.BookRepository.FindByID(ctx, id)
	if err != nil {
		return dto.BookShowData{}, err
	}

	if data.Id == "" {
		return dto.BookShowData{}, errors.New("data buku tidak di temukan atau id salah")
	}

	stocks, err := b.BookStockRepository.FindByBookId(ctx, data.Id)
	if err != nil {
		return dto.BookShowData{}, err
	}

	stocksData := make([]dto.BookStockData, 0)

	for _, v := range stocks {
		stocksData = append(stocksData, dto.BookStockData{
			Code:   v.Code,
			Status: v.Status,
		})
	}

	var coverUrl string
	if data.CoverId.Valid {
		cover, _ := b.MediaRepository.FindById(ctx, data.CoverId.String)
		if cover.Path != "" {
			coverUrl = path.Join(b.cnf.Server.Asset, cover.Path)
		}
	}

	return dto.BookShowData{
		Id:          data.Id,
		Isbn:        data.Isbn,
		Title:       data.Title,
		CoverUrl:    coverUrl,
		Description: data.Description,
		Stocks:      stocksData,
	}, nil

}

// Update implements domain.BookService.
func (b BookService) Update(ctx context.Context, req dto.UpdateBookRequest) error {

	persisted, err := b.BookRepository.FindByID(ctx, req.Id)

	if err != nil {
		return err
	}

	if persisted.Id == "" {
		return errors.New("data buku tidak di temukan")
	}

	coverId := sql.NullString{Valid: false, String: req.CoverId}
	if req.CoverId != "" {
		coverId.Valid = true
	}

	persisted.Isbn = req.Isbn
	persisted.Title = req.Title
	persisted.Description = req.Description
	persisted.CoverId = coverId
	persisted.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}

	return b.BookRepository.Update(ctx, &persisted)

}

package services

import (
	"context"
	"errors"
	"time"

	"ProjetoGobid/internal/store/pgstore"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewProductService(pool *pgxpool.Pool) ProductService {
	return ProductService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

func (ps ProductService) CreateProduct(
	ctx context.Context,
	sellerID uuid.UUID,
	productName,
	description string,
	basePrice float64,
	auctionEnd time.Time,
) (uuid.UUID, error) {
	id, err := ps.queries.CreateProduct(ctx, pgstore.CreateProductParams{
		SellerID:    sellerID,
		ProductName: productName,
		Description: description,
		Baseprice:   basePrice,
		AuctionEnd:  auctionEnd,
	})
	if err != nil {
		return uuid.UUID{}, err
	}
	return id, nil
}

func (ps ProductService) GetProduct(ctx context.Context) ([]pgstore.Product, error) {
	data, err := ps.queries.GetProduct(ctx)
	if err != nil {
		return data, err
	}
	return data, nil
}

var ErrProductNotFound = errors.New("product not found")

func (ps ProductService) GetProductById(ctx context.Context, id uuid.UUID) (pgstore.Product, error) {
	data, err := ps.queries.GetProductById(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return pgstore.Product{}, ErrProductNotFound
		}
		return pgstore.Product{}, err
	}
	return data, nil
}

package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lil-oren/cron/internal/dto"
)

type (
	ReviewRepository interface {
		RateOfProduct(ctx context.Context, code string) (*dto.RateOfProductModel, error)
	}
	reviewRepository struct {
		db *sqlx.DB
	}
)

func (r *reviewRepository) RateOfProduct(ctx context.Context, code string) (*dto.RateOfProductModel, error) {
	rate := new(dto.RateOfProductModel)
	query := `
		SELECT 
			count(id) AS rate_count,
			COALESCE(sum(rating), 0) AS rate_sum
		FROM reviews r
		WHERE r.product_code = $1
	`
	if err := r.db.GetContext(ctx, rate, query, code); err != nil {
		return nil, err
	}
	return rate, nil
}

func NewReviewRepository(db *sqlx.DB) ReviewRepository {
	return &reviewRepository{
		db: db,
	}
}

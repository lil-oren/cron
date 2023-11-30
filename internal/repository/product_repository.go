package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lil-oren/cron/internal/dto"
)

type (
	ProductRepository interface {
		FindRecommended(ctx context.Context) ([]dto.HomePageProductModel, error)
	}
	productRepository struct {
		db *sqlx.DB
	}
)

func (r *productRepository) FindRecommended(ctx context.Context) ([]dto.HomePageProductModel, error) {
	e := make([]dto.HomePageProductModel, 0)

	query := `
	SELECT
		p.product_code,
		p.thumbnail_url AS media_url,
		p.name,
		pv.price, 
		(pv.price - (pv.price * pv.discount / 100)) as discounted_price,
		pv.discount,
		(CASE 
			WHEN mp.count_purchased IS NULL THEN 0
			ELSE mp.count_purchased
		END) AS total_sold,
		d.name AS shop_location,
		s.name AS shop_name
	FROM products p
	LEFT JOIN 
		(
			SELECT 
				DISTINCT ON (pv.product_id) pv.product_id, 
				pv.discount, 
				pv.price 
			FROM 
				product_variants pv
		) pv ON p.id = pv.product_id
	LEFT JOIN 
	(
		SELECT
			od.product_code,
			count(od.order_id) AS count_purchased
		FROM order_details od 
		GROUP BY od.product_code
	) mp ON 
	p.product_code = mp.product_code
	LEFT JOIN shops s
		ON s.account_id = p.seller_id
	LEFT JOIN account_addresses aa
		ON aa.account_id = p.seller_id
		AND aa.is_shop
	LEFT JOIN districts d
		ON d.id = aa.district_id
	LIMIT 18;
	`

	err := r.db.SelectContext(ctx, &e, query)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

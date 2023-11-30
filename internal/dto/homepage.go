package dto

import (
	"github.com/shopspring/decimal"
)

type (
	HomePageProductModel struct {
		ImageUrl        string          `db:"media_url"`
		ProductCode     string          `db:"product_code"`
		Name            string          `db:"name"`
		Price           decimal.Decimal `db:"price"`
		DiscountedPrice decimal.Decimal `db:"discounted_price"`
		Discount        float32         `db:"discount"`
		TotalSold       int             `db:"total_sold"`
		ShopName        string          `db:"shop_name"`
		ShopLocation    string          `db:"shop_location"`
	}
	HomePageProductResponseBody struct {
		ImageUrl        string  `json:"image_url"`
		ProductCode     string  `json:"product_code"`
		Name            string  `json:"name"`
		Price           float64 `json:"price"`
		DiscountedPrice float64 `json:"discounted_price"`
		Discount        float32 `json:"discount"`
		TotalSold       int     `json:"total_sold"`
		ShopName        string  `json:"shop_name"`
		ShopLocation    string  `json:"shop_location"`
		Rating          float64 `json:"rating"`
	}
	HomePageCategoryResponseBody struct {
		CategoryID   int64  `json:"category_id"`
		CategoryName string `json:"category_name"`
		ImageURL     string `json:"image_url"`
	}
)

package usecase

import (
	"context"
	"math"

	"github.com/lil-oren/cron/internal/dependency"
	"github.com/lil-oren/cron/internal/dto"
	"github.com/lil-oren/cron/internal/repository"
	"github.com/lil-oren/cron/internal/shared"
)

type (
	HomepageUsecase interface {
		GetRecommendedProducts(ctx context.Context) error
	}
	homepageUsecase struct {
		pr     repository.ProductRepository
		cr     repository.CacheRepository
		rr     repository.ReviewRepository
		logger dependency.Logger
	}
)

func (uc *homepageUsecase) GetRecommendedProducts(ctx context.Context) error {
	e, err := uc.pr.FindRecommended(ctx)
	if err != nil {
		uc.logger.Infof("Failed to update recommended products", nil)
		return err
	}
	res := make([]dto.HomePageProductResponseBody, 0)
	for _, val := range e {
		rate, err := uc.rr.RateOfProduct(ctx, val.ProductCode)
		if err != nil {
			uc.logger.Infof("Failed to update recommended products", nil)
			return err
		}
		rating := rate.RateSum / rate.RateCount
		if math.IsNaN(rating) {
			rating = 0
		}
		v := dto.HomePageProductResponseBody{
			Name:            val.Name,
			ImageUrl:        val.ImageUrl,
			Discount:        val.Discount,
			Price:           val.Price.InexactFloat64(),
			DiscountedPrice: val.DiscountedPrice.InexactFloat64(),
			TotalSold:       val.TotalSold,
			ShopName:        val.ShopName,
			ShopLocation:    val.ShopLocation,
			Rating:          shared.RoundFloat(rating, 1),
			ProductCode:     val.ProductCode,
		}
		res = append(res, v)
	}

	err = uc.cr.SetRecommendedProduct(ctx, res)
	if err != nil {
		uc.logger.Infof("Failed to update recommended products", nil)
		return err
	}

	uc.logger.Infof("Updated Recommended Products", nil)
	return nil
}

func NewHomepageUsecase(
	pr repository.ProductRepository,
	cr repository.CacheRepository,
	rr repository.ReviewRepository,
	logger dependency.Logger,
) HomepageUsecase {
	return &homepageUsecase{
		pr:     pr,
		cr:     cr,
		rr:     rr,
		logger: logger,
	}
}

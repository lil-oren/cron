package job

import (
	"context"

	"github.com/go-co-op/gocron"
	"github.com/lil-oren/cron/internal/usecase"
)

type HomepageJob struct {
	huc usecase.HomepageUsecase
	s   *gocron.Scheduler
}

func (j HomepageJob) UpdateRecommendedProduct() {
	j.s.Every(5).Minutes().Tag("recom").Do(j.huc.GetRecommendedProducts, context.TODO())
}

func NewHomepageJob(huc usecase.HomepageUsecase, s *gocron.Scheduler) HomepageJob {
	return HomepageJob{
		huc: huc,
		s:   s,
	}
}

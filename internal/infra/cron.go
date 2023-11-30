package infra

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/lil-oren/cron/internal/dependency"
	"github.com/lil-oren/cron/internal/job"
	"github.com/lil-oren/cron/internal/repository"
	"github.com/lil-oren/cron/internal/usecase"
)

type (
	server struct {
		v            *validator.Validate
		repositories repositories
		usecases     usecases
		jobs         jobs
		cfg          dependency.Config
		sc           *gocron.Scheduler
	}

	repositories struct {
		productRepository repository.ProductRepository
		reviewRepository  repository.ReviewRepository
		cacheRepository   repository.CacheRepository
	}

	usecases struct {
		homepageUsecase usecase.HomepageUsecase
	}

	jobs struct {
		homepageJob job.HomepageJob
	}
)

func (s *server) initRepository(db *sqlx.DB, rd *redis.Client, cfg dependency.Config) {
	s.repositories.productRepository = repository.NewProductRepository(db)
	s.repositories.cacheRepository = repository.NewCacheRepository(rd, s.cfg)
	s.repositories.reviewRepository = repository.NewReviewRepository(db)
}

func (s *server) initUsecase(rd *redis.Client, logger dependency.Logger) {
	s.usecases.homepageUsecase = usecase.NewHomepageUsecase(
		s.repositories.productRepository,
		s.repositories.cacheRepository,
		s.repositories.reviewRepository,
		logger,
	)
}

func (s *server) initJobs() {
	s.jobs.homepageJob = job.NewHomepageJob(s.usecases.homepageUsecase, s.sc)
}

func (s *server) startCronJob(cfg dependency.Config, logger dependency.Logger) *gocron.Scheduler {
	s.initJobs()

	go func() {
		logger.Infof("CronJob is running", nil)
		s.jobs.homepageJob.UpdateRecommendedProduct()
		s.sc.StartBlocking()
	}()

	return s.sc
}

func initGracefulShutdown(goCron *gocron.Scheduler, cfg dependency.Config) {
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.App.GracefulTimeout)*time.Second)
	defer cancel()

	// stop cronjob
	goCron.Stop()

	<-ctx.Done()
	log.Println("Server exiting")
}

func InitApp(db *sqlx.DB, rc *redis.Client, cfg dependency.Config, logger dependency.Logger) {
	s := server{
		v:   validator.New(),
		cfg: cfg,
		sc:  gocron.NewScheduler(time.Local),
	}

	s.initRepository(db, rc, cfg)
	s.initUsecase(rc, logger)

	goCron := s.startCronJob(cfg, logger)

	initGracefulShutdown(goCron, cfg)
}

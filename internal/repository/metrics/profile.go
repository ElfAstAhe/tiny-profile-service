package metrics

import (
	"context"
	"time"

	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/infra/metrics"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-profile-service/internal/domain"
)

type ProfileMetricsRepository struct {
	*repository.BaseOwnedMetricsRepository[*domain.Profile, string, string]

	repo domain.ProfileRepository
}

var _ libdomain.OwnedRepository[*domain.Profile, string, string] = (*ProfileMetricsRepository)(nil)
var _ domain.ProfileRepository = (*ProfileMetricsRepository)(nil)

func NewProfileMetricsRepository(repo domain.ProfileRepository) *ProfileMetricsRepository {
	return &ProfileMetricsRepository{
		repo:                       repo,
		BaseOwnedMetricsRepository: repository.NewBaseOwnedMetricsRepository[*domain.Profile, string, string]("ProfileRepository", repo),
	}
}

func (prm *ProfileMetricsRepository) FindByID(ctx context.Context, id string) (res *domain.Profile, err error) {
	defer func(start time.Time) {
		metrics.ObserveRepositoryOp(prm.GetRepositoryName(), "FindByID", err, start)
	}(time.Now())

	return prm.repo.FindByID(ctx, id)
}

func (prm *ProfileMetricsRepository) FindByUserID(ctx context.Context, userID string) (res *domain.Profile, err error) {
	defer func(start time.Time) {
		metrics.ObserveRepositoryOp(prm.GetRepositoryName(), "FindByUserID", err, start)
	}(time.Now())

	return prm.repo.FindByUserID(ctx, userID)
}

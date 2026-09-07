package metrics

import (
	"context"
	"time"

	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/infra/metrics"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-profile-service/internal/domain"
)

type PersonMetricsRepository struct {
	*repository.BaseCRUDMetricsRepository[*domain.Person, string]

	repo domain.PersonRepository
}

var _ libdomain.CRUDRepository[*domain.Person, string] = (*PersonMetricsRepository)(nil)
var _ domain.PersonRepository = (*PersonMetricsRepository)(nil)

func NewPersonsMetricsRepository(repo domain.PersonRepository) *PersonMetricsRepository {
	return &PersonMetricsRepository{
		repo:                      repo,
		BaseCRUDMetricsRepository: repository.NewBaseCRUDMetricsRepository[*domain.Person, string]("PersonRepository", repo),
	}
}

func (pm *PersonMetricsRepository) FindByExternalID(ctx context.Context, externalID string) (res *domain.Person, err error) {
	defer func(start time.Time) {
		metrics.ObserveRepositoryOp(pm.GetRepositoryName(), "FindByExternalID", err, start)
	}(time.Now())

	return pm.repo.FindByExternalID(ctx, externalID)
}

package telemetry

import (
	"context"
	"fmt"

	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-profile-service/internal/domain"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type PersonRepository struct {
	*repository.BaseCRUDTraceRepository[*domain.Person, string]

	repo domain.PersonRepository
}

var _ libdomain.CRUDRepository[*domain.Person, string] = (*PersonRepository)(nil)
var _ domain.PersonRepository = (*PersonRepository)(nil)

func NewPersonRepository(repo domain.PersonRepository) *PersonRepository {
	return &PersonRepository{
		repo:                    repo,
		BaseCRUDTraceRepository: repository.NewBaseCRUDTraceRepository[*domain.Person, string]("PersonRepository", repo),
	}
}

func (por *PersonRepository) FindByExternalID(ctx context.Context, externalID string) (*domain.Person, error) {
	ctx, span := por.StartSpan(ctx, fmt.Sprintf("%s.FindByExternalID", por.GetRepositoryName()))
	defer span.End()

	span.SetAttributes(
		attribute.String("param.external_id", fmt.Sprintf("%v", externalID)),
	)

	res, err := por.repo.FindByExternalID(ctx, externalID)
	if err != nil {
		span.AddEvent("FindByExternalID_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return por.GetNilEntity(), err
	}

	return res, nil
}

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

type ProfileRepository struct {
	*repository.BaseOwnedTraceRepository[*domain.Profile, string, string]

	repo domain.ProfileRepository
}

var _ libdomain.OwnedRepository[*domain.Profile, string, string] = (*ProfileRepository)(nil)
var _ domain.ProfileRepository = (*ProfileRepository)(nil)

func NewProfileRepository(repo domain.ProfileRepository) *ProfileRepository {
	return &ProfileRepository{
		repo:                     repo,
		BaseOwnedTraceRepository: repository.NewBaseOwnedTraceRepository[*domain.Profile, string, string]("ProfileRepository", repo),
	}
}

func (ptr *ProfileRepository) FindByID(ctx context.Context, id string) (*domain.Profile, error) {
	ctx, span := ptr.StartSpan(ctx, fmt.Sprintf("%s.FindByID", ptr.GetRepositoryName()))
	defer span.End()

	span.SetAttributes(
		attribute.String("param.id", fmt.Sprintf("%v", id)),
	)

	res, err := ptr.repo.FindByID(ctx, id)
	if err != nil {
		span.AddEvent("FindByID_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return ptr.GetNilEntity(), err
	}

	return res, nil
}

func (ptr *ProfileRepository) FindByUserID(ctx context.Context, userID string) (*domain.Profile, error) {
	ctx, span := ptr.StartSpan(ctx, fmt.Sprintf("%s.FindByUserID", ptr.GetRepositoryName()))
	defer span.End()

	span.SetAttributes(
		attribute.String("param.user_id", fmt.Sprintf("%v", userID)),
	)

	res, err := ptr.repo.FindByUserID(ctx, userID)
	if err != nil {
		span.AddEvent("FindByUserID_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return ptr.GetNilEntity(), err
	}

	return res, nil
}

package postgres

import (
	"context"
	"strings"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-profile-service/internal/domain"
	domrepo "github.com/ElfAstAhe/tiny-profile-service/internal/repository"
)

type ProfilePgRepository struct {
	*repository.BaseOwnedRepository[*domain.Profile, string, string]
}

var _ libdomain.OwnedRepository[*domain.Profile, string, string] = (*ProfilePgRepository)(nil)
var _ domain.ProfileRepository = (*ProfilePgRepository)(nil)

func NewProfileRepository(
	executor db.Executor,
	decipher db.ErrorDecipher,
) (*ProfilePgRepository, error) {
	// new instance
	res := &ProfilePgRepository{}
	// sql builders
	queryBuilders := repository.NewBaseOwnedQueryBuildersBuilder().NewInstance().
		WithListAll(func() string {
			return sqlProfileListAll
		}).
		WithListAllByOwners(func() string {
			return sqlProfileListAllByOwners
		}).
		WithList(func() string {
			return sqlProfileList
		}).
		WithFind(func() string {
			return sqlProfileFind
		}).
		WithCreate(func() string {
			return sqlProfileCreate
		}).
		WithChange(func() string {
			return sqlProfileChange
		}).
		WithDelete(func() string {
			return sqlProfileDelete
		}).
		WithDeleteAll(func() string {
			return sqlProfileDeleteAll
		}).
		Build()

	// callbacks
	callbacks, err := repository.NewBaseRepositoryCallbacksBuilder[*domain.Profile, string]().NewInstance().
		WithEntityScanner(res.entityScanner).
		WithNewEntityFactory(domain.NewEmptyProfile).
		WithAfterListYield(res.afterListYield).
		Build()
	if err != nil {
		return nil, err
	}
	// base owned
	base, err := repository.NewBaseOwnedRepository[*domain.Profile, string, string](
		executor,
		decipher,
		repository.NewEntityInfo("profiles", "Profile"),
		queryBuilders,
		callbacks,
		repository.LinkStrategyOneToMany,
		nil,
	)
	if err != nil {
		return nil, err
	}

	res.BaseOwnedRepository = base

	return res, nil
}

func (pp *ProfilePgRepository) FindByID(ctx context.Context, id string) (*domain.Profile, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errs.NewInvalidArgumentError("id", "must not be blank")
	}

	res, err := pp.GetHelper().Get(ctx, domrepo.SourceLabelProfileFindByID, sqlProfileFindByID, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pp *ProfilePgRepository) FindByUserID(ctx context.Context, userID string) (*domain.Profile, error) {
	if strings.TrimSpace(userID) == "" {
		return nil, errs.NewInvalidArgumentError("userID", "must not be blank")
	}

	res, err := pp.GetHelper().Get(ctx, domrepo.SourceLabelProfileFindByUserID, sqlProfileFindByUserID, userID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pp *ProfilePgRepository) entityScanner(
	scanner repository.Scannable,
	sourceLabel string,
	dest *domain.Profile,
	params ...any,
) error {
	switch sourceLabel {
	default:
		return scanner.Scan(
			&dest.ID,
			&dest.UserID,
			&dest.PersonID,
			&dest.TimeZone,
			&dest.Lang,
			&dest.Active,
			&dest.Deleted,
			&dest.CreatedAt,
			&dest.UpdatedAt,
		)
	}
}

func (pp *ProfilePgRepository) afterListYield(entity *domain.Profile, params ...any) (*domain.Profile, bool, error) {
	if entity.IsDeleted() {
		return nil, false, errs.NewDalSoftDeletedError(pp.GetHelper().GetInfo().Entity, entity.GetID())
	}

	return entity, true, nil
}

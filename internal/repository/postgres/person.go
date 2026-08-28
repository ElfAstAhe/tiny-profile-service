package postgres

import (
	"context"
	"database/sql"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-profile-service/internal/domain"
)

type PersonPgRepository struct {
	*repository.BaseCRUDRepository[*domain.Person, string]
}

var _ libdomain.CRUDRepository[*domain.Person, string] = (*PersonPgRepository)(nil)
var _ domain.PersonRepository = (*PersonPgRepository)(nil)

func NewPersonRepository(executor db.Executor, errDecipher db.ErrorDecipher) (*PersonPgRepository, error) {
	res := &PersonPgRepository{}

	// sql builders
	queryBuilders := repository.NewBaseCRUDQueryBuildersBuilder().NewInstance().
		WithFind(func() string {
			return sqlPersonFind
		}).
		WithList(func() string {
			return sqlPersonList
		}).
		WithCreate(func() string {
			return sqlPersonCreate
		}).
		WithChange(func() string {
			return sqlPersonChange
		}).
		WithDelete(func() string {
			return sqlPersonDelete
		}).
		Build()
	// callbacks
	callbacks, _ := repository.NewBaseRepositoryCallbacksBuilder[*domain.Person, string]().NewInstance().
		WithEntityScanner(res.entityScanner).
		WithNewEntityFactory(domain.NewEmptyPerson).
		WithAfterListYield(res.afterListYield).
		WithValidateCreate(res.validateCreate).
		WithBeforeCreate(res.beforeCreate).
		WithCreator(res.creator).
		WithValidateChange(res.validateChange).
		WithBeforeChange(res.beforeChange).
		WithChanger(res.changer).
		Build()
	// base CRUD
	base, err := repository.NewBaseCRUDRepository[*domain.Person, string](
		executor,
		errDecipher,
		repository.NewEntityInfo("persons", "Person"),
		queryBuilders,
		callbacks,
	)
	if err != nil {
		return nil, errs.NewCommonError("error create PersonPgRepository", err)
	}

	res.BaseCRUDRepository = base

	return res, nil
}

func (p *PersonPgRepository) FindByExternalID(ctx context.Context, externalID string) (*domain.Person, error) {

}

func (p *PersonPgRepository) entityScanner(
	scanner repository.Scannable,
	sourceLabel string,
	entity *domain.Person,
	params ...any,
) error {
	switch sourceLabel {
	default:
		return scanner.Scan(
			&entity.ID,

			&entity.Source,
			&entity.EventDate,
			&entity.Event,
			&entity.Status,
			&entity.RequestID,
			&entity.TraceID,
			&entity.Username,
			&entity.AccessToken,
			&entity.RefreshToken,
			&entity.CreatedAt,
			&entity.UpdatedAt,
		)
	}
}

func (p *PersonPgRepository) afterListYield(entity *domain.Person, params ...any) (*domain.Person, bool, error) {
	return entity, true, nil
}

func (p *PersonPgRepository) validateCreate(entity *domain.Person, params ...any) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "person entity is nil")
	}

	return entity.ValidateCreate()
}

func (p *PersonPgRepository) beforeCreate(entity *domain.Person, params ...any) error {
	if err := entity.BeforeCreate(); err != nil {
		return errs.NewDalError("PersonPgRepository.beforeCreate", "before create failed", err)
	}

	return nil
}

func (p *PersonPgRepository) creator(ctx context.Context, querier db.Querier, entity *domain.Person, params ...any) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, p.GetQueryBuilders().GetCreate()(),
		entity.ID,

		entity.Source,
		entity.EventDate,
		entity.Event,
		entity.Status,
		entity.RequestID,
		entity.TraceID,
		entity.Username,
		entity.AccessToken,
		entity.RefreshToken,
		entity.CreatedAt,
	), nil
}

func (p *PersonPgRepository) validateChange(entity *domain.Person, params ...any) error {
	if entity == nil {
		return errs.NewInvalidArgumentError("entity", "person entity is nil")
	}

	return entity.ValidateChange()
}

func (p *PersonPgRepository) beforeChange(entity *domain.Person, params ...any) error {
	if err := entity.BeforeChange(); err != nil {
		return errs.NewDalError("PersonPgRepository.beforeChange", "before change failed", err)
	}

	return nil
}

func (p *PersonPgRepository) changer(ctx context.Context, querier db.Querier, entity *domain.Person, params ...any) (*sql.Row, error) {
	return querier.QueryRowContext(ctx, p.GetQueryBuilders().GetChange()(),
		entity.ID,

		entity.Source,
		entity.EventDate,
		entity.Event,
		entity.Status,
		entity.RequestID,
		entity.TraceID,
		entity.Username,
		entity.AccessToken,
		entity.RefreshToken,
		entity.UpdatedAt,
	), nil
}

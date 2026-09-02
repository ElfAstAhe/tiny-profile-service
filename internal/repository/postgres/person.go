package postgres

import (
	"context"
	"database/sql"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	librepo "github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-profile-service/internal/domain"
	"github.com/ElfAstAhe/tiny-profile-service/internal/repository"
)

type PersonPgRepository struct {
	*librepo.BaseCRUDRepository[*domain.Person, string]
	profileRepo domain.ProfileRepository
}

var _ libdomain.CRUDRepository[*domain.Person, string] = (*PersonPgRepository)(nil)
var _ domain.PersonRepository = (*PersonPgRepository)(nil)

func NewPersonRepository(
	executor db.Executor,
	errDecipher db.ErrorDecipher,
	profileRepo domain.ProfileRepository,
) (*PersonPgRepository, error) {
	res := &PersonPgRepository{
		profileRepo: profileRepo,
	}

	// sql builders
	queryBuilders := librepo.NewBaseCRUDQueryBuildersBuilder().NewInstance().
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
	callbacks, _ := librepo.NewBaseRepositoryCallbacksBuilder[*domain.Person, string]().NewInstance().
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
	base, err := librepo.NewBaseCRUDRepository[*domain.Person, string](
		executor,
		errDecipher,
		librepo.NewEntityInfo("persons", "Person"),
		queryBuilders,
		callbacks,
	)
	if err != nil {
		return nil, errs.NewCommonError("error create PersonPgRepository", err)
	}

	res.BaseCRUDRepository = base

	return res, nil
}

func (p *PersonPgRepository) Find(ctx context.Context, id string) (*domain.Person, error) {
	res, err := p.BaseCRUDRepository.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	profiles, err := p.profileRepo.ListAllByPersonID(ctx, id)
	if err != nil {
		return nil, err
	}
	res.Profiles = profiles

	return res, nil
}

func (p *PersonPgRepository) FindByExternalID(ctx context.Context, externalID string) (*domain.Person, error) {
	if externalID == "" {
		return nil, errs.NewInvalidArgumentError("externalID", "externalID is empty")
	}
	res, err := p.GetHelper().Get(ctx, repository.SourceLabelPersonFindByExternalID, externalID)
	if err != nil {
		return nil, err
	}
	profiles, err := p.profileRepo.ListAllByPersonID(ctx, externalID)
	if err != nil {
		return nil, err
	}
	res.Profiles = profiles

	return res, nil
}

func (p *PersonPgRepository) List(ctx context.Context, limit, offset int) ([]*domain.Person, error) {
	res, err := p.BaseCRUDRepository.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	allProfiles, err := p.profileRepo.ListAllByPersons(ctx, libdomain.EntitiesToIDList(res))
	if err != nil {
		return nil, err
	}
	for _, person := range res {
		if profiles, ok := allProfiles[person.ID]; ok {
			person.Profiles = profiles
		}
	}

	return res, nil
}

func (p *PersonPgRepository) entityScanner(
	scanner librepo.Scannable,
	sourceLabel string,
	entity *domain.Person,
	params ...any,
) error {
	switch sourceLabel {
	default:
		return scanner.Scan(
			&entity.ID,
			&entity.ExternalID,
			&entity.LastName,
			&entity.FirstName,
			&entity.Patronymic,
			&entity.Birthday,
			&entity.Department,
			&entity.Position,
			&entity.Status,
			&entity.AvatarURL,
			&entity.Active,
			&entity.Deleted,
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
		entity.ExternalID,
		entity.LastName,
		entity.FirstName,
		entity.Patronymic,
		entity.Birthday,
		entity.Department,
		entity.Position,
		entity.Status,
		entity.AvatarURL,
		entity.Active,
		entity.Deleted,
		entity.CreatedAt,
		entity.UpdatedAt,
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
		entity.ExternalID,
		entity.LastName,
		entity.FirstName,
		entity.Patronymic,
		entity.Birthday,
		entity.Department,
		entity.Position,
		entity.Status,
		entity.AvatarURL,
		entity.Active,
		entity.UpdatedAt,
	), nil
}

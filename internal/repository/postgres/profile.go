package postgres

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-profile-service/internal/domain"
)

type ProfileRepository struct {
	*repository.BaseCRUDRepository[*domain.Profile, string]
}

var _ libdomain.CRUDRepository[*domain.Profile, string] = (*ProfileRepository)(nil)
var _ domain.ProfileRepository = (*ProfileRepository)(nil)

func NewProfileRepository(
	executor *db.Executor,
	decipher *db.ErrorDecipher,
) (*ProfileRepository, error) {
	// new instance
	res := &ProfileRepository{}
	// sql builders
	// callbacks
	// base CRUD

	return res, nil
}

func (p ProfileRepository) FindByUserID(ctx context.Context, userID string) (*domain.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileRepository) ListAllByPersonID(ctx context.Context, personID string) ([]*domain.Profile, error) {
	//TODO implement me
	panic("implement me")
}

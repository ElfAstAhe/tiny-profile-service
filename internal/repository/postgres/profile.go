package postgres

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-profile-service/internal/domain"
)

type ProfilePgRepository struct {
	*repository.BaseCRUDRepository[*domain.Profile, string]
}

var _ libdomain.CRUDRepository[*domain.Profile, string] = (*ProfilePgRepository)(nil)
var _ domain.ProfileRepository = (*ProfilePgRepository)(nil)

func NewProfileRepository(
	executor *db.Executor,
	decipher *db.ErrorDecipher,
) (*ProfilePgRepository, error) {
	// new instance
	res := &ProfilePgRepository{}
	// sql builders
	// callbacks
	// base CRUD

	return res, nil
}

func (p *ProfilePgRepository) FindByUserID(ctx context.Context, userID string) (*domain.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProfilePgRepository) ListAllByPersonID(ctx context.Context, personID string) ([]*domain.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProfilePgRepository) ListAllByPersons(ctx context.Context, personIDs []string) (map[string][]*domain.Profile, error) {
	//TODO implement me
	panic("implement me")
}

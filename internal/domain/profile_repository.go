package domain

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/domain"
)

type ProfileRepository interface {
	domain.CRUDRepository[*Profile, string]

	FindByUserID(ctx context.Context, userID string) (*Profile, error)

	ListAllByPersonID(ctx context.Context, personID string) ([]*Profile, error)
}

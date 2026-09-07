package domain

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/domain"
)

type ProfileRepository interface {
	domain.OwnedRepository[*Profile, string, string]

	FindByID(ctx context.Context, id string) (*Profile, error)
	FindByUserID(ctx context.Context, userID string) (*Profile, error)
}

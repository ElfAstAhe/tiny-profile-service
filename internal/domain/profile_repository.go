package domain

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/domain"
)

type ProfileRepository interface {
	domain.OwnedRepository[*Profile, string, string]

	FindByUserID(ctx context.Context, userID string) (*Profile, error)
}

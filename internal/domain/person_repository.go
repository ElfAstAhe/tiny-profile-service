package domain

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/domain"
)

type PersonRepository interface {
	domain.CRUDRepository[*Person, string]

	FindByExternalID(ctx context.Context, externalID string) (*Person, error)
}

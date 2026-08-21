package domain

import (
	"github.com/ElfAstAhe/go-service-template/pkg/domain"
)

type PersonRepository interface {
	domain.CRUDRepository[*Person, string]
}

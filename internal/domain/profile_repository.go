package domain

import (
	"github.com/ElfAstAhe/go-service-template/pkg/domain"
)

type ProfileRepository interface {
	domain.OwnedRepository[*Profile, string, string]
}

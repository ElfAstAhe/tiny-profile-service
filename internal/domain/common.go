package domain

import (
	_ "time/tzdata"

	"github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/google/uuid"
)

func defaultBeforeCreate(entity domain.Entity[string]) error {
	newID, err := uuid.NewV7()
	if err != nil {
		return errs.NewBllError("defaultBeforeCreate", "generate new id", err)
	}

	entity.SetID(newID.String())

	return nil
}

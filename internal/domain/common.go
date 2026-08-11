package domain

import (
	"fmt"
	"time"
	_ "time/tzdata"

	"github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/google/uuid"
	"golang.org/x/text/language"
)

func defaultBeforeCreate(entity domain.Entity[string]) error {
	newID, err := uuid.NewV7()
	if err != nil {
		return errs.NewBllError("defaultBeforeCreate", "generate new id", err)
	}

	entity.SetID(newID.String())

	return nil
}

func validateLang(lang string) error {
	if _, err := language.Parse(lang); err != nil {
		return errs.NewBllError("validateLang", fmt.Sprintf("invalid language [%s]", lang), err)
	}

	return nil
}

func validateTimeZone(timezone string) error {
	if _, err := time.LoadLocation(timezone); err != nil {
		return errs.NewBllError("validateTimezone", fmt.Sprintf("invalid timezone [%s]", timezone), err)
	}

	return nil
}

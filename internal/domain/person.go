package domain

import (
	"hash/fnv"
	"strconv"
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	auditdomain "github.com/ElfAstAhe/tiny-audit-service/pkg/domain"
	auditrepository "github.com/ElfAstAhe/tiny-audit-service/pkg/repository"
)

type Person struct {
	ID         string    `db:"id"`
	ExternalID string    `db:"external_id"`
	LastName   string    `db:"last_name"`
	FirstName  string    `db:"first_name"`
	Patronymic string    `db:"patronymic"`
	Deleted    bool      `db:"deleted"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

var _ domain.Entity[string] = (*Person)(nil)
var _ domain.SoftDeleteEntity[bool] = (*Person)(nil)
var _ auditdomain.Auditable = (*Person)(nil)
var _ auditrepository.AuditableEntity[string] = (*Person)(nil)

func NewPerson(
	id string,
	externalID string,
	lastName string,
	firstName string,
	patronymic string,
	deleted bool,
) *Person {
	return &Person{
		ID:         id,
		ExternalID: externalID,
		LastName:   lastName,
		FirstName:  firstName,
		Patronymic: patronymic,
		Deleted:    deleted,
	}
}

func NewEmptyPerson() *Person {
	return &Person{}
}

func (pe *Person) GetID() string {
	return pe.ID
}

func (pe *Person) SetID(id string) {
	pe.ID = id
}

func (pe *Person) IsExists() bool {
	return pe.ID != ""
}

func (pe *Person) BeforeCreate() error {
	if err := defaultBeforeCreate(pe); err != nil {
		return errs.NewBllError("Person.BeforeCreate", "default before create failed", err)
	}

	if pe.CreatedAt.IsZero() {
		pe.CreatedAt = time.Now()
	}
	pe.UpdatedAt = time.Now()

	return nil
}

func (pe *Person) BeforeChange() error {
	pe.UpdatedAt = time.Now()

	return nil
}

func (pe *Person) ValidateCreate() error {
	if pe.ID != "" {
		return errs.NewBllValidateError("Person.ValidateCreate", "id must be empty", nil)
	}
	// ..

	return nil
}

func (pe *Person) ValidateChange() error {
	if pe.ID == "" {
		return errs.NewBllValidateError("Person.ValidateChange", "id cannot be empty", nil)
	}

	return nil
}

func (pe *Person) GetDeleted() bool {
	return pe.Deleted
}

func (pe *Person) SetDeleted(deleted bool) {
	pe.Deleted = deleted
}

func (pe *Person) IsDeleted() bool {
	return pe.Deleted
}

func (pe *Person) GetInternalTypeName() string {
	return utils.GetFullTypeName(pe)
}

func (pe *Person) GetTypeName() string {
	return "Person"
}

func (pe *Person) GetTypeDescription() string {
	return "Person model"
}

func (pe *Person) GetInstanceID() string {
	return pe.ID
}

func (pe *Person) GetInstanceName() string {
	// ..

	return ""
}

func (pe *Person) HashCode() uint32 {
	h := fnv.New32a()

	h.Write([]byte(pe.ID))
	h.Write([]byte(pe.ExternalID))
	h.Write([]byte(pe.LastName))
	h.Write([]byte(pe.FirstName))
	h.Write([]byte(pe.Patronymic))
	if pe.Deleted {
		h.Write([]byte{1})
	} else {
		h.Write([]byte{0})
	}
	h.Write([]byte(pe.CreatedAt.Format(time.RFC3339)))
	h.Write([]byte(pe.UpdatedAt.Format(time.RFC3339)))

	return h.Sum32()
}

func (pe *Person) ToAuditMap() map[string]*auditdomain.AuditField {
	res := make(map[string]*auditdomain.AuditField)

	res["id"] = auditdomain.NewAuditField(pe.ID, "УИЭ")
	res["external_id"] = auditdomain.NewAuditField(pe.ExternalID, "Внешний ID (интеграция)")
	res["last_name"] = auditdomain.NewAuditField(pe.LastName, "Фамилия")
	res["first_name"] = auditdomain.NewAuditField(pe.FirstName, "Имя")
	res["patronymic"] = auditdomain.NewAuditField(pe.Patronymic, "Отчество")
	res["deleted"] = auditdomain.NewAuditField(strconv.FormatBool(pe.Deleted), "Признак soft deleted")
	res["created_at"] = auditdomain.NewAuditField(pe.CreatedAt.Format(time.RFC3339), "Создано")
	res["updated_at"] = auditdomain.NewAuditField(pe.UpdatedAt.Format(time.RFC3339), "Изменено")

	return res
}

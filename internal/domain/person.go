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
	"github.com/davecgh/go-spew/spew"
)

type Person struct {
	ID         string    `db:"id"`
	ExternalID string    `db:"external_id"`
	LastName   string    `db:"last_name"`
	FirstName  string    `db:"first_name"`
	Patronymic string    `db:"patronymic"`
	Department string    `db:"department"`
	Position   string    `db:"position"`
	Status     string    `db:"status"`
	AvatarURL  string    `db:"avatar_url"`
	Active     bool      `db:"active"`
	Deleted    bool      `db:"deleted"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	Profiles   []*Profile
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
	department string,
	position string,
	status string,
	avatarURL string,
	active bool,
	deleted bool,
) *Person {
	return &Person{
		ID:         id,
		ExternalID: externalID,
		LastName:   lastName,
		FirstName:  firstName,
		Patronymic: patronymic,
		Department: department,
		Position:   position,
		Status:     status,
		AvatarURL:  avatarURL,
		Active:     active,
		Deleted:    deleted,
		Profiles:   make([]*Profile, 0),
	}
}

func NewEmptyPerson() *Person {
	return &Person{
		Active:   true,
		Deleted:  false,
		Profiles: make([]*Profile, 0),
	}
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
	return spew.Sprintf("LN: %s FN: %s P: %s", pe.LastName, pe.FirstName, pe.Patronymic)
}

func (pe *Person) HashCode() uint32 {
	h := fnv.New32a()

	h.Write([]byte(pe.ID))
	h.Write([]byte(pe.ExternalID))
	h.Write([]byte(pe.LastName))
	h.Write([]byte(pe.FirstName))
	h.Write([]byte(pe.Patronymic))
	h.Write([]byte(pe.Department))
	h.Write([]byte(pe.Position))
	h.Write([]byte(pe.Status))
	if pe.Active {
		h.Write([]byte{1})
	} else {
		h.Write([]byte{0})
	}
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
	return map[string]*auditdomain.AuditField{
		"id":          auditdomain.NewAuditField(pe.ID, "id"),
		"external_id": auditdomain.NewAuditField(pe.ExternalID, "Внешний ID (интеграция)"),
		"last_name":   auditdomain.NewAuditField(pe.LastName, "Фамилия"),
		"first_name":  auditdomain.NewAuditField(pe.FirstName, "Имя"),
		"patronymic":  auditdomain.NewAuditField(pe.Patronymic, "Отчество"),
		"department":  auditdomain.NewAuditField(pe.Department, "Департамент"),
		"position":    auditdomain.NewAuditField(pe.Position, "Позиция"),
		"status":      auditdomain.NewAuditField(pe.Status, "Статус"),
		"avatar_url":  auditdomain.NewAuditField(pe.AvatarURL, "Аватар"),
		"active":      auditdomain.NewAuditField(strconv.FormatBool(pe.Active), "Признк active"),
		"deleted":     auditdomain.NewAuditField(strconv.FormatBool(pe.Deleted), "Признак soft deleted"),
		"created_at":  auditdomain.NewAuditField(pe.CreatedAt.Format(time.RFC3339), "Создано"),
		"updated_at":  auditdomain.NewAuditField(pe.UpdatedAt.Format(time.RFC3339), "Изменено"),
	}
}

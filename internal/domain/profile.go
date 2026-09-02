package domain

import (
	"hash/fnv"
	"strconv"
	"time"

	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	auditdomain "github.com/ElfAstAhe/tiny-audit-service/pkg/domain"
	auditrepository "github.com/ElfAstAhe/tiny-audit-service/pkg/repository"
	"github.com/davecgh/go-spew/spew"
)

type Profile struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	PersonID  string    `db:"person_id"`
	TimeZone  string    `db:"time_zone"`
	Lang      string    `db:"lang"`
	Active    bool      `db:"active"`
	Deleted   bool      `db:"deleted"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

var _ libdomain.Entity[string] = (*Profile)(nil)
var _ libdomain.SoftDeleteEntity[bool] = (*Profile)(nil)
var _ auditdomain.Auditable = (*Profile)(nil)
var _ auditrepository.AuditableEntity[string] = (*Profile)(nil)

func (pr *Profile) GetID() string {
	return pr.ID
}

func (pr *Profile) SetID(id string) {
	pr.ID = id
}

func (pr *Profile) IsExists() bool {
	return pr.ID != ""
}

func (pr *Profile) BeforeCreate() error {
	if err := libdomain.AssignUUIDv7(pr); err != nil {
		return errs.NewBllError("Profile.BeforeCreate", "default before create failed", err)
	}

	if pr.CreatedAt.IsZero() {
		pr.CreatedAt = time.Now()
	}
	pr.UpdatedAt = time.Now()

	return nil
}

func (pr *Profile) BeforeChange() error {
	pr.UpdatedAt = time.Now()

	return nil
}

func (pr *Profile) ValidateCreate() error {
	if pr.ID != "" {
		return errs.NewBllValidateError("Profile.ValidateCreate", "id must be empty", nil)
	}

	if pr.UserID == "" {
		return errs.NewBllValidateError("Profile.ValidateCreate", "user_id must not be empty", nil)
	}

	if err := utils.ValidateLanguage(pr.Lang); err != nil {
		return errs.NewBllValidateError("Profile.ValidateCreate", "lang validation failed", err)
	}

	if err := utils.ValidateTimeZone(pr.TimeZone); err != nil {
		return errs.NewBllValidateError("Profile.ValidateCreate", "time_zone validation failed", err)
	}

	return nil
}

func (pr *Profile) ValidateChange() error {
	if pr.ID == "" {
		return errs.NewBllValidateError("Profile.ValidateChange", "id cannot be empty", nil)
	}

	if pr.UserID == "" {
		return errs.NewBllValidateError("Profile.ValidateChange", "user_id must not be empty", nil)
	}

	if err := utils.ValidateLanguage(pr.Lang); err != nil {
		return errs.NewBllValidateError("Profile.ValidateChange", "lang validation failed", err)
	}

	if err := utils.ValidateTimeZone(pr.TimeZone); err != nil {
		return errs.NewBllValidateError("Profile.ValidateChange", "time_zone validation failed", err)
	}

	return nil
}

func (pr *Profile) GetDeleted() bool {
	return pr.Deleted
}

func (pr *Profile) SetDeleted(deleted bool) {
	pr.Deleted = deleted
}

func (pr *Profile) IsDeleted() bool {
	return pr.Deleted
}

func (pr *Profile) GetInternalTypeName() string {
	return utils.GetFullTypeName(pr)
}

func (pr *Profile) GetTypeName() string {
	return "Profile"
}

func (pr *Profile) GetTypeDescription() string {
	return "Profile model"
}

func (pr *Profile) GetInstanceID() string {
	return pr.ID
}

func (pr *Profile) GetInstanceName() string {
	return spew.Sprintf("Profile of user_id [%s]", pr.UserID)
}

//goland:noinspection DuplicatedCode
func (pr *Profile) HashCode() uint32 {
	h := fnv.New32a()

	h.Write([]byte(pr.ID))
	h.Write([]byte(pr.UserID))
	h.Write([]byte(pr.PersonID))
	h.Write([]byte(pr.TimeZone))
	h.Write([]byte(pr.Lang))
	if pr.Active {
		h.Write([]byte{1})
	} else {
		h.Write([]byte{0})
	}
	if pr.Deleted {
		h.Write([]byte{1})
	} else {
		h.Write([]byte{0})
	}
	h.Write([]byte(pr.CreatedAt.Format(time.RFC3339)))
	h.Write([]byte(pr.UpdatedAt.Format(time.RFC3339)))

	return h.Sum32()
}

func (pr *Profile) ToAuditMap() map[string]*auditdomain.AuditField {
	return map[string]*auditdomain.AuditField{
		"id":         auditdomain.NewAuditField(pr.ID, "id"),
		"user_id":    auditdomain.NewAuditField(pr.UserID, "user_id"),
		"person_id":  auditdomain.NewAuditField(pr.PersonID, "person_id"),
		"time_zone":  auditdomain.NewAuditField(pr.TimeZone, "time_zone"),
		"lang":       auditdomain.NewAuditField(pr.Lang, "lang"),
		"active":     auditdomain.NewAuditField(strconv.FormatBool(pr.Active), "Признк active"),
		"deleted":    auditdomain.NewAuditField(strconv.FormatBool(pr.Deleted), "Признак soft deleted"),
		"created_at": auditdomain.NewAuditField(pr.CreatedAt.Format(time.RFC3339), "Создано"),
		"updated_at": auditdomain.NewAuditField(pr.UpdatedAt.Format(time.RFC3339), "Изменено"),
	}
}

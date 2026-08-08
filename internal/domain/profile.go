package domain

import (
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/domain"
	auditdomain "github.com/ElfAstAhe/tiny-audit-service/pkg/domain"
	auditrepository "github.com/ElfAstAhe/tiny-audit-service/pkg/repository"
)

type Profile struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	Deleted   bool      `db:"deleted"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

var _ domain.Entity[string] = (*Profile)(nil)
var _ domain.SoftDeleteEntity[bool] = (*Profile)(nil)
var _ auditdomain.Auditable = (*Profile)(nil)
var _ auditrepository.AuditableEntity[string] = (*Profile)(nil)

func (pr *Profile) GetID() string {
	//TODO implement me
	panic("implement me")
}

func (pr *Profile) SetID(id string) {
	//TODO implement me
	panic("implement me")
}

func (pr *Profile) IsExists() bool {
	//TODO implement me
	panic("implement me")
}

func (pr *Profile) BeforeCreate() error {
	//TODO implement me
	panic("implement me")
}

func (pr *Profile) BeforeChange() error {
	//TODO implement me
	panic("implement me")
}

func (pr *Profile) ValidateCreate() error {
	//TODO implement me
	panic("implement me")
}

func (pr *Profile) ValidateChange() error {
	//TODO implement me
	panic("implement me")
}

func (pr *Profile) GetDeleted() bool {
	//TODO implement me
	panic("implement me")
}

func (pr *Profile) SetDeleted(deleted bool) {
	//TODO implement me
	panic("implement me")
}

func (pr *Profile) IsDeleted() bool {
	//TODO implement me
	panic("implement me")
}

func (pr *Profile) GetInternalTypeName() string {
	//TODO implement me
	panic("implement me")
}

func (pr *Profile) GetTypeName() string {
	//TODO implement me
	panic("implement me")
}

func (pr *Profile) GetTypeDescription() string {
	//TODO implement me
	panic("implement me")
}

func (pr *Profile) GetInstanceID() string {
	//TODO implement me
	panic("implement me")
}

func (pr *Profile) GetInstanceName() string {
	//TODO implement me
	panic("implement me")
}

func (pr *Profile) HashCode() uint32 {
	//TODO implement me
	panic("implement me")
}

func (pr *Profile) ToAuditMap() map[string]*auditdomain.AuditField {
	//TODO implement me
	panic("implement me")
}

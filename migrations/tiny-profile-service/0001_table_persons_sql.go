package tiny_profile_service

const (
	sqlCreateTablePersons string = `
create table if not exists persons (
    id varchar(50) not null,
    external_id varchar(50) null,
    last_name varchar(100) null,
    first_name varchar(100) null,
    patronymic varchar(100) null,
    birthday timestamptz null,
    department varchar(100) null,
    position varchar(100) null,
    status varchar(100) null,
    avatar_url text null,
    active bool not null default false,
    deleted bool not null default false,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    constraint persons_pk primary key (id)
)
`
	sqlDropTablePersons string = `
drop table if exists persons cascade
`
	sqlCreateIndexPersonsByID string = `
create index if not exists idx_persons_id on persons(deleted asc, id asc)
`
	sqlDropIndexPersonsByID string = `
drop index if exists idx_persons_id
`
	sqlCreateIndexPersonsByExternalID string = `
create index if not exists idx_persons_external_id on persons(deleted asc, external_id asc)
`
	sqlDropIndexPersonsByExternalID string = `
drop index if exists idx_persons_external_id
`
)

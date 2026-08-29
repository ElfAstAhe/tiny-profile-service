package tiny_profile_service

const (
	sqlCreateTableProfiles string = `
create table if not exists profiles (
    id varchar(50) not null,
    user_id varchar(100) not null,
    person_id varchar(50) null,
    time_zone varchar(50) null,
    lang varchar(10) null default 'ru',
    active bool not null default false,
    deleted bool not null default false,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    constraint profiles_pk primary key (id),
    constraint profiles_person_fk foreign key (person_id) references persons(id),
    constraint profiles_uk unique (user_id)
)
`

	sqlDropTableProfiles string = `
drop table if exists profiles
`

	sqlCreateIndexProfilesByID string = `
create index if not exists idx_profiles_id on profiles(deleted desc, id asc)
`

	sqlDropIndexProfilesByID string = `
drop index if exists idx_profiles_id
`

	sqlCreateIndexProfilesByUserID string = `
create index if not exists idx_profiles_user_id on profiles(deleted desc, user_id asc)
`

	sqlDropIndexProfilesByUserID string = `
drop index if exists idx_profiles_user_id
`
	sqlCreateIndexProfilesByPersonID string = `
create index if not exists idx_profiles_person on profiles(deleted desc, person_id asc)
`
	sqlDropIndexProfilesByPersonID string = `
drop index if exists idx_profiles_person
`
)

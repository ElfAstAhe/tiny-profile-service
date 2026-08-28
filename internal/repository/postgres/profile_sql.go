package postgres

const (
	sqlProfileList string = `
select
    id,
    user_id,
    person_id,
    time_zone,
    lang,
    active,
    deleted,
    created_at,
    updated_at
from
    profiles
order by
    user_id asc,
    id asc
offset $2
limit $1
`
	sqlProfileFind string = `
select
    id,
    user_id,
    person_id,
    time_zone,
    lang,
    active,
    deleted,
    created_at,
    updated_at
from
    profiles
where
    id = $1
and person_id = $2
`
	sqlProfileCreate string = `
insert into profiles (
    id,
    user_id,
    person_id,
    time_zone,
    lang,
    active,
    deleted,
    created_at,
    updated_at
)
values($1, $2, $3, $4, $5, $6, false, $7, $8)
returning
    id,
    user_id,
    person_id,
    time_zone,
    lang,
    active,
    deleted,
    created_at,
    updated_at
`
	sqlProfileChange string = `
update
    profiles
set
    time_zone = $2,
    lang = $3,
    active = $4,
    updated = $5
where
    id = $1
returning
    id,
    user_id,
    person_id,
    time_zone,
    lang,
    active,
    deleted,
    created_at,
    updated_at
`
	sqlProfileDelete string = `
update
    profiles
set
    deleted = true
where
    id = $1
`
	sqlProfileFindByUserID string = `
select
    id,
    user_id,
    person_id,
    time_zone,
    lang,
    active,
    deleted,
    created_at,
    updated_at
from
    profiles
where
    user_id = $1
`
	sqlProfileFindByPersonID string = `
select
    id,
    user_id,
    person_id,
    time_zone,
    lang,
    active,
    deleted,
    created_at,
    updated_at
from
    profiles
where
    deleted = false
and person_id = $1
`
)

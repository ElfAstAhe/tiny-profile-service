package postgres

const (
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
where
    deleted = false
and person_id = $1
order by
    user_id asc,
    id asc
offset $2
limit $2
`
	sqlProfileListAll string = `
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
order by
    user_id asc,
    id asc
`
	sqlProfileListAllByOwners string = `
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
    person_id = any($1)
order by
    person_id asc,
    user_id asc,
    id asc
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
values($1, $2, $3, $4, $5, $6, $7, $8, $9)
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
    deleted = $5,
    updated = $6
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
	sqlProfileDeleteAll string = `
update
    profiles
set
    deleted = true
where
    person_id = $1
`
	sqlProfileDelete string = `
update
    profiles
set
    deleted = true
where
    id = $1
`
	sqlProfileFindByID string = `
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
)

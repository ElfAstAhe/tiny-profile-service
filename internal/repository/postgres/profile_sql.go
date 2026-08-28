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

	sqlProfileCreate string = ``

	sqlProfileChange string = ``

	sqlProfileDelete string = ``
)

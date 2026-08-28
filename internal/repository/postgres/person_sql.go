package postgres

const (
	sqlPersonList string = `
select
    id,
    external_id,
    last_name,
    first_name,
    patronymic,
    department,
    position,
    status,
    avatar_url,
    active,
    deleted,
    created_at,
    updated_at
from
    persons
order by
    id asc
offset $2
limit $1
`

	sqlPersonFind string = `
select
    id,
    external_id,
    last_name,
    first_name,
    patronymic,
    department,
    position,
    status,
    avatar_url,
    active,
    deleted,
    created_at,
    updated_at
from
    persons
where
    id = $1
`

	sqlPersonCreate string = `
insert into Persons(
    id,
    external_id,
    last_name,
    first_name,
    patronymic,
    department,
    position,
    status,
    avatar_url,
    active,
    deleted,
    created_at,
    updated_at
)
values(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
)
returning
    id,
    external_id,
    last_name,
    first_name,
    patronymic,
    department,
    position,
    status,
    avatar_url,
    active,
    deleted,
    created_at,
    updated_at
`

	sqlPersonChange string = `
update
    persons
set
    external_id = $2,
    last_name = $3,
    first_name = $4,
    patronymic = $5,
    department = $6,
    position = $7,
    status = $8,
    avatar_url = $9,
    active = $10,
    deleted = $11,
    updated_at = $12
where
    id = $1
returning 
    id,
    external_id,
    last_name,
    first_name,
    patronymic,
    department,
    position,
    status,
    avatar_url,
    active,
    deleted,
    created_at,
    updated_at
`

	sqlPersonDelete string = `
update
    persons
set
    deleted = true
where
    id = $1
`
)

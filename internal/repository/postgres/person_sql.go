package postgres

const (
	sqlPersonList string = `
select
    id,
    external_id,
    last_name,
    first_name,
    patronymic,
    birthday,
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
    deleted = false
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
    birthday,
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
    deleted = false
and id = $1
`
	sqlPersonCreate string = `
insert into Persons(
    id,
    external_id,
    last_name,
    first_name,
    patronymic,
	birthday,
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
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13,$14
)
returning
    id,
    external_id,
    last_name,
    first_name,
    patronymic,
    birthday,
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
    birthday = $6,
    department = $7,
    position = $8,
    status = $9,
    avatar_url = $10,
    active = $11,
    updated_at = $12
where
    id = $1
returning 
    id,
    external_id,
    last_name,
    first_name,
    patronymic,
    birthday,
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
	sqlPersonFindByExternalID = `
select
    id,
    external_id,
    last_name,
    first_name,
    patronymic,
    birthday,
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
    deleted = false
and external_id = $1
`
)

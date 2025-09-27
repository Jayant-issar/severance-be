-- name: CreateUser :one
insert into users (
    id,
    username,
    email,
    password_hash
) values (
    $1,$2,$3,$4
) RETURNING *;

-- name: GetUser :one
select * from users
where username = $1 LIMIT 1;

-- name: GetUserByID :one
select * from users where id = $1;

-- name: GetUserByEmail :one
select * from users where email = $1;

-- name: UpdateUser :one
update users set username = $2, email = $3, password_hash = $4 where id = $1 returning *;

-- name: DeleteUser :exec
delete from users where id = $1;

-- name: ListUsers :many
select * from users order by created_at desc;
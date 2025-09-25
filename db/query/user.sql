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
-- name: CreateQuestion :one
insert into questions (id, title, description, difficulty, tags) values ($1, $2, $3, $4, $5) returning *;

-- name: GetQuestion :one
select * from questions where id = $1;

-- name: UpdateQuestion :one
update questions set title = $2, description = $3, difficulty = $4, tags = $5 where id = $1 returning *;

-- name: DeleteQuestion :exec
delete from questions where id = $1;

-- name: ListQuestions :many
select * from questions order by created_at desc limit $1 offset $2;
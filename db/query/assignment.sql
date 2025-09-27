-- name: CreateAssignment :one
insert into assignments (id, title, description, difficulty, tags) values ($1, $2, $3, $4, $5) returning *;

-- name: GetAssignment :one
select * from assignments where id = $1;

-- name: UpdateAssignment :one
update assignments set title = $2, description = $3, difficulty = $4, tags = $5 where id = $1 returning *;

-- name: DeleteAssignment :exec
delete from assignments where id = $1;

-- name: ListAssignments :many
select * from assignments order by created_at desc;
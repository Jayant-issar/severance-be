-- name: CreateSubmission :one
insert into submissions (id, user_id, assignment_id, code, language, status, runtime_ms, memory_kb) values ($1, $2, $3, $4, $5, $6, $7, $8) returning *;

-- name: GetSubmission :one
select * from submissions where id = $1;

-- name: GetSubmissionsByUser :many
select * from submissions where user_id = $1 order by created_at desc;

-- name: GetSubmissionsByAssignment :many
select * from submissions where assignment_id = $1 order by created_at desc;

-- name: UpdateSubmission :one
update submissions set status = $2, runtime_ms = $3, memory_kb = $4 where id = $1 returning *;

-- name: DeleteSubmission :exec
delete from submissions where id = $1;
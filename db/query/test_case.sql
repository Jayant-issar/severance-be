-- name: CreateTestCase :one
insert into test_cases (id, assignment_id, input, expected_output, is_hidden) values ($1, $2, $3, $4, $5) returning *;

-- name: GetTestCase :one
select * from test_cases where id = $1;

-- name: GetTestCasesByAssignment :many
select * from test_cases where assignment_id = $1 order by created_at;

-- name: UpdateTestCase :one
update test_cases set assignment_id = $2, input = $3, expected_output = $4, is_hidden = $5 where id = $1 returning *;

-- name: DeleteTestCase :exec
delete from test_cases where id = $1;
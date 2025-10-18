-- name: CreateAIReview :one
insert into ai_reviews (id, submission_id, feedback, score, review_agent) values ($1, $2, $3, $4, $5) returning *;

-- name: GetAIReview :one
select * from ai_reviews where id = $1;

-- name: GetAIReviewsBySubmission :many
select * from ai_reviews where submission_id = $1 order by created_at desc;

-- name: UpdateAIReview :one
update ai_reviews set feedback = $2, score = $3, review_agent = $4 where id = $1 returning *;

-- name: DeleteAIReview :exec
delete from ai_reviews where id = $1;
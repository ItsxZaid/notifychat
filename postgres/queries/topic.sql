-- name: CreateTopic :one
INSERT INTO topics (name, description)
VALUES ($1, $2)
RETURNING *;

-- name: GetTopic :one
SELECT * FROM topics
WHERE id = $1;

-- name: GetAllTopics :many
SELECT * FROM topics;

-- name: DeleteTopic :exec
DELETE FROM topics
WHERE id = $1;

-- name: UpdateTopic :one
UPDATE topics
SET
  name = $1,
  description = $2
WHERE
  id = $3
RETURNING *;

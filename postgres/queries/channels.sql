-- name: CreateChannel :one
INSERT INTO channels (topic_id, type, config, template)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetChannelsByTopicID :many
SELECT * FROM channels
WHERE topic_id = $1;

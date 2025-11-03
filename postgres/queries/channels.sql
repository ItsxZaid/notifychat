-- name: CreateChannel :one
INSERT INTO channels (campaign_id, type, config, template)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetChannelsByCampaignID :many
SELECT * FROM channels
WHERE campaign_id = $1;

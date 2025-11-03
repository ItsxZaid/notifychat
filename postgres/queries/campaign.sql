-- name: CreateCampaign :one
INSERT INTO campaigns (name)
VALUES ($1)
RETURNING *;

-- name: GetCampaign :one
SELECT * FROM campaigns
WHERE id = $1;

-- name: CreateChannel :one
INSERT INTO channels (campaign_id, type, config, template)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetChannelsByCampaignID :many
SELECT * FROM channels
WHERE campaign_id = $1;

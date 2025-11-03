-- name: CreateCampaign :one
INSERT INTO campaigns (name)
VALUES ($1)
RETURNING *;

-- name: GetCampaign :one
SELECT * FROM campaigns
WHERE id = $1;

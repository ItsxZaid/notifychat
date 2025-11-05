-- name: CreateCampaign :one
INSERT INTO campaigns (name)
VALUES ($1)
RETURNING *;

-- name: GetCampaign :one
SELECT * FROM campaigns
WHERE id = $1;

-- name: GetAllCampaigns :many
SELECT * FROM campaigns;

-- name: DeleteCampaign :exec
DELETE FROM campaigns
WHERE id = $1;

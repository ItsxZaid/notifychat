package domain

import (
	"context"
	"time"
)

type Campaign struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type CampaignRepository interface {
	GetCampaign(ctx context.Context, id string) (*Campaign, error)
	CreateCampaign(ctx context.Context, name string) (*Campaign, error)
}

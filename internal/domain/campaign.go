package domain

import (
	"context"
	"time"
)

type Campaign struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

type CampaignRepository interface {
	GetCompaign(ctx context.Context, id string) (*Campaign, error)
}

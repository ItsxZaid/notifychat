package store

import (
	"context"
	"solution-for-x/notifychat/internal/domain"
	"solution-for-x/notifychat/internal/store/db.go"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CampaignStore struct {
	db *db.Queries
}

var _ domain.CampaignRepository = (*CampaignStore)(nil)

func NewCampaignStore(pool *pgxpool.Pool) *CampaignStore {
	return &CampaignStore{
		db: db.New(pool),
	}
}

func (cs *CampaignStore) GetCompaign(ctx context.Context, id string) (*domain.Campaign, error) {
	// Implementation here
	return nil, nil
}

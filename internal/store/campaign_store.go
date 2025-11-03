package store

import (
	"context"
	"solution-for-x/notifychat/internal/domain"
	"solution-for-x/notifychat/internal/store/sqlc_generated"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CampaignStore struct {
	db *sqlc_generated.Queries
}

var _ domain.CampaignRepository = (*CampaignStore)(nil)

func NewCampaignStore(pool *pgxpool.Pool) *CampaignStore {
	return &CampaignStore{
		db: sqlc_generated.New(pool),
	}
}

func (cs *CampaignStore) GetCompaign(ctx context.Context, id string) (*domain.Campaign, error) {
	// Implementation here
	return nil, nil
}

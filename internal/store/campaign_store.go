package store

import (
	"context"
	"itsxzaid/notifychat/internal/domain"
	"itsxzaid/notifychat/internal/store/sqlc_generated"

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

func (cs *CampaignStore) GetCampaign(ctx context.Context, id string) (*domain.Campaign, error) {
	// Implementation here
	return nil, nil
}

func (cs *CampaignStore) CreateCampaign(ctx context.Context, name string) (*domain.Campaign, error) {
	if campaign, err := cs.db.CreateCampaign(ctx, name); err != nil {
		return nil, err
	} else {
		return &domain.Campaign{
			ID:        campaign.ID.String(),
			Name:      campaign.Name,
			CreatedAt: campaign.CreatedAt.Time,
		}, nil
	}
}

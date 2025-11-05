package store

import (
	"context"
	"fmt"
	"itsxzaid/notifychat/internal/domain"
	"itsxzaid/notifychat/internal/store/sqlc_generated"

	"github.com/jackc/pgx/v5/pgtype"
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

func (cs *CampaignStore) GetAllCampaigns(ctx context.Context) ([]domain.Campaign, error) {
	sqlcCampaigns, err := cs.db.GetAllCampaigns(ctx)
	if err != nil {
		return nil, err
	}

	if sqlcCampaigns == nil {
		return []domain.Campaign{}, nil
	}

	domainCampaigns := make([]domain.Campaign, 0, len(sqlcCampaigns))

	for _, sqlcCampaign := range sqlcCampaigns {
		domainCampaigns = append(domainCampaigns, domain.Campaign{
			ID:        sqlcCampaign.ID.String(),
			Name:      sqlcCampaign.Name,
			CreatedAt: sqlcCampaign.CreatedAt.Time,
		})
	}

	return domainCampaigns, nil
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

func (cs *CampaignStore) DeleteCampaign(ctx context.Context, id string) error {
	var pgUUID pgtype.UUID

	if err := pgUUID.Scan(id); err != nil {
		return domain.ErrInvalidInput
	}

	if err := cs.db.DeleteCampaign(ctx, pgUUID); err != nil {
		fmt.Printf("%v", err)
		return err
	}

	return nil
}

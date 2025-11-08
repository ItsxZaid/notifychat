package store

import (
	"context"
	"fmt"
	"itsxzaid/notifychat/internal/domain"
	"itsxzaid/notifychat/internal/store/sqlc_generated"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ChannelStore struct {
	db *sqlc_generated.Queries
}

var _ domain.ChannelRepository = (*ChannelStore)(nil)

func NewChannelStore(pool *pgxpool.Pool) *ChannelStore {
	return &ChannelStore{
		db: sqlc_generated.New(pool),
	}
}

func (cs *ChannelStore) CreateChannel(ctx context.Context, payload sqlc_generated.CreateChannelParams) (*domain.Channel, error) {
	if channel, err := cs.db.CreateChannel(ctx, payload); err != nil {
		return nil, err
	} else {
		var channelType = domain.ChannelType(channel.Type)

		cfg, err := domain.UnmarshalChannelConfig(channel.Config, channelType)
		if err != nil {
			return nil, err
		}

		return &domain.Channel{
			ID:        channel.ID.String(),
			TopicID:   channel.TopicID.String(),
			Type:      channelType,
			Config:    cfg,
			Template:  channel.Template,
			CreatedAt: channel.CreatedAt.Time,
		}, nil
	}
}

func (cs *ChannelStore) GetChannelsByTopicID(ctx context.Context, topic_id string) ([]domain.Channel, error) {
	var pgUUID pgtype.UUID

	if err := pgUUID.Scan(topic_id); err != nil {
		return nil, domain.ErrInvalidInput
	}

	sqlcChannels, err := cs.db.GetChannelsByTopicID(ctx, pgUUID)
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}

	domainChannels := make([]domain.Channel, 0, len(sqlcChannels))

	for _, sqlcChannel := range sqlcChannels {
		var channelType = domain.ChannelType(sqlcChannel.Type)

		cfg, err := domain.UnmarshalChannelConfig(sqlcChannel.Config, channelType)
		if err != nil {
			return nil, err
		}

		domainChannels = append(domainChannels, domain.Channel{
			ID:        sqlcChannel.ID.String(),
			TopicID:   sqlcChannel.TopicID.String(),
			Type:      channelType,
			Config:    cfg,
			Template:  sqlcChannel.Template,
			CreatedAt: sqlcChannel.CreatedAt.Time,
		})
	}

	return domainChannels, nil
}

package store

import (
	"context"
	"fmt"
	"itsxzaid/notifychat/internal/domain"
	"itsxzaid/notifychat/internal/store/sqlc_generated"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TopicStore struct {
	db *sqlc_generated.Queries
}

// Compile-time interface satisfaction check. Always satisfy baby
var _ domain.TopicRepository = (*TopicStore)(nil)

func NewTopicStore(pool *pgxpool.Pool) *TopicStore {
	return &TopicStore{
		db: sqlc_generated.New(pool),
	}
}

func (cs *TopicStore) GetAllTopics(ctx context.Context) ([]domain.Topic, error) {
	sqlcTopics, err := cs.db.GetAllTopics(ctx)
	if err != nil {
		return nil, err
	}

	if sqlcTopics == nil {
		return []domain.Topic{}, nil
	}

	domainTopics := make([]domain.Topic, 0, len(sqlcTopics))

	for _, sqlcTopic := range sqlcTopics {
		domainTopics = append(domainTopics, domain.Topic{
			ID:          sqlcTopic.ID.String(),
			Name:        sqlcTopic.Name,
			Description: sqlcTopic.Description.String,
			CreatedAt:   sqlcTopic.CreatedAt.Time,
		})
	}

	return domainTopics, nil
}

func (cs *TopicStore) CreateTopic(ctx context.Context, payload sqlc_generated.CreateTopicParams) (*domain.Topic, error) {
	if topic, err := cs.db.CreateTopic(ctx, payload); err != nil {
		return nil, err
	} else {
		return &domain.Topic{
			ID:          topic.ID.String(),
			Name:        topic.Name,
			Description: topic.Description.String,
			CreatedAt:   topic.CreatedAt.Time,
		}, nil
	}
}

func (cs *TopicStore) UpdateTopic(ctx context.Context, payload sqlc_generated.UpdateTopicParams) (*domain.Topic, error) {
	if topic, err := cs.db.UpdateTopic(ctx, payload); err != nil {
		return nil, err
	} else {
		return &domain.Topic{
			ID:          topic.ID.String(),
			Name:        topic.Name,
			Description: topic.Description.String,
			CreatedAt:   topic.CreatedAt.Time,
		}, nil
	}
}

func (cs *TopicStore) DeleteTopic(ctx context.Context, id string) error {
	var pgUUID pgtype.UUID

	if err := pgUUID.Scan(id); err != nil {
		return domain.ErrInvalidInput
	}

	if err := cs.db.DeleteTopic(ctx, pgUUID); err != nil {
		fmt.Printf("%v", err)
		return err
	}

	return nil
}

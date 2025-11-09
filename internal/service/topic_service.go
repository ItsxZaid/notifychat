package service

import (
	"context"
	"errors"
	"itsxzaid/notifychat/internal/domain"
	"itsxzaid/notifychat/internal/store/sqlc_generated"

	"github.com/jackc/pgx/v5/pgtype"
)

type TopicService struct {
	repo domain.TopicRepository
}

type CreateTopicParams struct {
	Name string

	// Pointer type cause Description is optional and can be nil.
	Description *string
}

type UpdateTopicParams struct {
	TopicID     string
	Name        *string
	Description *string
}

func NewTopicService(repo domain.TopicRepository) *TopicService {
	return &TopicService{
		repo: repo,
	}
}

func (ts *TopicService) GetAllTopics(ctx context.Context) ([]domain.Topic, error) {
	return ts.repo.GetAllTopics(ctx)
}

func (ts *TopicService) CreateTopic(ctx context.Context, params CreateTopicParams) (*domain.Topic, error) {
	var pgDescription pgtype.Text

	if params.Description != nil {
		pgDescription = pgtype.Text{
			String: *params.Description,
			Valid:  true,
		}
	}

	topic, err := ts.repo.CreateTopic(ctx, sqlc_generated.CreateTopicParams{
		Name:        params.Name,
		Description: pgDescription,
	})
	if err != nil {
		return nil, err
	}

	return topic, nil
}

func (ts *TopicService) UpdateTopic(ctx context.Context, params UpdateTopicParams) (*domain.Topic, error) {
	var pgDescription pgtype.Text

	if params.Description != nil {
		pgDescription = pgtype.Text{
			String: *params.Description,
			Valid:  true,
		}
	}

	var pgUUID pgtype.UUID

	if err := pgUUID.Scan(params.TopicID); err != nil {
		return nil, errors.New("The topic ID is not a valid UUID")
	}

	updatedTopic, err := ts.repo.UpdateTopic(ctx, sqlc_generated.UpdateTopicParams{
		ID:          pgUUID,
		Name:        *params.Name,
		Description: pgDescription,
	})
	if err != nil {
		return nil, err
	}

	return updatedTopic, nil
}

func (ts *TopicService) DeleteTopic(ctx context.Context, topicID string) error {
	return ts.repo.DeleteTopic(ctx, topicID)
}

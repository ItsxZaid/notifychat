package domain

import (
	"context"
	"itsxzaid/notifychat/internal/store/sqlc_generated"
	"time"
)

type Topic struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type TopicRepository interface {
	GetAllTopics(ctx context.Context) ([]Topic, error)
	CreateTopic(ctx context.Context, payload sqlc_generated.CreateTopicParams) (*Topic, error)
	UpdateTopic(ctx context.Context, payload sqlc_generated.UpdateTopicParams) (*Topic, error)
	DeleteTopic(ctx context.Context, id string) error
}

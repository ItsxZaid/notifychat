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

// NOTE: An interface is a "plug".
// When i pass a TopicRepository around, i am not copying
// the whole TopicStore.
// The interface is just two small pointers:
// Pointer-1: Point to the "actual value" (my *TopicStore)
// Pointer-2: Point to a "method table" that lists the real functions.
// this is why it fast and powerful.

type TopicRepository interface {
	GetAllTopics(ctx context.Context) ([]Topic, error)
	CreateTopic(ctx context.Context, payload sqlc_generated.CreateTopicParams) (*Topic, error)
	UpdateTopic(ctx context.Context, payload sqlc_generated.UpdateTopicParams) (*Topic, error)
	DeleteTopic(ctx context.Context, id string) error
}

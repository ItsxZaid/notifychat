package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"itsxzaid/notifychat/internal/store/sqlc_generated"
	"time"
)

// ChannelConfig is a sealed marker interface for valid config types.
type ChannelConfig interface {
	isChannelConfig()
}

type TelegramConfig struct {
	BotToken string `json:"bot_token"`
}

type WhatsappConfig struct {
	SessionID string `json:"session_id"`
}

func (TelegramConfig) isChannelConfig() {}
func (WhatsappConfig) isChannelConfig() {}

type ChannelType string

const (
	Telegram ChannelType = "telegram"
	Whatsapp ChannelType = "whatsapp"
)

type Channel struct {
	ID      string      `json:"id"`
	TopicID string      `json:"topic_id"`
	Type    ChannelType `json:"type"`
	// Config must be one of the types that implements ChannelConfig
	// (TelegramConfig, WhatsappConfig).
	Config    ChannelConfig `json:"config"`
	Template  string        `json:"template"`
	CreatedAt time.Time     `json:"created_at"`
}

type ChannelRepository interface {
	GetChannelsByTopicID(ctx context.Context, topic_id string) ([]Channel, error)
	CreateChannel(ctx context.Context, payload sqlc_generated.CreateChannelParams) (*Channel, error)
}

// UnmarshalChannelConfig converts raw JSON bytes into the proper ChannelConfig type.
// You need to pass the type (e.g., "telegram" or "whatsapp") so it knows which struct to use.
func UnmarshalChannelConfig(data []byte, typ ChannelType) (ChannelConfig, error) {
	switch typ {
	case Telegram:
		var cfg TelegramConfig
		if err := json.Unmarshal(data, &cfg); err != nil {
			return nil, fmt.Errorf("failed to unmarshal telegram config: %w", err)
		}
		return cfg, nil
	case Whatsapp:
		var cfg WhatsappConfig
		if err := json.Unmarshal(data, &cfg); err != nil {
			return nil, fmt.Errorf("failed to unmarshal whatsapp config: %w", err)
		}
		return cfg, nil
	default:
		return nil, fmt.Errorf("unknown channel type: %s", typ)
	}
}

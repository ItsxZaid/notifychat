package domain

import (
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

type Channel struct {
	ID      string `json:"id"`
	TopicID string `json:"topic_id"`
	// Config must be one of the types that implements ChannelConfig
	// (TelegramConfig, WhatsappConfig).
	Config    ChannelConfig `json:"config"`
	Template  string        `json:"template"`
	CreatedAt time.Time     `json:"created_at"`
}

type ChannelRepository interface{}

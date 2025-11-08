package service

import "itsxzaid/notifychat/internal/domain"

type ChannelService struct {
	repo domain.ChannelRepository
}

func NewChannelService(repo domain.ChannelRepository) *ChannelService {
	return &ChannelService{
		repo: repo,
	}
}

package service

import "itsxzaid/notifychat/internal/domain"

type TopicService struct {
	repo domain.TopicRepository
}

func NewTopicService(repo domain.TopicRepository) *TopicService {
	return &TopicService{
		repo: repo,
	}
}

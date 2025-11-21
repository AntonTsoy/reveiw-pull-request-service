package service

type SubscriptionRepository interface {
}

type SubsService struct {
	repo SubscriptionRepository
}

func NewSubsService(repo SubscriptionRepository) *SubsService {
	return &SubsService{repo: repo}
}

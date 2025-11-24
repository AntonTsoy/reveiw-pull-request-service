package service

import (
	"github.com/AntonTsoy/review-pull-request-service/internal/database"
	"github.com/AntonTsoy/review-pull-request-service/internal/repository"
)

type Service struct {
	TeamService *TeamService
}

func NewService(db *database.Database, repository *repository.Repository) *Service {
	return &Service{
		TeamService: newTeamService(db, repository.TeamRepository, repository.UserRepository),
	}
}

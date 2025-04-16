package service

import (
	"github.com/uptrace/bun"
	"test_task/internal/manager"
	"test_task/internal/repository"
)

type Service struct {
	userManager user.ManagerSDK
}

type ServiceSDK interface {
	GetUserManager() user.ManagerSDK
}

func NewService(db bun.IDB) *Service {
	userManager := user.NewManager(repository.NewUserRepo(db))

	return &Service{
		userManager: userManager,
	}
}

func (s *Service) GetUserManager() user.ManagerSDK {
	return s.userManager
}

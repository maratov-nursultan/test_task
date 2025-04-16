package service

import (
	"github.com/maratov-nursultan/profile/internal/manager"
	"github.com/maratov-nursultan/profile/internal/repository"
	"github.com/uptrace/bun"
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

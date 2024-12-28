package v1

import (
	"github.com/webbsalad/storya-content-backend/internal/config"
	"github.com/webbsalad/storya-content-backend/internal/repository/content"
	user_content_service "github.com/webbsalad/storya-content-backend/internal/service/user_content"
)

type Service struct {
	contentRepository content.Repository
	config            config.Config
}

func NewService(contentRepository content.Repository, config config.Config) user_content_service.Service {
	return &Service{
		contentRepository: contentRepository,
		config:            config,
	}
}

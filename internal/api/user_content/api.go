package user_content

import (
	desc "github.com/webbsalad/storya-content-backend/internal/pb/github.com/webbsalad/storya-content-backend/content"
	service "github.com/webbsalad/storya-content-backend/internal/service/user_content"
)

type Implementation struct {
	desc.UnimplementedUserContentServiceServer

	UserContentService service.Service
}

func NewUserContentImplementation(userContentService service.Service) desc.UserContentServiceServer {
	return &Implementation{
		UserContentService: userContentService,
	}
}

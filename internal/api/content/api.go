package content

import (
	desc "github.com/webbsalad/storya-content-backend/internal/pb/github.com/webbsalad/storya-content-backend/content"
	service "github.com/webbsalad/storya-content-backend/internal/service/content"
)

type Implementation struct {
	desc.UnimplementedContentServiceServer

	ContentService service.Service
}

func NewImplementation(contentService service.Service) desc.ContentServiceServer {
	return &Implementation{
		ContentService: contentService,
	}
}

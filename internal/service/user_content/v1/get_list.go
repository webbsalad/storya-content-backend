package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/storya-content-backend/internal/model"
)

func (s *Service) GetList(ctx context.Context, userID model.UserID, contentType model.ContentType) ([]model.UserItem, error) {
	items, err := s.contentRepository.GetList(ctx, userID, contentType)
	if err != nil {
		return nil, fmt.Errorf("get list items: %w", err)
	}

	return items, nil
}

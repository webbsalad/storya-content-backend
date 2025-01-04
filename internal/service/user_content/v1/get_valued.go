package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/storya-content-backend/internal/model"
)

func (s *Service) GetValued(ctx context.Context, userID model.UserID, contentType model.ContentType, value model.Value) ([]model.Item, error) {
	items, err := s.contentRepository.GetValued(ctx, userID, contentType, value)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return items, nil
}

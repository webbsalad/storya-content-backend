package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/storya-content-backend/internal/model"
)

func (s *Service) GetRand(ctx context.Context, contentType model.ContentType, count int32) ([]model.Item, error) {
	items, err := s.contentRepository.GetRand(ctx, contentType, count)
	if err != nil {
		return nil, fmt.Errorf("get items: %w", err)

	}

	return items, nil
}

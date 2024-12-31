package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/storya-content-backend/internal/model"
)

func (s *Service) Create(ctx context.Context, item model.Item, contentType model.ContentType) (model.Item, error) {
	item, err := s.contentRepository.Create(ctx, item, contentType)
	if err != nil {
		return model.Item{}, fmt.Errorf("create item: %w", err)
	}

	return item, nil
}

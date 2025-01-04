package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/storya-content-backend/internal/model"
)

func (s *Service) Update(ctx context.Context, item model.Item, contentType model.ContentType) (model.Item, error) {
	updatedItem, err := s.contentRepository.Update(ctx, item, contentType)
	if err != nil {
		return model.Item{}, fmt.Errorf("update: %w", err)
	}

	return updatedItem, nil
}

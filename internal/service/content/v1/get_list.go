package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/storya-content-backend/internal/model"
)

func (s *Service) GetList(ctx context.Context, itemIDs []model.ItemID, contentType model.ContentType) ([]model.Item, error) {
	items, err := s.contentRepository.GetList(ctx, itemIDs, contentType)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return items, nil
}

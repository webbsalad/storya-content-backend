package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/storya-content-backend/internal/model"
)

func (s *Service) Get(ctx context.Context, itemID model.ItemID, contentType model.ContentType) (model.Item, error) {
	item, err := s.contentRepository.Get(ctx, itemID, contentType)
	if err != nil {
		return model.Item{}, fmt.Errorf("get item: %w", err)
	}

	return item, nil
}

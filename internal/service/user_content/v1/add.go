package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/storya-content-backend/internal/model"
)

func (s *Service) Add(ctx context.Context, userID model.UserID, itemID model.ItemID, contentType model.ContentType, value model.Value) (model.ItemID, error) {
	storedItemID, err := s.contentRepository.Add(ctx, userID, itemID, contentType, value)
	if err != nil {
		return model.ItemID{}, fmt.Errorf("add: %w", err)
	}

	return storedItemID, nil
}

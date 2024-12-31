package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/storya-content-backend/internal/model"
)

func (s *Service) Remove(ctx context.Context, userID model.UserID, itemID model.ItemID, contentType model.ContentType) error {
	if err := s.contentRepository.Remove(ctx, userID, itemID, contentType); err != nil {
		return fmt.Errorf("remove: %w", err)
	}

	return nil
}

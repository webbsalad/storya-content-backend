package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/storya-content-backend/internal/model"
)

func (s *Service) Delete(ctx context.Context, itemID model.ItemID, contentType model.ContentType) error {
	if err := s.contentRepository.Delete(ctx, itemID, contentType); err != nil {
		return fmt.Errorf("delete item: %w", err)
	}

	return nil
}

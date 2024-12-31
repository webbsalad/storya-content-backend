package user_content

import (
	"context"

	"github.com/webbsalad/storya-content-backend/internal/model"
)

type Service interface {
	Remove(ctx context.Context, userID model.UserID, itemID model.ItemID, contentType model.ContentType) error
}

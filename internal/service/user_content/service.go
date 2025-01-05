package user_content

import (
	"context"

	"github.com/webbsalad/storya-content-backend/internal/model"
)

type Service interface {
	GetUserItems(ctx context.Context, userID model.UserID, contentType model.ContentType) ([]model.UserItem, error)
	GetValued(ctx context.Context, userID model.UserID, contentType model.ContentType, value model.Value) ([]model.Item, error)
	Remove(ctx context.Context, userID model.UserID, itemID model.ItemID, contentType model.ContentType) error
}

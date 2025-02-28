package content

import (
	"context"

	"github.com/webbsalad/storya-content-backend/internal/model"
)

type Repository interface {
	Get(ctx context.Context, itemID model.ItemID, contentType model.ContentType) (model.Item, error)
	GetList(ctx context.Context, itemIDs []model.ItemID, contentType model.ContentType) ([]model.Item, error)
	GetUserItems(ctx context.Context, userID model.UserID, contentType model.ContentType) ([]model.UserItem, error)
	GetRand(ctx context.Context, contentType model.ContentType, count int32) ([]model.Item, error)
	GetValued(ctx context.Context, userID model.UserID, contentType model.ContentType, value model.Value) ([]model.Item, error)
	Create(ctx context.Context, item model.Item, contentType model.ContentType) (model.Item, error)
	Add(ctx context.Context, userID model.UserID, itemID model.ItemID, contentType model.ContentType, value model.Value) (model.ItemID, error)
	Update(ctx context.Context, item model.Item, contentType model.ContentType) (model.Item, error)
	Delete(ctx context.Context, itemID model.ItemID, contentType model.ContentType) error
	Remove(ctx context.Context, userID model.UserID, itemID model.ItemID, contentType model.ContentType) error
}

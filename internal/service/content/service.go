package content

import (
	"context"

	"github.com/webbsalad/storya-content-backend/internal/model"
)

type Service interface {
	Get(ctx context.Context, itemID model.ItemID, contentType model.ContentType) (model.Item, error)
	GetList(ctx context.Context, itemIDs []model.ItemID, contentType model.ContentType) ([]model.Item, error)
	GetRand(ctx context.Context, contentType model.ContentType, count int32) ([]model.Item, error)
	Create(ctx context.Context, item model.Item, contentType model.ContentType) (model.Item, error)
	Update(ctx context.Context, item model.Item, contentType model.ContentType) (model.Item, error)
	Delete(ctx context.Context, itemID model.ItemID, contentType model.ContentType) error
}

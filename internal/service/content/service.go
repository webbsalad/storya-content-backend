package content

import (
	"context"

	"github.com/webbsalad/storya-content-backend/internal/model"
)

type Service interface {
	Get(ctx context.Context, itemID model.ItemID, contentType model.ContentType) (model.Item, error)
}

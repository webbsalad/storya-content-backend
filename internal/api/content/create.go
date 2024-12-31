package content

import (
	"context"

	"github.com/webbsalad/storya-content-backend/internal/convertor"
	"github.com/webbsalad/storya-content-backend/internal/model"
	"github.com/webbsalad/storya-content-backend/internal/pb/github.com/webbsalad/storya-content-backend/content"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Create(ctx context.Context, req *content.CreateItemRequest) (*content.Item, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	reqItem := model.Item{
		Title: req.GetTitle(),
		Year:  req.GetYear(),
		Tags:  convertor.ToTagsFromDesc(req.GetTags()),
	}

	item, err := i.ContentService.Create(ctx, reqItem, model.ContentType(req.GetType()))
	if err != nil {
		return nil, convertor.ConvertError(err)
	}

	return convertor.ToDescFromItem(item), nil

}

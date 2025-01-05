package content

import (
	"context"

	"github.com/webbsalad/storya-content-backend/internal/convertor"
	"github.com/webbsalad/storya-content-backend/internal/model"
	"github.com/webbsalad/storya-content-backend/internal/pb/github.com/webbsalad/storya-content-backend/content"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Update(ctx context.Context, req *content.UpdateItemRequest) (*content.Item, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	item, err := convertor.UpdateItemRequetToItem(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	updatedItem, err := i.ContentService.Update(ctx, item, model.ContentType(req.GetType()))
	if err != nil {
		return nil, convertor.ConvertError(err)
	}

	return convertor.ToDescFromItem(updatedItem), nil

}

package content

import (
	"context"

	"github.com/webbsalad/storya-content-backend/internal/convertor"
	"github.com/webbsalad/storya-content-backend/internal/model"
	"github.com/webbsalad/storya-content-backend/internal/pb/github.com/webbsalad/storya-content-backend/content"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetList(ctx context.Context, req *content.GetListRequest) (*content.GetListResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	itemIDs, err := convertor.ItemIDsFromStrings(req.ItemIds)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid data: %v", err)
	}

	items, err := i.ContentService.GetList(ctx, itemIDs, model.ContentType(req.ContentType))
	if err != nil {
		return nil, convertor.ConvertError(err)
	}

	return &content.GetListResponse{
		Items: convertor.ToDescFromItems(items),
	}, nil
}

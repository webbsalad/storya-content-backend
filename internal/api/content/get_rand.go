package content

import (
	"context"

	"github.com/webbsalad/storya-content-backend/internal/convertor"
	"github.com/webbsalad/storya-content-backend/internal/model"
	"github.com/webbsalad/storya-content-backend/internal/pb/github.com/webbsalad/storya-content-backend/content"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetRand(ctx context.Context, req *content.GetRandRequest) (*content.GetRandResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	items, err := i.ContentService.GetRand(ctx, model.ContentType(req.GetContentType()), req.GetCount())
	if err != nil {
		return nil, convertor.ConvertError(err)
	}

	return &content.GetRandResponse{
		Items: convertor.ToDescFromItems(items),
	}, nil
}

package user_content

import (
	"context"

	"github.com/webbsalad/storya-content-backend/internal/convertor"
	"github.com/webbsalad/storya-content-backend/internal/model"
	"github.com/webbsalad/storya-content-backend/internal/pb/github.com/webbsalad/storya-content-backend/content"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Add(ctx context.Context, req *content.AddRequest) (*content.AddResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	userID, err := model.UserIDFromString(req.GetUserID())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id data: %v", err)
	}

	itemID, err := model.ItemIDFromString(req.GetItemId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid item id data: %v", err)
	}

	storedItemID, err := i.UserContentService.Add(ctx, userID, itemID, model.ContentType(req.GetContentType()), model.Value(req.GetValue()))
	if err != nil {
		return nil, convertor.ConvertError(err)
	}

	return &content.AddResponse{
		ItemId: storedItemID.String(),
	}, nil

}

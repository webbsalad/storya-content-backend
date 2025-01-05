package user_content

import (
	"context"

	"github.com/webbsalad/storya-content-backend/internal/convertor"
	"github.com/webbsalad/storya-content-backend/internal/model"
	"github.com/webbsalad/storya-content-backend/internal/pb/github.com/webbsalad/storya-content-backend/content"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetUserItems(ctx context.Context, req *content.GetUserItemsRequest) (*content.GetUserItemsResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	userID, err := model.UserIDFromString(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id: %v", err)
	}

	userItems, err := i.UserContentService.GetUserItems(ctx, userID, model.ContentType(req.GetContentType()))
	if err != nil {
		return nil, convertor.ConvertError(err)
	}

	return &content.GetUserItemsResponse{
		UserItems: convertor.ToDescFromUserItems(userItems),
	}, nil

}

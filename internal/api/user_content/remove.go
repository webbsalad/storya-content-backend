package user_content

import (
	"context"

	"github.com/webbsalad/storya-content-backend/internal/convertor"
	"github.com/webbsalad/storya-content-backend/internal/model"
	"github.com/webbsalad/storya-content-backend/internal/pb/github.com/webbsalad/storya-content-backend/content"
	"github.com/webbsalad/storya-content-backend/internal/utils/metadata"
	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Remove(ctx context.Context, req *content.RemoveItemRequest) (*emptypb.Empty, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	userID, err := metadata.GetUserID(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "%v", err)
	}

	itemID, err := model.ItemIDFromString(req.GetItemId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid item id: %v", err)
	}

	if err := i.UserContentService.Remove(ctx, userID, itemID, model.ContentType(req.GetContentType())); err != nil {
		return nil, convertor.ConvertError(err)
	}

	return &emptypb.Empty{}, nil
}

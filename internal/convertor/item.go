package convertor

import (
	"github.com/webbsalad/storya-content-backend/internal/model"
	"github.com/webbsalad/storya-content-backend/internal/pb/github.com/webbsalad/storya-content-backend/content"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToDescFromItem(in model.Item) *content.Item {
	return &content.Item{
		Id:        in.ID.String(),
		Title:     in.Title,
		Year:      in.Year,
		Type:      content.ContentType(in.Type),
		CreatedAt: timestamppb.New(in.CreatedAt),
		Tags:      toDescFromTags(in.Tags),
	}
}

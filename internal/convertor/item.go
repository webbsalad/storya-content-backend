package convertor

import (
	"fmt"

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

func ToDescFromItems(in []model.Item) []*content.Item {
	items := make([]*content.Item, len(in))
	for i, item := range in {
		items[i] = ToDescFromItem(item)
	}
	return items
}

func ToItemFromDesc(in *content.Item) (model.Item, error) {
	itemID, err := model.ItemIDFromString(in.GetId())
	if err != nil {
		return model.Item{}, fmt.Errorf("convert str to item id: %w", err)
	}

	return model.Item{
		ID:    itemID,
		Title: in.Title,
		Year:  in.Year,
		Type:  model.ContentType(in.Type),
		Tags:  ToTagsFromDesc(in.Tags),
	}, nil
}

func ToDescFromUserItem(in model.UserItem) *content.UserItem {
	return &content.UserItem{
		Item:  ToDescFromItem(in.Item),
		Value: content.Value(in.Value),
	}
}

func ToDescFromUserItems(in []model.UserItem) []*content.UserItem {
	userItems := make([]*content.UserItem, len(in))
	for i, userItem := range in {
		userItems[i] = ToDescFromUserItem(userItem)
	}
	return userItems
}

func UpdateItemRequetToItem(in *content.UpdateItemRequest) (model.Item, error) {
	itemID, err := model.ItemIDFromString(in.GetItemId())
	if err != nil {
		return model.Item{}, fmt.Errorf("convert str to item id: %w", err)
	}

	return model.Item{
		ID:    itemID,
		Title: in.GetTitle(),
		Year:  in.GetYear(),
		Type:  model.ContentType(in.Type),
		Tags:  ToTagsFromDesc(in.Tags),
	}, nil
}

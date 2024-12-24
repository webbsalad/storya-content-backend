package pg

import (
	"fmt"

	"github.com/webbsalad/storya-content-backend/internal/model"
)

func toItemFromDB(storedItem Item, storedTags []Tag) (model.Item, error) {
	itemID, err := model.ItemIDFromString(storedItem.ID)
	if err != nil {
		return model.Item{}, fmt.Errorf("convert str to item id: %w", err)
	}

	modelTags := make([]model.Tag, len(storedTags))
	for i, tg := range storedTags {
		modelTags[i] = model.Tag{
			Name: tg.Name,
		}
	}

	return model.Item{
		ID:        itemID,
		Title:     storedItem.Title,
		Year:      int32(storedItem.Year),
		CreatedAt: storedItem.CreatedAt,

		Tags: modelTags,
	}, nil
}

func FromContentTypeToString(ct model.ContentType) string {
	switch ct {
	case model.MOVIE:
		return "movie"
	case model.GAME:
		return "game"
	case model.BOOK:
		return "book"
	default:
		return ""
	}
}

package convertor

import (
	"fmt"

	"github.com/webbsalad/storya-content-backend/internal/model"
)

func ItemIDsFromStrings(in []string) ([]model.ItemID, error) {
	itemIDs := make([]model.ItemID, len(in))
	for i, itemIDStr := range in {
		itemID, err := model.ItemIDFromString(itemIDStr)
		if err != nil {
			return nil, fmt.Errorf("convert item: %w", err)
		}
		itemIDs[i] = itemID
	}

	return itemIDs, nil
}

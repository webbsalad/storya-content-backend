package convertor

import (
	"github.com/webbsalad/storya-content-backend/internal/model"
	"github.com/webbsalad/storya-content-backend/internal/pb/github.com/webbsalad/storya-content-backend/content"
)

func toDescFromTags(in []model.Tag) []*content.Tag {
	tags := make([]*content.Tag, len(in))

	for i, tg := range in {
		tags[i] = &content.Tag{
			Name: tg.Name,
		}
	}

	return tags
}

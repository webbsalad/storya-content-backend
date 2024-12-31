package convertor

import (
	"github.com/webbsalad/storya-content-backend/internal/model"
	"github.com/webbsalad/storya-content-backend/internal/pb/github.com/webbsalad/storya-content-backend/content"
)

func toDescFromTags(in []model.Tag) []*content.Tag {
	tags := make([]*content.Tag, len(in))

	for i, tg := range in {
		tags[i] = toDescFromTag(tg)
	}

	return tags
}

func toDescFromTag(tg model.Tag) *content.Tag {
	return &content.Tag{
		Name: tg.Name,
	}
}

func ToTagsFromDesc(in []*content.Tag) []model.Tag {
	tags := make([]model.Tag, len(in))

	for i, tg := range in {
		tags[i] = toTagFromDesc(tg)
	}

	return tags
}

func toTagFromDesc(tg *content.Tag) model.Tag {
	return model.Tag{
		Name: tg.Name,
	}
}

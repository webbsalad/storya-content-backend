package model

import "time"

type Item struct {
	ID        ItemID
	Type      ContentType
	Title     string
	Year      int32
	CreatedAt time.Time

	Tags []Tag
}

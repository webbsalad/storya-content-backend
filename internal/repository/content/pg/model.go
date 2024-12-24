package pg

import "time"

type Item struct {
	ID        string    `db:"id"`
	Title     string    `db:"title"`
	Year      int       `db:"year"`
	CreatedAt time.Time `db:"created_at"`
	Tags      []Tag
}

type Tag struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

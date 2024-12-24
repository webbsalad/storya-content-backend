package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/webbsalad/storya-content-backend/internal/model"
	"github.com/webbsalad/storya-content-backend/internal/repository/content"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) (content.Repository, error) {
	return &Repository{db: db}, nil
}

func (r *Repository) Get(ctx context.Context, itemID model.ItemID, contentType model.ContentType) (model.Item, error) {
	table := FromContentTypeToString(contentType)
	var storedItem Item
	var tagIDs []string
	var storedTags []Tag

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	itemQuery := psql.
		Select("*").
		From(table).
		Where(
			sq.Eq{"id": itemID.String()},
		)

	q, args, err := itemQuery.ToSql()
	if err != nil {
		return model.Item{}, fmt.Errorf("build query: %w", err)
	}

	if err = r.db.GetContext(ctx, &storedItem, q, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Item{}, model.ErrItemNotFound
		}
		return model.Item{}, fmt.Errorf("get item: %w", err)
	}

	tagsQuery := psql.
		Select("tag_id").
		From(fmt.Sprintf("%s_tags", table)).
		Where(
			sq.Eq{fmt.Sprintf("%s_id", table): itemID.String()},
		)

	q, args, err = tagsQuery.ToSql()
	if err != nil {
		return model.Item{}, fmt.Errorf("build tags query: %w", err)
	}

	if err = r.db.SelectContext(ctx, &tagIDs, q, args...); err != nil {
		return model.Item{}, fmt.Errorf("get tag IDs: %w", err)
	}

	tagsDetailsQuery := psql.
		Select("id", "name").
		From("tag").
		Where(
			sq.Eq{"id": tagIDs},
		)

	q, args, err = tagsDetailsQuery.ToSql()
	if err != nil {
		return model.Item{}, fmt.Errorf("build tags details query: %w", err)
	}

	if err = r.db.SelectContext(ctx, &storedTags, q, args...); err != nil {
		return model.Item{}, fmt.Errorf("get tag details: %w", err)
	}

	item, err := toItemFromDB(storedItem, storedTags)
	if err != nil {
		return model.Item{}, fmt.Errorf("convert to model item: %w", err)
	}

	item.Type = contentType

	return item, nil
}

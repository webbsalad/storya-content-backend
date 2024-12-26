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

func (r *Repository) GetList(ctx context.Context, userID model.UserID, contentType model.ContentType) ([]model.Item, error) {
	table := FromContentTypeToString(contentType)
	var contentIDs []string
	var items []model.Item

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	contentQuery := psql.
		Select("content_id").
		From("user_content").
		Where(
			sq.Eq{"user_id": userID.String(), "content_type": table},
		)

	q, args, err := contentQuery.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build content query: %w", err)
	}

	if err = r.db.SelectContext(ctx, &contentIDs, q, args...); err != nil {
		return nil, fmt.Errorf("get content IDs: %w", err)
	}

	for _, contentID := range contentIDs {
		var storedItem Item
		var tagIDs []string
		var storedTags []Tag

		itemQuery := psql.
			Select("*").
			From(table).
			Where(
				sq.Eq{"id": contentID},
			)

		q, args, err = itemQuery.ToSql()
		if err != nil {
			return nil, fmt.Errorf("build item query: %w", err)
		}

		if err = r.db.GetContext(ctx, &storedItem, q, args...); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				continue
			}
			return nil, fmt.Errorf("get item: %w", err)
		}

		tagsQuery := psql.
			Select("tag_id").
			From(fmt.Sprintf("%s_tags", table)).
			Where(
				sq.Eq{fmt.Sprintf("%s_id", table): contentID},
			)

		q, args, err = tagsQuery.ToSql()
		if err != nil {
			return nil, fmt.Errorf("build tags query: %w", err)
		}

		if err = r.db.SelectContext(ctx, &tagIDs, q, args...); err != nil {
			return nil, fmt.Errorf("get tag IDs: %w", err)
		}

		tagsDetailsQuery := psql.
			Select("id", "name").
			From("tag").
			Where(
				sq.Eq{"id": tagIDs},
			)

		q, args, err = tagsDetailsQuery.ToSql()
		if err != nil {
			return nil, fmt.Errorf("build tags details query: %w", err)
		}

		if err = r.db.SelectContext(ctx, &storedTags, q, args...); err != nil {
			return nil, fmt.Errorf("get tag details: %w", err)
		}

		item, err := toItemFromDB(storedItem, storedTags)
		if err != nil {
			return nil, fmt.Errorf("convert to model item: %w", err)
		}

		item.Type = contentType
		items = append(items, item)
	}

	return items, nil
}

func (r *Repository) GetRand(ctx context.Context, contentType model.ContentType, count int32) ([]model.Item, error) {
	table := FromContentTypeToString(contentType)
	var storedItems []Item
	var items []model.Item

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	randomQuery := psql.
		Select("*").
		From(table).
		OrderBy("RANDOM()").
		Limit(uint64(count))

	q, args, err := randomQuery.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build random query: %w", err)
	}

	if err = r.db.SelectContext(ctx, &storedItems, q, args...); err != nil {
		return nil, fmt.Errorf("get random items: %w", err)
	}

	for _, storedItem := range storedItems {
		var tagIDs []string
		var storedTags []Tag

		tagsQuery := psql.
			Select("tag_id").
			From(fmt.Sprintf("%s_tags", table)).
			Where(
				sq.Eq{fmt.Sprintf("%s_id", table): storedItem.ID},
			)

		q, args, err = tagsQuery.ToSql()
		if err != nil {
			return nil, fmt.Errorf("build tags query: %w", err)
		}

		if err = r.db.SelectContext(ctx, &tagIDs, q, args...); err != nil {
			return nil, fmt.Errorf("get tag IDs: %w", err)
		}

		tagsDetailsQuery := psql.
			Select("id", "name").
			From("tag").
			Where(
				sq.Eq{"id": tagIDs},
			)

		q, args, err = tagsDetailsQuery.ToSql()
		if err != nil {
			return nil, fmt.Errorf("build tags details query: %w", err)
		}

		if err = r.db.SelectContext(ctx, &storedTags, q, args...); err != nil {
			return nil, fmt.Errorf("get tag details: %w", err)
		}

		item, err := toItemFromDB(storedItem, storedTags)
		if err != nil {
			return nil, fmt.Errorf("convert to model item: %w", err)
		}

		item.Type = contentType
		items = append(items, item)
	}

	return items, nil
}

func (r *Repository) Delete(ctx context.Context, itemID model.ItemID, contentType model.ContentType) error {
	table := FromContentTypeToString(contentType)

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Delete(table).
		Where(
			sq.Eq{"id": itemID.String()},
		)

	q, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	res, err := r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("delete item: %w", err)
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("get affected rows: %w", err)
	}

	if rowAffected == 0 {
		return model.ErrItemNotFound
	}

	return nil
}

package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

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

func (r *Repository) GetUserItems(ctx context.Context, userID model.UserID, contentType model.ContentType) ([]model.UserItem, error) {
	table := FromContentTypeToString(contentType)
	var userContent []UserItem
	var userItems []model.UserItem

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	contentQuery := psql.
		Select("content_id", "value").
		From("user_content").
		Where(
			sq.And{
				sq.Eq{"user_id": userID.String()},
				sq.Eq{"content_type": contentType},
			},
		)

	q, args, err := contentQuery.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build content query: %w", err)
	}

	if err = r.db.SelectContext(ctx, &userContent, q, args...); err != nil {
		return nil, fmt.Errorf("get content IDs: %w", err)
	}

	for _, content := range userContent {
		var storedItem Item
		var tagIDs []string
		var storedTags []Tag
		var userItem model.UserItem

		itemQuery := psql.
			Select("*").
			From(table).
			Where(
				sq.Eq{"id": content.ItemID},
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
				sq.Eq{fmt.Sprintf("%s_id", table): content.ItemID},
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
		userItem.Item = item
		userItem.Value = model.Value(content.Value)
		userItems = append(userItems, userItem)
	}

	return userItems, nil
}

func (r *Repository) GetList(ctx context.Context, itemIDs []model.ItemID, contentType model.ContentType) ([]model.Item, error) {
	table := FromContentTypeToString(contentType)
	var items []model.Item

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	for _, itemID := range itemIDs {
		var storedItem Item
		var tagIDs []string
		var storedTags []Tag

		itemQuery := psql.
			Select("*").
			From(table).
			Where(
				sq.Eq{"id": itemID.String()},
			)

		q, args, err := itemQuery.ToSql()
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
				sq.Eq{fmt.Sprintf("%s_id", table): itemID.String()},
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

func (r *Repository) GetValued(ctx context.Context, userID model.UserID, contentType model.ContentType, value model.Value) ([]model.Item, error) {
	table := FromContentTypeToString(contentType)
	var contentIDs []string
	var items []model.Item

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	contentQuery := psql.
		Select("content_id").
		From("user_content").
		Where(
			sq.And{
				sq.Eq{"user_id": userID.String()},
				sq.Eq{"content_type": contentType},
				sq.Eq{"value": value},
			},
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

func (r *Repository) Create(ctx context.Context, item model.Item, contentType model.ContentType) (model.Item, error) {
	table := FromContentTypeToString(contentType)

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.Item{}, fmt.Errorf("begin transaction: %w", err)
	}

	defer func() {
		rollbackErr := tx.Rollback()
		if err == nil && rollbackErr != nil {
			err = fmt.Errorf("rollback: %w", rollbackErr)
		}
	}()

	var tagIDs []string
	for _, tag := range item.Tags {
		var tagID string

		tagQuery := psql.
			Select("id").
			From("tag").
			Where(sq.Eq{"name": tag.Name})

		q, args, err := tagQuery.ToSql()
		if err != nil {
			return model.Item{}, fmt.Errorf("build tag query: %w", err)
		}

		err = tx.GetContext(ctx, &tagID, q, args...)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				insertTagQuery := psql.
					Insert("tag").
					Columns("name").
					Values(tag.Name).
					Suffix("RETURNING id")

				q, args, err = insertTagQuery.ToSql()
				if err != nil {
					return model.Item{}, fmt.Errorf("build insert tag query: %w", err)
				}

				err = tx.GetContext(ctx, &tagID, q, args...)
				if err != nil {
					return model.Item{}, fmt.Errorf("insert tag: %w", err)
				}
			} else {
				return model.Item{}, fmt.Errorf("get tag: %w", err)
			}
		}

		tagIDs = append(tagIDs, tagID)
	}

	insertItemQuery := psql.
		Insert(table).
		Columns("title", "year", "created_at").
		Values(item.Title, item.Year, time.Now()).
		Suffix("RETURNING id")

	q, args, err := insertItemQuery.ToSql()
	if err != nil {
		return model.Item{}, fmt.Errorf("build insert item query: %w", err)
	}

	var itemID string
	err = tx.GetContext(ctx, &itemID, q, args...)
	if err != nil {
		return model.Item{}, fmt.Errorf("insert item: %w", err)
	}

	for _, tagID := range tagIDs {
		insertTagLinkQuery := psql.
			Insert(fmt.Sprintf("%s_tags", table)).
			Columns(fmt.Sprintf("%s_id", table), "tag_id").
			Values(itemID, tagID)

		q, args, err := insertTagLinkQuery.ToSql()
		if err != nil {
			return model.Item{}, fmt.Errorf("build insert tag link query: %w", err)
		}

		_, err = tx.ExecContext(ctx, q, args...)
		if err != nil {
			return model.Item{}, fmt.Errorf("insert tag link: %w", err)
		}
	}

	item.ID, err = model.ItemIDFromString(itemID)
	if err != nil {
		return model.Item{}, fmt.Errorf("convert str to item id: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return model.Item{}, fmt.Errorf("commit transaction: %w", err)
	}

	return item, nil
}

func (r *Repository) Add(ctx context.Context, userID model.UserID, itemID model.ItemID, contentType model.ContentType, value model.Value) (model.ItemID, error) {
	var strItemID string

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Insert("user_content").
		Columns("user_id", "content_id", "content_type", "value", "created_at").
		Values(userID.String(), itemID.String(), contentType, value, time.Now()).
		Suffix("ON CONFLICT (user_id, content_id, content_type) DO NOTHING RETURNING id")

	q, args, err := query.ToSql()
	if err != nil {
		return model.ItemID{}, fmt.Errorf("build query: %w", err)
	}

	err = r.db.GetContext(ctx, &strItemID, q, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ItemID{}, model.ErrUserContentAlreadyExist
		}
		return model.ItemID{}, fmt.Errorf("add user item: %w", err)
	}

	storedItem, err := model.ItemIDFromString(strItemID)
	if err != nil {
		return model.ItemID{}, fmt.Errorf("convert str to item id: %w", err)
	}

	return storedItem, nil
}

func (r *Repository) Update(ctx context.Context, item model.Item, contentType model.ContentType) (model.Item, error) {
	table := FromContentTypeToString(contentType)

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.Item{}, fmt.Errorf("begin transaction: %w", err)
	}

	defer func() {
		rollbackErr := tx.Rollback()
		if err == nil && rollbackErr != nil {
			err = fmt.Errorf("rollback: %w", rollbackErr)
		}
	}()

	updateItemQuery := psql.
		Update(table).
		Set("title", item.Title).
		Set("year", item.Year).
		Where(sq.Eq{"id": item.ID.String()})

	q, args, err := updateItemQuery.ToSql()
	if err != nil {
		return model.Item{}, fmt.Errorf("build update item query: %w", err)
	}

	_, err = tx.ExecContext(ctx, q, args...)
	if err != nil {
		return model.Item{}, fmt.Errorf("update item: %w", err)
	}

	deleteTagLinksQuery := psql.
		Delete(fmt.Sprintf("%s_tags", table)).
		Where(sq.Eq{fmt.Sprintf("%s_id", table): item.ID.String()})

	q, args, err = deleteTagLinksQuery.ToSql()
	if err != nil {
		return model.Item{}, fmt.Errorf("build delete tag links query: %w", err)
	}

	_, err = tx.ExecContext(ctx, q, args...)
	if err != nil {
		return model.Item{}, fmt.Errorf("delete tag links: %w", err)
	}

	var tagIDs []string
	for _, tag := range item.Tags {
		var tagID string
		tagQuery := psql.
			Select("id").
			From("tag").
			Where(sq.Eq{"name": tag.Name})

		q, args, err := tagQuery.ToSql()
		if err != nil {
			return model.Item{}, fmt.Errorf("build tag query: %w", err)
		}

		err = tx.GetContext(ctx, &tagID, q, args...)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				insertTagQuery := psql.
					Insert("tag").
					Columns("name").
					Values(tag.Name).
					Suffix("RETURNING id")

				q, args, err = insertTagQuery.ToSql()
				if err != nil {
					return model.Item{}, fmt.Errorf("build insert tag query: %w", err)
				}

				err = tx.GetContext(ctx, &tagID, q, args...)
				if err != nil {
					return model.Item{}, fmt.Errorf("insert tag: %w", err)
				}
			} else {
				return model.Item{}, fmt.Errorf("get tag: %w", err)
			}
		}

		tagIDs = append(tagIDs, tagID)
	}

	for _, tagID := range tagIDs {
		insertTagLinkQuery := psql.
			Insert(fmt.Sprintf("%s_tags", table)).
			Columns(fmt.Sprintf("%s_id", table), "tag_id").
			Values(item.ID.String(), tagID)

		q, args, err := insertTagLinkQuery.ToSql()
		if err != nil {
			return model.Item{}, fmt.Errorf("build insert tag link query: %w", err)
		}

		_, err = tx.ExecContext(ctx, q, args...)
		if err != nil {
			return model.Item{}, fmt.Errorf("insert tag link: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return model.Item{}, fmt.Errorf("commit transaction: %w", err)
	}

	return item, nil
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

func (r *Repository) Remove(ctx context.Context, userID model.UserID, itemID model.ItemID, contentType model.ContentType) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Delete("user_content").
		Where(
			sq.And{
				sq.Eq{"content_id": itemID.String()},
				sq.Eq{"user_id": userID.String()},
				sq.Eq{"content_type": contentType},
			},
		)

	q, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	res, err := r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("remove item: %w", err)
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

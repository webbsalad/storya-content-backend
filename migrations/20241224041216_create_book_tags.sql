-- +goose Up
-- +goose StatementBegin
CREATE TABLE book_tags (
    book_id UUID NOT NULL REFERENCES book(id) ON DELETE CASCADE,
    tag_id UUID NOT NULL REFERENCES tag(id) ON DELETE CASCADE,
    PRIMARY KEY (book_id, tag_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE book_tags;
-- +goose StatementEnd

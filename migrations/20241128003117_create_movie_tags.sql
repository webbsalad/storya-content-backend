-- +goose Up
-- +goose StatementBegin
CREATE TABLE movie_tags (
    movie_id UUID NOT NULL REFERENCES movie(id) ON DELETE CASCADE,
    tag_id UUID NOT NULL REFERENCES tag(id) ON DELETE CASCADE,
    PRIMARY KEY (movie_id, tag_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE movie_tags;
-- +goose StatementEnd

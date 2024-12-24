-- +goose Up
-- +goose StatementBegin
CREATE TABLE game_tags (
    game_id UUID NOT NULL REFERENCES game(id) ON DELETE CASCADE,
    tag_id UUID NOT NULL REFERENCES tag(id) ON DELETE CASCADE,
    PRIMARY KEY (game_id, tag_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE game_tags;
-- +goose StatementEnd

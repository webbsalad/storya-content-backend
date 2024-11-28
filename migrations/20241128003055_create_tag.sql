-- +goose Up
-- +goose StatementBegin
CREATE TABLE tag (
    id UUID DEFAULT uuid_generate_v4() NOT NULL CONSTRAINT tag_pkey PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tag;
-- +goose StatementEnd

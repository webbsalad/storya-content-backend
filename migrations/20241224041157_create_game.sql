-- +goose Up
-- +goose StatementBegin
CREATE TABLE game (
    id UUID DEFAULT uuid_generate_v4() NOT NULL CONSTRAINT game_pkey PRIMARY KEY,
    title TEXT NOT NULL,
    year INT CHECK (year >= 1800 AND year <= EXTRACT(YEAR FROM CURRENT_DATE)) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE game;
-- +goose StatementEnd

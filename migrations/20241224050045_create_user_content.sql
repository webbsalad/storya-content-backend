-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_content (
    id UUID DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL,
    content_id UUID NOT NULL,
    content_type TEXT CHECK (content_type IN ('game', 'book', 'movie')) NOT NULL,
    rating INT CHECK (rating >= -10 AND rating <= 10),
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,

    CONSTRAINT fk_user_content_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_content;
-- +goose StatementEnd

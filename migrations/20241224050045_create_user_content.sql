-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_content (
    id UUID DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL,
    content_id UUID NOT NULL,
    content_type INT CHECK (content_type >= 0 AND content_type <= 2) NOT NULL,
    value INT CHECK (value >= 0 AND value <= 2),
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,

    CONSTRAINT fk_user_content_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT unique_user_content UNIQUE (user_id, content_id, content_type)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_content;
-- +goose StatementEnd

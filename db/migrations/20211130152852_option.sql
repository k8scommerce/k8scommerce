-- +goose Up
-- +goose StatementBegin
CREATE TABLE option (
    id bigserial PRIMARY KEY,
    store_id bigint NOT NULL,
    name character varying NOT NULL,
    display_name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    UNIQUE (name)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE option;

-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE property (
    id bigserial PRIMARY KEY,
    name character varying NOT NULL,
    display_name character varying NOT NULL,
    fiterable boolean DEFAULT TRUE NOT NULL,
    filter_param character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    UNIQUE (name)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE property;

-- +goose StatementEnd

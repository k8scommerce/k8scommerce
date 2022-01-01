-- +goose Up
-- +goose StatementBegin
CREATE TABLE option_item (
    id bigserial PRIMARY KEY,
    option_id bigint NOT NULL,
    name character varying NOT NULL,
    display_name character varying NOT NULL,
    sort_order integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    UNIQUE (name),
    FOREIGN KEY (option_id) REFERENCES option (id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE option_item;

-- +goose StatementEnd

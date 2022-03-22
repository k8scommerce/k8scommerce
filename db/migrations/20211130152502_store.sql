-- +goose Up
--
-- +goose StatementBegin
CREATE TABLE store (
    id bigserial PRIMARY KEY,
    is_default boolean DEFAULT FALSE NOT NULL,
    name character varying NOT NULL,
    description text,
    url character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    deleted_at timestamp without time zone,
    UNIQUE (name, url)
);
-- +goose StatementEnd
--
--
-- +goose StatementBegin
CREATE TABLE store_setting (
    id bigserial PRIMARY KEY,
    store_id bigint not null,
    config JSONB default '{}' NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    deleted_at timestamp without time zone,
    FOREIGN KEY (store_id) REFERENCES store (id) ON DELETE CASCADE,
    UNIQUE (store_id)
);
-- +goose StatementEnd
--
--
--
-- +goose Down
--
-- +goose StatementBegin
DROP TABLE store;
DROP TABLE store_setting;
-- +goose StatementEnd
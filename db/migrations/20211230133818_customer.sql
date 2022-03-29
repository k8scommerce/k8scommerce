-- +goose Up
--
--
-- +goose StatementBegin
CREATE TABLE customer (
    id bigserial PRIMARY KEY,
    store_id bigint NOT NULL,
    first_name character varying NOT NULL,
    last_name character varying NOT NULL,
    company character varying NULL,
    phone character varying NULL,
    email character varying NOT NULL UNIQUE,
    password character varying NULL,
    is_verified boolean NOT NULL default false,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    deleted_at timestamp without time zone,
    UNIQUE(store_id, email)
);
CREATE INDEX IF NOT EXISTS idx_customer_store_id ON customer USING btree (store_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE customer_address (
    id bigserial PRIMARY KEY,
    store_id bigint NOT NULL,
    customer_id bigint NOT NULL,
    kind address_kind NOT NULL,
    is_default boolean DEFAULT FALSE NOT NULL,
    street character varying NOT NULL,
    apt_suite character varying NULL,
    city character varying NOT NULL,
    state_province character varying NOT NULL,
    country character varying NOT NULL,
    postal_code character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    deleted_at timestamp without time zone
);
CREATE INDEX IF NOT EXISTS idx_customer_address_store_id ON customer_address USING btree (store_id);
CREATE INDEX IF NOT EXISTS idx_customer_address_customer_id_kind ON customer_address USING btree (customer_id, kind);
-- +goose StatementEnd
--2
--
--
--
-- +goose Down
-- +goose StatementBegin
DROP TABLE customer;
DROP TABLE customer_address;
-- +goose StatementEnd
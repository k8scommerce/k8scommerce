-- +goose Up
-- +goose StatementBegin
CREATE TABLE inventory_brand (
    id bigserial PRIMARY KEY,
    store_id bigint NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    deleted_at timestamp without time zone
);
-- +goose StatementEnd
--
-- +goose StatementBegin
CREATE TABLE inventory_supplier (
    id bigserial PRIMARY KEY,
    store_id bigint NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    deleted_at timestamp without time zone
);
-- +goose StatementEnd
--
-- +goose StatementBegin
CREATE TABLE inventory_stock (
    id bigserial PRIMARY KEY,
    store_id bigint NOT NULL,
    warehouse_id bigint NOT NULL,
    inventory_item_id bigint NOT NULL,
    quantity int NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    deleted_at timestamp without time zone
);
-- +goose StatementEnd
--
-- +goose StatementBegin
CREATE TABLE inventory_item (
    id bigserial PRIMARY KEY,
    store_id bigint NOT NULL,
    sku character varying NOT NULL,
    name character varying NOT NULL,
    supplier_id bigint NOT NULL,
    brand_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    deleted_at timestamp without time zone
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE inventory_item CASCADE;
DROP TABLE IF EXISTS inventory_stock CASCADE;
DROP TABLE IF EXISTS inventory_supplier;
DROP TABLE IF EXISTS inventory_brand;
-- +goose StatementEnd
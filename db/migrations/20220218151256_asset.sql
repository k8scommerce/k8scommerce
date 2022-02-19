-- +goose Up
--
--
-- +goose StatementBegin
CREATE TYPE asset_kind AS ENUM ('image', 'document', 'audio', 'video', 'archive');
-- +goose StatementEnd
--
--
-- +goose StatementBegin
CREATE TABLE asset (
    id bigserial PRIMARY KEY,
    store_id bigint NOT NULL,
    product_id bigint NOT NULL,
    variant_id bigint NOT NULL,
    name character varying NOT NULL,
    url character varying NOT NULL,
    display_name character varying,
    kind asset_kind NOT NULL,
    content_type character varying NOT NULL,
    sort_order integer default 100,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    UNIQUE (store_id, name)
);
CREATE INDEX IF NOT EXISTS idx_asset_product_id ON asset USING btree (product_id);
CREATE INDEX IF NOT EXISTS idx_asset_variant_id ON asset USING btree (variant_id);
CREATE INDEX IF NOT EXISTS idx_asset_variant_id_kind ON asset USING btree (variant_id, kind);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS asset_kind;
DROP TABLE IF EXISTS asset;
-- +goose StatementEnd
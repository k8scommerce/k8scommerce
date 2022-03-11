-- +goose Up
-- +goose StatementBegin
CREATE TABLE variant (
    id bigserial PRIMARY KEY,
    product_id bigint NOT NULL,
    is_default boolean DEFAULT FALSE NOT NULL,
    sku character varying DEFAULT ''::character varying NOT NULL,
    sort_order integer NOT NULL,
    cost_amount bigint,
    cost_currency character varying DEFAULT 'USD',
    track_inventory boolean DEFAULT TRUE NOT NULL,
    tax_category_id bigint,
    shipping_category_id bigint,
    discontinue_on timestamp without time zone,
    weight numeric(8, 2) DEFAULT 0.0,
    height numeric(8, 2),
    width numeric(8, 2),
    depth numeric(8, 2),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    deleted_at timestamp without time zone,
    UNIQUE (sku),
    FOREIGN KEY (product_id) REFERENCES product (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_variant_product_id ON variant USING btree (product_id);
CREATE INDEX IF NOT EXISTS idx_variant_is_default ON variant USING btree (is_default);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE variant;

-- +goose StatementEnd

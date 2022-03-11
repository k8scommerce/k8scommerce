-- +goose Up
-- +goose StatementBegin
CREATE TABLE price (
    id bigserial PRIMARY KEY,
    variant_id bigint NOT NULL,
    sale_price bigint NOT NULL,
    retail_price bigint,
    currency character varying DEFAULT 'USD',
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    deleted_at timestamp without time zone,
    FOREIGN KEY (variant_id) REFERENCES variant (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_price_variant_id ON price USING btree (variant_id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE price;

-- +goose StatementEnd

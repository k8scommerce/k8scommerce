-- +goose Up
-- +goose StatementBegin
CREATE TABLE product (
    id bigserial PRIMARY KEY,
    store_id bigint NOT NULL,
    slug character varying NOT NULL,
    name character varying NOT NULL,
    short_description text,
    description text,
    meta_title character varying,
    meta_description text,
    meta_keywords character varying,
    promotionable boolean DEFAULT TRUE NOT NULL,
    featured boolean DEFAULT FALSE NOT NULL,
    available_on timestamp without time zone NULL,
    discontinue_on timestamp without time zone,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    deleted_at timestamp without time zone,
    UNIQUE (slug)
);

CREATE INDEX IF NOT EXISTS idx_product_store_id ON product USING btree (store_id);
CREATE INDEX IF NOT EXISTS idx_product_slug ON product USING btree (slug);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE product;

-- +goose StatementEnd

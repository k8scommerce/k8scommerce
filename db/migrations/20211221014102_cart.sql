-- +goose Up
-- +goose StatementBegin
DO $$ BEGIN CREATE TYPE cart_status AS enum (
    'new',
    'checkout',
    'paid',
    'complete',
    'abandoned'
);
EXCEPTION
WHEN duplicate_object THEN null;
END $$;
-- +goose StatementEnd
--
--
--
-- +goose StatementBegin
CREATE TABLE cart (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    store_id bigint NOT NULL,
    customer_id bigint NULL,
    first_name character varying NULL,
    last_name character varying NULL,
    company character varying NULL,
    phone character varying NULL,
    email character varying NULL,
    billing_address JSONB default '{}' NOT NULL,
    shipping_address JSONB default '{}' NOT NULL,
    discount_id bigint NULL,
    status cart_status DEFAULT 'new' NOT NULL,
    currency character varying DEFAULT 'USD',
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    expires_at timestamp without time zone NULL,
    abandoned_at timestamp without time zone NULL,
    FOREIGN KEY (store_id) REFERENCES store (id),
    FOREIGN KEY (customer_id) REFERENCES customer (id)
);
CREATE INDEX IF NOT EXISTS idx_cart_store_id ON cart USING btree (store_id);
CREATE INDEX IF NOT EXISTS idx_cart_customer_id ON cart USING btree (customer_id);
-- +goose StatementEnd
--
--
--
-- +goose StatementBegin
CREATE TABLE cart_item (
    cart_id uuid NOT NULL,
    sku character varying DEFAULT ''::character varying NOT NULL,
    quantity int NOT NULL default 1,
    price bigint NOT NULL default 0,
    currency character varying DEFAULT 'USD',
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    expires_at timestamp without time zone NULL,
    abandoned_at timestamp without time zone NULL,
    PRIMARY KEY (cart_id, sku),
    FOREIGN KEY (cart_id) REFERENCES cart (id),
    FOREIGN KEY (sku) REFERENCES variant (sku)
);
CREATE INDEX IF NOT EXISTS idx_cart_item_cart_id ON cart_item USING btree (cart_id);
CREATE INDEX IF NOT EXISTS idx_cart_item_sku ON cart_item USING btree (sku);
CREATE INDEX IF NOT EXISTS idx_cart_item_abandoned_at ON cart_item USING btree (abandoned_at);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE cart_item;
DROP TABLE cart;
DROP TYPE IF EXISTS cart_status;
-- +goose StatementEnd
-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS customer_password_reset (
    id bigserial PRIMARY KEY,
    store_id bigint NOT NULL,
    customer_id bigint NOT NULL,
    token character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    redeemed_at timestamp without time zone,
    expired_at timestamp without time zone,
    UNIQUE (token),
    FOREIGN KEY (customer_id) REFERENCES customer (id)
);
CREATE INDEX IF NOT EXISTS idx_customer_password_reset_store_id ON customer_password_reset USING btree (store_id);
-- +goose StatementEnd
--
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users_password_reset (
    id bigserial PRIMARY KEY,
    users_id bigint NOT NULL,
    token character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    redeemed_at timestamp without time zone,
    expired_at timestamp without time zone,
    UNIQUE (token),
    FOREIGN KEY (users_id) REFERENCES users (id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS customer_password_reset;
DROP TABLE IF EXISTS users_password_reset;
-- +goose StatementEnd
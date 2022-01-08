-- +goose Up
-- +goose StatementBegin
CREATE TABLE cart (
    customer_id bigint NOT NULL PRIMARY KEY,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL
);
-- +goose StatementEnd
--
--
--
-- +goose StatementBegin
CREATE TABLE cart_item (
    customer_id bigint NOT NULL,
    sku character varying DEFAULT ''::character varying NOT NULL,
    quantity int NOT NULL default 1,
    price bigint NOT NULL default 0,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    expires_at timestamp without time zone NOT NULL,
    abandoned_at timestamp without time zone NULL,
    PRIMARY KEY (customer_id, sku),
    FOREIGN KEY (customer_id) REFERENCES cart (customer_id),
    FOREIGN KEY (sku) REFERENCES variant (sku)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE cart_item;
DROP TABLE cart;
-- +goose StatementEnd
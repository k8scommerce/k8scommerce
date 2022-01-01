-- +goose Up
-- +goose StatementBegin
CREATE TABLE store_product (
    store_id bigint,
    product_id bigint,
    PRIMARY KEY (store_id, product_id),
    FOREIGN KEY (store_id) REFERENCES store (id),
    FOREIGN KEY (product_id) REFERENCES product (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE store_product;
-- +goose StatementEnd

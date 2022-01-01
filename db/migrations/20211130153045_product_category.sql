-- +goose Up
-- +goose StatementBegin
CREATE TABLE product_category (
    product_id bigint,
    category_id bigint,
    PRIMARY KEY (product_id, category_id),
    FOREIGN KEY (product_id) REFERENCES product (id),
    FOREIGN KEY (category_id) REFERENCES category (id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE product_category;

-- +goose StatementEnd

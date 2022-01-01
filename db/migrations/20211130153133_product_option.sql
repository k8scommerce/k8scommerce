-- +goose Up
-- +goose StatementBegin
CREATE TABLE product_option (
    product_id bigint,
    option_id bigint,
    PRIMARY KEY (product_id, option_id),
    FOREIGN KEY (product_id) REFERENCES product (id),
    FOREIGN KEY (option_id) REFERENCES option (id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE product_option;

-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE product_property (
    product_id bigint,
    property_id bigint,
    PRIMARY KEY (product_id, property_id),
    FOREIGN KEY (product_id) REFERENCES product (id),
    FOREIGN KEY (property_id) REFERENCES property (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE product_property;
-- +goose StatementEnd

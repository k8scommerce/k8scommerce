-- +goose Up
-- +goose StatementBegin
CREATE TABLE price (
    id bigserial PRIMARY KEY,
    variant_id bigint NOT NULL,
    amount bigint NOT NULL,
    compare_at_amount bigint,
    currency character varying DEFAULT 'USD',
    user_role_id bigint,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    deleted_at timestamp without time zone,
    FOREIGN KEY (variant_id) REFERENCES variant (id) ON DELETE CASCADE
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE price;

-- +goose StatementEnd

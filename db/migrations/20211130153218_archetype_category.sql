-- +goose Up
-- +goose StatementBegin
CREATE TABLE archetype_category (
    archetype_id bigint,
    category_id bigint,
    PRIMARY KEY (archetype_id, category_id),
    FOREIGN KEY (archetype_id) REFERENCES archetype (id),
    FOREIGN KEY (category_id) REFERENCES category (id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE archetype_category;

-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE archetype_option (
    archetype_id bigint,
    option_id bigint,
    PRIMARY KEY (archetype_id, option_id),
    FOREIGN KEY (archetype_id) REFERENCES archetype (id),
    FOREIGN KEY (option_id) REFERENCES option (id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE archetype_option;

-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE archetype_property (
    archetype_id bigint,
    property_id bigint,
    PRIMARY KEY (archetype_id, property_id),
    FOREIGN KEY (archetype_id) REFERENCES archetype (id),
    FOREIGN KEY (property_id) REFERENCES property (id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE archetype_property;

-- +goose StatementEnd

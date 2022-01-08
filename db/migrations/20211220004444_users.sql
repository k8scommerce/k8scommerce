-- +goose Up
--
--
-- +goose StatementBegin
CREATE TABLE users (
    id bigserial PRIMARY KEY,
    store_id bigint NOT NULL,
    first_name character varying NOT NULL,
    last_name character varying NOT NULL,
    email character varying NOT NULL UNIQUE,
    password character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL
);
-- +goose StatementEnd
--
--
-- +goose StatementBegin
CREATE TABLE permission_group (
    id bigserial PRIMARY KEY,
    store_id bigint NOT NULL,
    group_name character varying NOT NULL UNIQUE,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL
);
-- +goose StatementEnd
--
--
-- +goose StatementBegin
CREATE TABLE users_permission_group (
    user_id bigserial,
    permission_group_id bigserial,
    PRIMARY KEY (user_id, permission_group_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (permission_group_id) REFERENCES permission_group (id)
);
-- +goose StatementEnd
--
--
-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE users_permission_group;
DROP TABLE permission_group;
-- +goose StatementEnd
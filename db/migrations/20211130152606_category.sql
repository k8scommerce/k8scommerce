-- +goose Up
-- +goose StatementBegin
CREATE TABLE category (
    id bigserial PRIMARY KEY,
    parent_id bigint,
    store_id bigint NOT NULL,
    slug character varying NOT NULL,
    name character varying NOT NULL,
    description text,
    meta_title character varying,
    meta_description character varying,
    meta_keywords character varying,
    hide_from_nav boolean DEFAULT FALSE,
    lft bigint,
    rgt bigint,
    depth integer,
    sort_order integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    FOREIGN KEY (parent_id) REFERENCES category (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_category_parent_id ON category USING btree (parent_id);

CREATE INDEX IF NOT EXISTS idx_category_lft ON category USING btree (lft);

CREATE INDEX IF NOT EXISTS idx_category_rgt ON category USING btree (rgt);

CREATE INDEX IF NOT EXISTS idx_category_depth ON category USING btree (depth);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE category;

-- +goose StatementEnd

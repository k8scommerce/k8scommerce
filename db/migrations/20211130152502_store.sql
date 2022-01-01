-- +goose Up
--
-- +goose StatementBegin
CREATE TABLE store (
    id bigserial PRIMARY KEY,
    is_default boolean DEFAULT FALSE NOT NULL,
    name character varying NOT NULL,
    description text,
    url character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    deleted_at timestamp without time zone,
    UNIQUE (name, url)
);
-- +goose StatementEnd
--
--
-- +goose StatementBegin
CREATE TABLE store_setting (
    id bigserial PRIMARY KEY,
    store_id bigint not null,
    seo_title character varying,
    seo_robots character varying,
    meta_description text,
    meta_keywords text,
    facebook character varying,
    twitter character varying,
    instagram character varying,
    code character varying,
    default_currency character varying DEFAULT 'USD' NOT NULL,
    supported_currencies character varying,
    default_locale character varying DEFAULT 'America/Denver' NOT NULL,
    supported_locales character varying,
    default_country_id bigint DEFAULT 1 NOT NULL,
    contact_phone character varying,
    mail_from_address character varying,
    customer_support_email character varying,
    new_order_notifications_email character varying,
    checkout_zone_id bigint,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    deleted_at timestamp without time zone
);
-- +goose StatementEnd--
--
--
-- +goose StatementBegin
CREATE TABLE store_address (
    id bigserial PRIMARY KEY,
    store_id bigint not null,
    kind address_kind NOT NULL,
    is_default boolean DEFAULT FALSE NOT NULL,
    address text,
    city text,
    state_province character varying,
    postal_code character varying,
    country character varying,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    deleted_at timestamp without time zone
);
-- +goose StatementEnd
--
-- +goose Down
--
-- +goose StatementBegin
DROP TABLE store;
-- +goose StatementEnd
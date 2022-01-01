-- +goose Up
-- +goose StatementBegin
DO $$ BEGIN
CREATE TYPE transaction_kind AS enum (
    'unknown',
    'sale',
    'refund',
    'chargeback'
);
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;
-- +goose StatementEnd
--
-- +goose StatementBegin
CREATE TABLE payment_gateway (
    id bigserial PRIMARY KEY,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NULL,
    deleted_at timestamp without time zone
);
-- +goose StatementEnd
--
-- +goose StatementBegin
CREATE TABLE payment_transaction (
    id bigserial PRIMARY KEY,
    store_id bigint not null,
    gateway_id bigint,
    reference_code character varying NOT NULL,
    auth_code character varying NOT NULL,
    response  character varying NOT NULL,
    amount bigint,
    currency character varying DEFAULT 'USD',
    kind transaction_kind NOT NULL,
    is_recurring boolean default false,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_payment_transaction_reference_code ON payment_transaction USING btree (reference_code);
CREATE INDEX idx_payment_transaction_kind ON payment_transaction USING btree (kind);
CREATE INDEX idx_payment_transaction_created_at ON payment_transaction ((created_at::DATE));
-- +goose StatementEnd
--
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION before_payment_transaction_update() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at := CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$;

DROP TRIGGER IF EXISTS trigger_payment_transaction_update ON payment_transaction;
CREATE TRIGGER trigger_payment_transaction_update BEFORE UPDATE ON payment_transaction FOR EACH ROW EXECUTE FUNCTION before_payment_transaction_update();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE payment_gateway;
DROP TYPE payment_transaction;
DROP TYPE transaction_kind;
DROP TRIGGER before_payment_transaction_update;
-- +goose StatementEnd

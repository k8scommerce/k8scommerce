-- +goose Up
-- +goose StatementBegin
SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = ON;
SET check_function_bodies = FALSE;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = OFF;
SET TIME ZONE 'UTC';
-- +goose StatementEnd
--
--
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS citext;
COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';
-- +goose StatementEnd
--
--
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION empty (text) RETURNS boolean LANGUAGE sql IMMUTABLE AS $_$
SELECT $1 ~ '^[[:space:]]*$';
$_$;
COMMENT ON FUNCTION empty (text) IS 'Find empty strings or strings containing only whitespace';
-- +goose StatementEnd
--
--
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_max_bigint_value () RETURNS bigint AS $$
DECLARE max_bigint_value BIGINT;
BEGIN
SELECT -(
        (
            (2 ^(8 * pg_column_size(1::BIGINT) -2))::BIGINT << 1
        ) + 1
    ) INTO max_bigint_value;
RETURN max_bigint_value;
END $$ LANGUAGE "plpgsql";
-- +goose StatementEnd
--
--
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_min_bigint_value () RETURNS bigint AS $$
DECLARE min_bigint_value BIGINT;
BEGIN
SELECT (2 ^(8 * pg_column_size(1::bigint) -2))::bigint << 1 INTO min_bigint_value;
RETURN min_bigint_value;
END $$ LANGUAGE "plpgsql";
-- +goose StatementEnd
--
--
-- +goose StatementBegin
CREATE OR REPLACE PROCEDURE build_h_tally_table () LANGUAGE plpgsql AS $$ BEGIN DROP TABLE IF EXISTS h_tally;
CREATE TABLE h_tally (n int PRIMARY KEY);
INSERT INTO h_tally
SELECT COALESCE(
        CAST(
            (
                ROW_NUMBER() OVER (
                    ORDER BY (
                            SELECT NULL
                        )
                ) - 1
            ) * 32 + 1 AS INT
        ),
        0
    ) AS n
FROM information_schema.columns ac1
    CROSS JOIN information_schema.columns ac2
LIMIT 5000;
END $$;
-- +goose StatementEnd
--
--
-- +goose StatementBegin
CALL build_h_tally_table ();
-- +goose StatementEnd
--
--
-- +goose StatementBegin
CREATE TYPE money AS (number NUMERIC, currency_code CHAR(3));
-- +goose StatementEnd
--
--
-- +goose StatementBegin
CREATE TYPE address_kind AS ENUM ('billing', 'shipping', 'mailing');
-- +goose StatementEnd
--
--
-- +goose StatementBegin
CREATE EXTENSION pgcrypto;
-- +goose StatementEnd
--
--
-- +goose Down
DROP EXTENSION IF EXISTS citext;
DROP EXTENSION IF EXISTS pgcrypto;
DROP FUNCTION IF EXISTS empty;
DROP FUNCTION IF EXISTS get_max_bigint_value;
DROP FUNCTION IF EXISTS get_min_bigint_value;
DROP TYPE IF EXISTS address_kind;
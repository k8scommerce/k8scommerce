-- +goose Up
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
-- CREATE SCHEMA IF NOT EXISTS tiger;
-- CREATE SCHEMA IF NOT EXISTS tiger_data;
-- CREATE SCHEMA IF NOT EXISTS topology;
-- COMMENT ON SCHEMA topology IS 'PostGIS Topology schema';
-- CREATE EXTENSION IF NOT EXISTS address_standardizer;
-- COMMENT ON EXTENSION address_standardizer IS 'Used to parse an address into constituent elements. Generally used to support geocoding address normalization step.';
-- CREATE EXTENSION IF NOT EXISTS address_standardizer_data_us;
-- COMMENT ON EXTENSION address_standardizer_data_us IS 'Address Standardizer US dataset example';
-- CREATE EXTENSION IF NOT EXISTS btree_gist;
-- COMMENT ON EXTENSION btree_gist IS 'support for indexing common datatypes in GiST';
CREATE EXTENSION IF NOT EXISTS citext;
COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';
-- CREATE EXTENSION IF NOT EXISTS fuzzystrmatch;
-- COMMENT ON EXTENSION fuzzystrmatch IS 'determine similarities and distance between strings';
-- CREATE EXTENSION IF NOT EXISTS postgis;
-- COMMENT ON EXTENSION postgis IS 'PostGIS geometry, geography, and raster spatial types and functions';
-- CREATE EXTENSION IF NOT EXISTS postgis_tiger_geocoder WITH SCHEMA tiger;
-- COMMENT ON EXTENSION postgis_tiger_geocoder IS 'PostGIS tiger geocoder and reverse geocoder';
-- CREATE EXTENSION IF NOT EXISTS postgis_topology WITH SCHEMA topology;
-- COMMENT ON EXTENSION postgis_topology IS 'PostGIS topology spatial types and functions';
-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';
-- CREATE EXTENSION IF NOT EXISTS citext;
-- COMMENT ON EXTENSION "citext" IS 'enabled case insensitivity for text like email addresses';
-- +goose StatementBegin
DO $$ BEGIN CREATE DOMAIN day_of_week AS smallint NOT NULL DEFAULT 0 CONSTRAINT day_of_week_check CHECK (
    (
        (VALUE >= 0)
        AND (VALUE <= 6)
    )
);
EXCEPTION
WHEN duplicate_object THEN NULL;
END $$;
-- +goose StatementEnd
COMMENT ON DOMAIN day_of_week IS '
    day_of_week is a physical day of a week represented by a number such that 
        0=SUNDAY
        1=MONDAY
        2=TUESDAY
        3=WEDNESDAY
        4=THURSDAY
        5=FRIDAY
        6=SATURDAY
';
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION empty (text) RETURNS boolean LANGUAGE sql IMMUTABLE AS $_$
SELECT $1 ~ '^[[:space:]]*$';
$_$;
-- +goose StatementEnd
COMMENT ON FUNCTION empty (text) IS 'Find empty strings or strings containing only whitespace';
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
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_min_bigint_value () RETURNS bigint AS $$
DECLARE min_bigint_value BIGINT;
BEGIN
SELECT (2 ^(8 * pg_column_size(1::bigint) -2))::bigint << 1 INTO min_bigint_value;
RETURN min_bigint_value;
END $$ LANGUAGE "plpgsql";
-- +goose StatementEnd
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
-- +goose StatementBegin
CALL build_h_tally_table ();
-- +goose StatementEnd
--
--
-- +goose StatementBegin
CREATE TYPE money AS (number NUMERIC, currency_code CHAR(3));
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TYPE address_kind AS ENUM ('billing', 'shipping', 'mailing');
-- +goose StatementEnd

-- +goose StatementBegin
CREATE EXTENSION pgcrypto;
-- +goose StatementEnd
--
--
-- +goose Down
DROP EXTENSION IF EXISTS citext;
DROP EXTENSION IF EXISTS pgcrypto;
DROP FUNCTION IF EXISTS smallint;
DROP FUNCTION IF EXISTS day_of_week;
DROP FUNCTION IF EXISTS empty;
DROP FUNCTION IF EXISTS get_max_bigint_value;
DROP FUNCTION IF EXISTS get_min_bigint_value;
DROP TYPE IF EXISTS address_kind;
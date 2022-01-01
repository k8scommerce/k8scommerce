-- +goose Up
--
--
--
-- +goose StatementBegin
CREATE OR REPLACE PROCEDURE update_category_nested_set_hierarchy ()
LANGUAGE plpgsql
AS $$
DECLARE
    rec RECORD;
    total int;
BEGIN
    SELECT
        count(*) INTO total
    FROM
        category_node_hierarchy;
    FOR rec IN
    SELECT
        CAST(SUBSTRING(h.sort_path, t.n, 32) AS int) AS id,
        COUNT(*) AS node_count --Includes current node
    FROM
        category_node_hierarchy h,
        h_tally t
    WHERE
        t.n BETWEEN 1
        AND LENGTH(sort_path)
    GROUP BY
        SUBSTRING(h.sort_path, t.n, 32) LOOP
            UPDATE
                category_node_hierarchy
            SET
                node_count = rec.node_count,
                rgt = (rec.node_count - 1) * 2 + lft + 1
            WHERE
                id = rec.id;
        END LOOP;
    --
    --
    UPDATE
        category t
    SET
        lft = tn.lft,
        rgt = tn.rgt,
        depth = tn.h_level
    FROM
        category_node_hierarchy tn
    WHERE
        t.id = tn.id;
    --
    --
    DROP TABLE IF EXISTS category_node_hierarchy;
END
$$;

-- +goose StatementEnd
--
--
--
-- +goose StatementBegin
CREATE OR REPLACE PROCEDURE rebuild_category_nested_set (storeID int)
LANGUAGE plpgsql
AS $$
BEGIN
    DROP TABLE IF EXISTS category_node_hierarchy;
    CREATE TABLE category_node_hierarchy (
        id int PRIMARY KEY,
        parent_id int,
        h_level int,
        lft int,
        rgt int, --Place holder
        node_number int,
        node_count int, --Place holder
        sort_path VARBIT
    );
            WITH RECURSIVE cteBuildPath AS (
                SELECT
                    anchor.id,
                    anchor.parent_id,
                    1 AS h_level,
                    CAST(CAST(anchor.id AS bit(32)) AS VARBIT) AS sort_path
                FROM
                    category AS anchor
                WHERE
                    parent_id IS NULL
                    AND store_id = storeID
                UNION ALL
                SELECT
                    recur.id,
                    recur.parent_id,
                    cte.h_level + 1 AS h_level,
                    CAST( --This does the concatenation to build SortPath
                        cte.sort_path || CAST(recur.id AS bit(32)) AS VARBIT) AS sort_path
                FROM
                    category AS recur
                    INNER JOIN cteBuildPath AS cte ON cte.id = recur.parent_id
                WHERE
                    recur.store_id = storeID) INSERT INTO category_node_hierarchy
            SELECT
                coalesce(sorted.id, 0) AS id,
                sorted.parent_id,
                coalesce(sorted.h_level, 0) AS h_level,
                coalesce(CAST(0 AS int), 0) AS lft, --Place holder
                coalesce(CAST(0 AS int), 0) AS rgt, --Place holder
                ROW_NUMBER() OVER (ORDER BY sorted.sort_path) AS node_number,
                coalesce(CAST(0 AS int), 0) AS node_count, --Place holder
                coalesce(sorted.sort_path, sorted.sort_path) AS sort_path
            FROM
                cteBuildPath AS sorted;
        -- CREATE UNIQUE INDEX uk_node_hierarchy_id ON node_hierarchy (id);
        -- CREATE UNIQUE INDEX uk_node_hierarchy_sort_path ON node_hierarchy (sort_path);
        CREATE INDEX uk_node_hierarchy_lft ON category_node_hierarchy (lft);
        UPDATE
            category_node_hierarchy
        SET
            lft = 2 * node_number - h_level;
        CALL update_category_nested_set_hierarchy ();
END
$$;

-- +goose StatementEnd
-- +goose Down
-- DROP PROCEDURE IF EXISTS build_h_tally_table;
-- DROP PROCEDURE IF EXISTS update_category_nested_set_hierarchy;
-- DROP PROCEDURE IF EXISTS rebuild_category_nested_set;

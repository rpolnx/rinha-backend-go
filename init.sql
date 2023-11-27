\ c rinha;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE EXTENSION IF NOT EXISTS "pg_trgm";

CREATE TABLE people (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    username varchar(32),
    name varchar(100),
    birthday date,
    stack text []
);

CREATE
OR REPLACE FUNCTION immutable_array_to_string(text [], text) RETURNS text as $ $
SELECT
    array_to_string($ 1, $ 2);
$ $ LANGUAGE sql IMMUTABLE;

-- DROP INDEX idx_term_search IF EXISTS;
-- CREATE INDEX idx_term_search ON people USING GIN (
--     to_tsvector(
--         'english',
--         coalesce(name, '') || ' ' || coalesce(username, '') || ' ' || immutable_array_to_string(coalesce(stack, '{}'), ' ')
--     )
-- );

CREATE INDEX idx_term_search ON people USING GIN (
    (
        coalesce(name, '') || ' ' || coalesce(username, '') || ' ' || immutable_array_to_string(coalesce(stack, '{}'), ' ')
    ) gin_trgm_ops
);

-- explain analyze
-- SELECT
--     *
-- FROM
--     "people"
-- WHERE
--     'java' < % (
--         coalesce(name, '') || ' ' || coalesce(username, '') || ' ' || immutable_array_to_string(coalesce(stack, '{}'), ' ')
--     )
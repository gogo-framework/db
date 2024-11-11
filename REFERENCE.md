# SQL Operations by Database Type

## Comparison Operators
Common across all DBs:
- `=` (Equal)
- `<>` or `!=` (Not Equal)
- `<` (Less Than)
- `>` (Greater Than)
- `<=` (Less Than or Equal)
- `>=` (Greater Than or Equal)
- `BETWEEN`
- `IN`
- `NOT IN`
- `LIKE`
- `IS NULL`
- `IS NOT NULL`

PostgreSQL specific:
- `ILIKE` (Case-insensitive LIKE)
- `@@` (Text search match)
- `<->` (Distance operator)
- `~` (Matches regex, case-sensitive)
- `~*` (Matches regex, case-insensitive)
- `!~` (Does not match regex, case-sensitive)
- `!~*` (Does not match regex, case-insensitive)

MySQL specific:
- `REGEXP` or `RLIKE`
- `NOT REGEXP`
- `<=>` (NULL-safe equal)
- `SOUNDS LIKE`

## Logical Operators
Common across all DBs:
- `AND`
- `OR`
- `NOT`
- `XOR` (MySQL explicit support)

## Arithmetic Operators
Common across all DBs:
- `+` (Addition)
- `-` (Subtraction)
- `*` (Multiplication)
- `/` (Division)
- `%` or `MOD` (Modulo)

PostgreSQL specific:
- `^` (Exponentiation)
- `|/` (Square root)
- `||/` (Cube root)
- `!` (Factorial)
- `@` (Absolute value)

## Bitwise Operators
Common across all DBs:
- `&` (AND)
- `|` (OR)
- `~` (NOT)
- `<<` (Left shift)
- `>>` (Right shift)
- `^` (XOR)

## Aggregate Functions
Common across all DBs:
- `COUNT`
- `SUM`
- `AVG`
- `MAX`
- `MIN`
- `GROUP_CONCAT` (MySQL) / `STRING_AGG` (PostgreSQL) / `GROUP_CONCAT` (SQLite)

PostgreSQL specific:
- `ARRAY_AGG`
- `JSON_AGG`
- `JSONB_AGG`
- `BOOL_AND`
- `BOOL_OR`
- `EVERY`

MySQL specific:
- `BIT_AND`
- `BIT_OR`
- `BIT_XOR`
- `STD`
- `STDDEV`
- `VARIANCE`

## Window Functions
Common across PostgreSQL and MySQL (SQLite has limited support):
- `ROW_NUMBER`
- `RANK`
- `DENSE_RANK`
- `FIRST_VALUE`
- `LAST_VALUE`
- `LAG`
- `LEAD`
- `NTH_VALUE`

PostgreSQL specific:
- `PERCENT_RANK`
- `CUME_DIST`
- `NTILE`

## Join Types
Common across all DBs:
- `INNER JOIN`
- `LEFT JOIN`
- `RIGHT JOIN` (Not in SQLite)
- `FULL OUTER JOIN` (PostgreSQL only)
- `CROSS JOIN`

PostgreSQL specific:
- `LATERAL JOIN`

## Set Operations
Common across all DBs:
- `UNION`
- `UNION ALL`
- `INTERSECT` (Limited in MySQL)
- `EXCEPT` (PostgreSQL) / `MINUS` (MySQL)

## Data Modification
Common across all DBs:
- `INSERT`
- `UPDATE`
- `DELETE`
- `TRUNCATE` (Limited in SQLite)

PostgreSQL specific:
- `UPSERT` (`INSERT ... ON CONFLICT`)
- `RETURNING`

MySQL specific:
- `REPLACE`
- `INSERT ... ON DUPLICATE KEY UPDATE`

## JSON Operations
PostgreSQL:
- `->` (Get JSON array element)
- `->>` (Get JSON array element as text)
- `#>` (Get JSON path)
- `#>>` (Get JSON path as text)
- `@>` (Contains)
- `<@` (Contained by)
- `?` (Key exists)
- `?|` (Any key exists)
- `?&` (All keys exist)

MySQL:
- `->` (Extract value)
- `->>` (Extract value as text)
- `JSON_EXTRACT`
- `JSON_CONTAINS`
- `JSON_CONTAINS_PATH`
- `JSON_ARRAY`
- `JSON_OBJECT`

SQLite:
- `->` (Extract value)
- `->>` (Extract value as text)
- `json_extract`
- `json_array`
- `json_object`


## Table Operations
Common across all DBs:
- `CREATE TABLE`
- `DROP TABLE`
- `ALTER TABLE`
  - `ADD COLUMN`
  - `DROP COLUMN`
  - `RENAME COLUMN`
  - `RENAME TABLE`
  - `ALTER COLUMN` (type modifications)

PostgreSQL specific:
- `ALTER TABLE ... INHERIT`
- `ALTER TABLE ... NO INHERIT`
- `ALTER TABLE ... OF type`
- `ALTER TABLE ... NOT OF`
- `CLUSTER`

MySQL specific:
- `ALTER TABLE ... ENGINE`
- `ALTER TABLE ... CHARACTER SET`
- `ALTER TABLE ... COLLATE`
- `OPTIMIZE TABLE`
- `REPAIR TABLE`
- `ANALYZE TABLE`

## Index Operations
Common across all DBs:
- `CREATE INDEX`
- `DROP INDEX`
- `CREATE UNIQUE INDEX`

PostgreSQL specific:
- `CREATE INDEX CONCURRENTLY`
- `DROP INDEX CONCURRENTLY`
- `REINDEX`
- Index types:
  - B-tree (default)
  - Hash
  - GiST
  - SP-GiST
  - GIN
  - BRIN

MySQL specific:
- `ALTER TABLE ... ADD INDEX`
- `ALTER TABLE ... ADD SPATIAL INDEX`
- `ALTER TABLE ... ADD FULLTEXT INDEX`
- Index types:
  - B-tree (default)
  - Hash
  - FULLTEXT
  - SPATIAL

SQLite specific:
- Index types:
  - B-tree (default)
  - Virtual table indexes (FTS, R-tree)

## Constraint Operations
Common across all DBs:
- `PRIMARY KEY`
- `FOREIGN KEY`
- `UNIQUE`
- `CHECK`
- `NOT NULL`
- `DEFAULT`

PostgreSQL specific:
- `EXCLUDE`
- `DEFERRABLE`
- `INITIALLY DEFERRED`
- `NO INHERIT`

MySQL specific:
- `SPATIAL INDEX`
- `FULLTEXT INDEX`

## Schema Operations
Common across PostgreSQL and MySQL (SQLite is single-schema):
- `CREATE SCHEMA`
- `DROP SCHEMA`
- `ALTER SCHEMA`
- `RENAME SCHEMA` (PostgreSQL only)

## View Operations
Common across all DBs:
- `CREATE VIEW`
- `DROP VIEW`
- `ALTER VIEW`
- `REPLACE VIEW`

PostgreSQL specific:
- `CREATE MATERIALIZED VIEW`
- `REFRESH MATERIALIZED VIEW`
- `CREATE RECURSIVE VIEW`

## Trigger Operations
Common across all DBs:
- `CREATE TRIGGER`
- `DROP TRIGGER`
- Trigger timing:
  - `BEFORE`
  - `AFTER`
  - `INSTEAD OF` (PostgreSQL and SQLite)

PostgreSQL specific:
- Statement-level triggers
- Constraint triggers
- Event triggers

## Sequence Operations (PostgreSQL mainly, some MySQL support)
- `CREATE SEQUENCE`
- `ALTER SEQUENCE`
- `DROP SEQUENCE`
- `NEXTVAL`
- `CURRVAL`
- `SETVAL`

## Partitioning Operations
PostgreSQL:
- `CREATE TABLE ... PARTITION BY`
  - RANGE
  - LIST
  - HASH
- `ATTACH PARTITION`
- `DETACH PARTITION`

MySQL:
- `CREATE TABLE ... PARTITION BY`
  - RANGE
  - LIST
  - HASH
  - KEY
- `ALTER TABLE ... ADD PARTITION`
- `ALTER TABLE ... DROP PARTITION`
- `ALTER TABLE ... REORGANIZE PARTITION`

## Table Space Operations (PostgreSQL and MySQL)
- `CREATE TABLESPACE`
- `ALTER TABLESPACE`
- `DROP TABLESPACE`

## Role/User Operations
PostgreSQL:
- `CREATE ROLE`
- `ALTER ROLE`
- `DROP ROLE`
- `GRANT`
- `REVOKE`

MySQL:
- `CREATE USER`
- `ALTER USER`
- `DROP USER`
- `GRANT`
- `REVOKE`

## Comment Operations
PostgreSQL:
- `COMMENT ON`

MySQL:
- Comments in creation: `CREATE TABLE ... COMMENT '...'`
- `ALTER TABLE ... COMMENT '...'`

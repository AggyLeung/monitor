# Database Design

## PostgreSQL (System of Record)

- Dynamic modeling:
  - `ci_type`
  - `ci_type_attribute`
- CI core:
  - `ci`
  - `ci_attribute_value` (EAV using JSONB)
- Fallback graph relation:
  - `relation`
- Audit and task:
  - `audit_log`
  - `sync_task`
- Compensation queue:
  - `graph_sync_failed`

Migration file: `go-core/migrations/001_init.sql`

## Neo4j (Graph Query Plane)

Node label:
- `:CI`

Node properties:
- `id`, `name`, `type`, `status`

Relationship types:
- `RUNS_ON`
- `DEPENDS_ON`
- `CONNECTS_TO`
- `CONTAINS`

Migration scripts:
- `go-core/migrations/neo4j/001_indexes.cypher`
- `go-core/migrations/neo4j/002_constraints.cypher`

## Consistency Strategy

Write path:
1. Write `ci` and `ci_attribute_value` in PostgreSQL transaction.
2. Publish graph sync job asynchronously.
3. If Neo4j update fails, persist to `graph_sync_failed`.

Read path:
1. Static attributes from PostgreSQL.
2. Topology from Neo4j.
3. If Neo4j is unavailable, degrade to PostgreSQL `relation` table.

Compensation:
- Failed graph sync records are queryable and retryable via:
  - `GET /api/v1/sync/failed`
  - `POST /api/v1/sync/failed/:id/retry`

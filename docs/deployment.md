# Deployment

## Runtime Model

- `go-core`: API gateway and data-plane service
- `python-discovery`: scheduled/triggered discovery worker
- `python-workflow`: workflow and ticket/report automation
- `frontend`: static SPA served by Nginx

## Environment Variables

Recommended secret-managed variables:

- `POSTGRES_DSN`
- `REDIS_ADDR`, `REDIS_PASSWORD`
- `NEO4J_URI`, `NEO4J_USER`, `NEO4J_PASSWORD`
- `JWT_SECRET`
- `CMDB_API_KEY`

In Kubernetes:

- Put non-sensitive settings in `ConfigMap`.
- Put credentials and keys in `Secret`.

## Observability

- Go: expose Prometheus metrics endpoint (recommended `/metrics`).
- Python: structured JSON logs for Loki ingestion.

## Gateway

- Use Nginx/Traefik for:
  - TLS termination
  - path-based routing
  - rate limiting
  - auth forwarding

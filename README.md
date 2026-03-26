# Enterprise Network Monitoring System (MVP)

This repository contains a runnable MVP for an enterprise network monitoring system with a closed-loop flow:

`monitoring -> detection -> analysis -> remediation -> verification`

## Repository Layout

- `docs/`: Architecture and operating model documents
- `schemas/`: Unified data schemas
- `services/collector-gateway/`: Push telemetry ingress
- `services/detection-engine/`: Rule-based anomaly detection
- `services/incident-center/`: Incident lifecycle orchestration
- `services/auto-remediation-runner/`: Automated runbook execution
- `services/verification-service/`: Post-fix validation
- `docker-compose.yml`: Local one-command startup

## Quick Start

1. Start all services:

```bash
docker compose up --build
```

2. Push a sample telemetry point:

```bash
curl -X POST http://localhost:8001/telemetry \
  -H "Content-Type: application/json" \
  -d '{
    "device_id":"edge-router-01",
    "metric_name":"cpu_utilization",
    "value":92.5,
    "timestamp":"2026-03-26T13:00:00Z",
    "tags":{"site":"shanghai-a","service":"wan-gateway"}
  }'
```

3. Query incidents:

```bash
curl http://localhost:8003/incidents
```

## Why This Matches Your Five Frameworks

1. Event lifecycle: incident states (`open -> remediating -> verifying -> resolved/investigating`) and timeline records.
2. Observability pillars: metric ingestion implemented, log/trace/profiling reserved via schema and architecture docs.
3. Zero-intrusive perspective: collector designed for push telemetry; eBPF integration points documented.
4. AIOps collaboration: detection engine supports rules now, with model hooks in service boundaries.
5. Closed loop: incident center automatically triggers remediation and verification, then writes back outcomes.

## Next Expansion (Recommended)

1. Add Kafka/NATS for decoupled ingestion and replay.
2. Add TSDB/log/trace backends (VictoriaMetrics + Loki + Tempo).
3. Add topology graph and causal reasoning service.
4. Add NLP retrieval for similar incident recommendation.

## New Microservice Split (Go + Python)

This repository now also includes a split architecture:

- `go-core/`: Go Core API (Gin + GORM + Neo4j + JWT + Casbin + Redis Streams)
- `python-discovery/`: Discovery service (FastAPI + boto3/netmiko + Redis worker)
- `python-workflow/`: Workflow service (FastAPI, change/report process entry)
- `frontend/`: React + TypeScript + Ant Design + Redux Toolkit + Axios + ECharts + G6
- `docker-compose.microservices.yml`: full stack with PostgreSQL, Neo4j, Redis
- `docs/database-design.md`: PostgreSQL + Neo4j model and consistency strategy
- `docs/deployment.md`: deployment and runtime operations notes

### Start Split Architecture

```bash
docker compose -f docker-compose.microservices.yml up --build
```

Frontend routes included:

- `/resources`: CI list/filter/search/add/edit/soft-delete
- `/resources/:id`: CI detail + relation preview + change history
- `/topology/:id`: fullscreen topology with drill and export
- `/ci-types`: CI type and attribute template management
- `/tasks`: sync tasks, status and logs
- `/users`: user and role assignment

Frontend implementation notes:

- Dynamic CI attribute form is rendered from CI type templates.
- Axios client uses request/response interceptors for JWT and 401 redirect.
- Vite dev server proxies `/api` to Go Core API.
- Production frontend is built as static assets and served by Nginx.

Backend data notes:

- PostgreSQL schema is migration-managed in `go-core/migrations/001_init.sql`.
- Neo4j indexes/constraints are migration-managed in `go-core/migrations/neo4j/`.
- Graph sync compensation queue is stored in `graph_sync_failed`.

## Local Development

Dependency services:

```bash
docker compose -f docker-compose.dev.yml up -d
```

Run services locally:

1. Go Core API

```bash
cd go-core
go run cmd/server/main.go
```

2. Python Discovery

```bash
cd python-discovery
python -m venv .venv
# Windows
.venv\Scripts\activate
# Linux/macOS
source .venv/bin/activate
pip install -r requirements.txt
python main.py
```

3. Frontend

```bash
cd frontend
npm install
npm run dev
```

## Deployment Notes

- Containerize all services and orchestrate with Kubernetes or Docker Compose.
- Manage secrets by ConfigMap/Secret (DB passwords, API keys).
- Expose Prometheus metrics from Go services.
- Use structured logs in Python and send to Loki.
- Use Nginx or Traefik as gateway for SSL termination and routing.

## CI/CD Pipeline

1. Push code and trigger unit tests (Go/Python/Frontend).
2. Build and push Docker images.
3. Deploy to staging and run integration tests.
4. Roll out production with rolling update strategy.

CD workflow file:

- `.github/workflows/cd.yml`

Branch strategy:

- `develop` -> staging deploy
- `main`/`master` -> production deploy

Required repository/environment secrets:

- `STAGING_DEPLOY_WEBHOOK`
- `PROD_DEPLOY_WEBHOOK`

Recommended GitHub Environments:

- `staging` (optional approval)
- `production` (required approval)

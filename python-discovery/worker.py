from __future__ import annotations

import json
import os
import time
from datetime import datetime, timezone

import redis

from client.cmdb_client import CMDBClient
from discovery.network import scan_devices


def run_worker() -> None:
    redis_addr = os.getenv("REDIS_ADDR", "redis")
    redis_port = int(os.getenv("REDIS_PORT", "6379"))
    stream = os.getenv("REDIS_STREAM", "cmdb_tasks")
    base_url = os.getenv("CMDB_BASE_URL", "http://go-core:8080")
    api_key = os.getenv("CMDB_API_KEY", "replace-me")

    r = redis.Redis(host=redis_addr, port=redis_port, decode_responses=True)
    client = CMDBClient(base_url, api_key)
    last_id = "0-0"

    while True:
        records = r.xread({stream: last_id}, count=10, block=5000)
        if not records:
            continue

        for _, messages in records:
            for message_id, fields in messages:
                last_id = message_id
                raw = fields.get("data", "{}")
                task = json.loads(raw)
                payload = task.get("payload", {})
                scope = payload.get("scope", "default")
                hosts = ["10.10.1.1", "10.10.1.2"] if scope == "all" else ["10.10.1.1"]
                scanned = scan_devices(client, hosts)
                client.post_sync_callback(
                    {
                        "task_id": task.get("id"),
                        "status": "success",
                        "scanned": scanned,
                        "finished_at": datetime.now(timezone.utc).isoformat(),
                    }
                )
        time.sleep(0.1)


if __name__ == "__main__":
    run_worker()

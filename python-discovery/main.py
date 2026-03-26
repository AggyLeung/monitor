from __future__ import annotations

import os
from datetime import datetime, timezone

from fastapi import FastAPI

from client.cmdb_client import CMDBClient
from discovery.network import scan_devices

app = FastAPI(title="python-discovery", version="0.1.0")


def build_client() -> CMDBClient:
    base_url = os.getenv("CMDB_BASE_URL", "http://go-core:8080")
    api_key = os.getenv("CMDB_API_KEY", "replace-me")
    return CMDBClient(base_url, api_key)


@app.get("/health")
def health() -> dict:
    return {"status": "ok", "service": "python-discovery"}


@app.post("/scan/network")
def scan_network() -> dict:
    hosts = os.getenv("DISCOVERY_HOSTS", "10.10.1.1,10.10.1.2").split(",")
    count = scan_devices(build_client(), hosts)
    return {"scanned": count, "at": datetime.now(timezone.utc).isoformat()}

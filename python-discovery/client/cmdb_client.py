from __future__ import annotations

from typing import Any, Dict

import requests


class CMDBClient:
    def __init__(self, base_url: str, api_key: str) -> None:
        self.base_url = base_url.rstrip("/")
        self.session = requests.Session()
        self.session.headers.update({"Authorization": f"Bearer {api_key}"})

    def upsert_ci(self, ci_data: Dict[str, Any]) -> Dict[str, Any]:
        ci_id = ci_data.get("id")
        if not ci_id:
            raise ValueError("ci_data.id is required")
        resp = self.session.put(f"{self.base_url}/api/v1/cis/{ci_id}", json=ci_data, timeout=10)
        if resp.status_code == 404:
            resp = self.session.post(f"{self.base_url}/api/v1/cis", json=ci_data, timeout=10)
        resp.raise_for_status()
        return resp.json()

    def post_sync_callback(self, payload: Dict[str, Any]) -> Dict[str, Any]:
        resp = self.session.post(f"{self.base_url}/api/v1/sync/callback", json=payload, timeout=10)
        resp.raise_for_status()
        return resp.json()

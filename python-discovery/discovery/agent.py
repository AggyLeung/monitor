from __future__ import annotations

from typing import Dict

from client.cmdb_client import CMDBClient


def sync_agent_payload(client: CMDBClient, payload: Dict) -> Dict:
    return client.upsert_ci(payload)

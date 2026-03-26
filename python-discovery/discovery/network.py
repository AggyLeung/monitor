from __future__ import annotations

from typing import Iterable

from client.cmdb_client import CMDBClient


def scan_devices(client: CMDBClient, hosts: Iterable[str]) -> int:
    # Placeholder scanner; replace with netmiko/snmp walk in next iteration.
    total = 0
    for host in hosts:
        payload = {
            "id": f"net-{host.replace('.', '-')}",
            "name": host,
            "type_id": 2,
            "status": "active",
            "attributes": {"ip": host, "vendor": "unknown"},
        }
        client.upsert_ci(payload)
        total += 1
    return total

from __future__ import annotations

from typing import Any, Dict, List

import boto3

from client.cmdb_client import CMDBClient


def _first_tag(tags: List[Dict[str, Any]]) -> str | None:
    if not tags:
        return None
    return tags[0].get("Value")


def sync_ec2(client: CMDBClient) -> int:
    ec2 = boto3.client("ec2")
    instances = ec2.describe_instances()
    count = 0
    for reservation in instances.get("Reservations", []):
        for instance in reservation.get("Instances", []):
            ci_data = {
                "id": instance["InstanceId"],
                "name": _first_tag(instance.get("Tags", [])) or instance["InstanceId"],
                "type_id": 1,
                "status": "active",
                "attributes": {
                    "cpu": instance.get("CpuOptions", {}).get("CoreCount"),
                    "memory": instance.get("MemoryInfo", {}).get("SizeInMiB"),
                    "ip": instance.get("PublicIpAddress"),
                    "state": instance.get("State", {}).get("Name"),
                },
            }
            client.upsert_ci(ci_data)
            count += 1
    return count

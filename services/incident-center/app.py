import os
import uuid
from datetime import datetime, timezone
from typing import Any, Dict, List

import requests
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel, Field

app = FastAPI(title="incident-center", version="0.1.0")

REMEDIATION_URL = os.getenv("REMEDIATION_URL", "http://127.0.0.1:8004")
VERIFICATION_URL = os.getenv("VERIFICATION_URL", "http://127.0.0.1:8005")

INCIDENTS: Dict[str, Dict[str, Any]] = {}


class Alert(BaseModel):
    alert_id: str
    metric_name: str
    severity: str
    observed_value: float
    threshold: float
    device_id: str
    timestamp: datetime
    tags: Dict[str, str] = Field(default_factory=dict)


def now_utc() -> str:
    return datetime.now(timezone.utc).isoformat()


def add_timeline(incident: Dict[str, Any], event: str, detail: str = "") -> None:
    incident["timeline"].append({"at": now_utc(), "event": event, "detail": detail})
    incident["updated_at"] = now_utc()


@app.get("/health")
def health() -> dict:
    return {"status": "ok", "service": "incident-center", "count": len(INCIDENTS)}


@app.get("/incidents")
def list_incidents() -> List[Dict[str, Any]]:
    return list(INCIDENTS.values())


@app.get("/incidents/{incident_id}")
def get_incident(incident_id: str) -> Dict[str, Any]:
    if incident_id not in INCIDENTS:
        raise HTTPException(status_code=404, detail="incident not found")
    return INCIDENTS[incident_id]


@app.post("/incidents")
def create_incident(alert: Alert) -> Dict[str, Any]:
    incident_id = str(uuid.uuid4())
    incident = {
        "incident_id": incident_id,
        "status": "open",
        "created_at": now_utc(),
        "updated_at": now_utc(),
        "alert": alert.model_dump(mode="json"),
        "timeline": [],
    }
    add_timeline(incident, "incident_opened", f"from alert {alert.alert_id}")
    INCIDENTS[incident_id] = incident

    incident["status"] = "remediating"
    add_timeline(incident, "remediation_started")

    remediation_result = {"executed": False}
    try:
        rem_resp = requests.post(
            f"{REMEDIATION_URL}/runbook/execute",
            json={"incident": incident},
            timeout=3,
        )
        rem_resp.raise_for_status()
        remediation_result = rem_resp.json()
        add_timeline(incident, "remediation_finished", remediation_result.get("action", ""))
    except Exception as exc:
        add_timeline(incident, "remediation_failed", str(exc))

    incident["status"] = "verifying"
    add_timeline(incident, "verification_started")

    verification_result = {"verified": False}
    try:
        ver_resp = requests.post(
            f"{VERIFICATION_URL}/verify",
            json={"incident": incident, "remediation_result": remediation_result},
            timeout=3,
        )
        ver_resp.raise_for_status()
        verification_result = ver_resp.json()
    except Exception as exc:
        verification_result = {"verified": False, "reason": str(exc)}

    if verification_result.get("verified"):
        incident["status"] = "resolved"
        add_timeline(incident, "incident_resolved", verification_result.get("reason", ""))
    else:
        incident["status"] = "investigating"
        add_timeline(incident, "incident_needs_investigation", verification_result.get("reason", ""))

    return {
        "created": True,
        "incident_id": incident_id,
        "status": incident["status"],
        "verification_result": verification_result,
    }

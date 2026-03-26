import os
import uuid
from datetime import datetime, timezone
from typing import Dict, Optional

import requests
from fastapi import FastAPI
from pydantic import BaseModel, Field

app = FastAPI(title="detection-engine", version="0.1.0")
INCIDENT_CENTER_URL = os.getenv("INCIDENT_CENTER_URL", "http://127.0.0.1:8003")

# Rule + model hybrid entry point: rules are active now, model_score is reserved.
THRESHOLDS = {
    "cpu_utilization": 85.0,
    "latency_ms": 1000.0,
    "packet_loss_pct": 5.0,
}


class TelemetryPoint(BaseModel):
    device_id: str
    metric_name: str
    value: float
    timestamp: datetime
    received_at: Optional[datetime] = None
    tags: Optional[Dict[str, str]] = Field(default_factory=dict)


def severity_for(metric_name: str, value: float, threshold: float) -> str:
    ratio = value / threshold if threshold > 0 else 0
    if ratio >= 1.6:
        return "critical"
    if ratio >= 1.3:
        return "high"
    if ratio >= 1.1:
        return "medium"
    return "low"


@app.get("/health")
def health() -> dict:
    return {"status": "ok", "service": "detection-engine"}


@app.post("/analyze")
def analyze(point: TelemetryPoint) -> dict:
    threshold = THRESHOLDS.get(point.metric_name, float("inf"))
    is_anomaly = point.value > threshold
    model_score = 0.0

    if not is_anomaly:
        return {"anomaly": False, "model_score": model_score}

    alert = {
        "alert_id": str(uuid.uuid4()),
        "metric_name": point.metric_name,
        "severity": severity_for(point.metric_name, point.value, threshold),
        "observed_value": point.value,
        "threshold": threshold,
        "device_id": point.device_id,
        "timestamp": datetime.now(timezone.utc).isoformat(),
        "tags": point.tags,
    }

    incident_result = {"created": False}
    try:
        resp = requests.post(f"{INCIDENT_CENTER_URL}/incidents", json=alert, timeout=3)
        resp.raise_for_status()
        incident_result = resp.json()
    except Exception as exc:
        incident_result = {"created": False, "error": str(exc)}

    return {"anomaly": True, "alert": alert, "incident": incident_result, "model_score": model_score}

import os
from datetime import datetime, timezone
from typing import Dict, Optional

import requests
from fastapi import FastAPI
from pydantic import BaseModel, Field

app = FastAPI(title="collector-gateway", version="0.1.0")

DETECTION_ENGINE_URL = os.getenv("DETECTION_ENGINE_URL", "http://127.0.0.1:8002")
INGESTED_COUNT = 0


class TelemetryPoint(BaseModel):
    device_id: str
    metric_name: str
    value: float
    timestamp: datetime
    tags: Optional[Dict[str, str]] = Field(default_factory=dict)


@app.get("/health")
def health() -> dict:
    return {"status": "ok", "service": "collector-gateway", "ingested_count": INGESTED_COUNT}


@app.post("/telemetry")
def ingest_telemetry(point: TelemetryPoint) -> dict:
    global INGESTED_COUNT
    INGESTED_COUNT += 1

    payload = point.model_dump(mode="json")
    payload["received_at"] = datetime.now(timezone.utc).isoformat()

    detection_result = {}
    try:
        response = requests.post(
            f"{DETECTION_ENGINE_URL}/analyze",
            json=payload,
            timeout=2,
        )
        response.raise_for_status()
        detection_result = response.json()
    except Exception as exc:
        detection_result = {"anomaly": False, "error": str(exc)}

    return {
        "accepted": True,
        "telemetry": payload,
        "detection_result": detection_result,
    }

from typing import Any, Dict

from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI(title="verification-service", version="0.1.0")

HARD_LIMITS = {
    "cpu_utilization": 98.0,
    "latency_ms": 2500.0,
    "packet_loss_pct": 20.0,
}


class VerificationRequest(BaseModel):
    incident: Dict[str, Any]
    remediation_result: Dict[str, Any]


@app.get("/health")
def health() -> dict:
    return {"status": "ok", "service": "verification-service"}


@app.post("/verify")
def verify(req: VerificationRequest) -> dict:
    incident = req.incident
    alert = incident.get("alert", {})
    metric_name = alert.get("metric_name", "unknown")
    observed_value = float(alert.get("observed_value", 0))
    hard_limit = HARD_LIMITS.get(metric_name, float("inf"))
    runbook_ok = bool(req.remediation_result.get("executed"))
    metric_within_guardrail = observed_value < hard_limit

    verified = runbook_ok and metric_within_guardrail
    reason = "verification passed" if verified else "requires manual investigation"
    return {
        "verified": verified,
        "reason": reason,
        "guardrail": {"metric_name": metric_name, "hard_limit": hard_limit, "observed_value": observed_value},
    }

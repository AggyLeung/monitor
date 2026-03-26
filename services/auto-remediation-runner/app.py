from datetime import datetime, timezone
from typing import Any, Dict

from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI(title="auto-remediation-runner", version="0.1.0")


class RunbookRequest(BaseModel):
    incident: Dict[str, Any]


@app.get("/health")
def health() -> dict:
    return {"status": "ok", "service": "auto-remediation-runner"}


@app.post("/runbook/execute")
def execute_runbook(req: RunbookRequest) -> dict:
    incident = req.incident
    alert = incident.get("alert", {})
    metric_name = alert.get("metric_name", "unknown")

    if metric_name == "cpu_utilization":
        action = "scale_out_or_restart_hot_process"
    elif metric_name == "latency_ms":
        action = "route_shift_and_qos_tuning"
    elif metric_name == "packet_loss_pct":
        action = "interface_reset_and_path_switch"
    else:
        action = "generic_safe_recovery"

    return {
        "executed": True,
        "action": action,
        "at": datetime.now(timezone.utc).isoformat(),
        "operator": "automation",
    }

from __future__ import annotations

from datetime import datetime, timezone
from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI(title="python-workflow", version="0.1.0")


class WorkflowJob(BaseModel):
    ticket_id: str
    action: str
    payload: dict = {}


@app.get("/health")
def health() -> dict:
    return {"status": "ok", "service": "python-workflow"}


@app.post("/workflow/execute")
def execute(job: WorkflowJob) -> dict:
    return {
        "accepted": True,
        "ticket_id": job.ticket_id,
        "action": job.action,
        "executed_at": datetime.now(timezone.utc).isoformat(),
    }

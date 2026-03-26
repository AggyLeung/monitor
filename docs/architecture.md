# Architecture (MVP)

## Core Services

1. Collector Gateway
- Accepts push telemetry from devices/agents.
- Normalizes payload and forwards to detection.

2. Detection Engine
- Runs lightweight rule checks for anomaly pre-detection.
- Creates alert objects and sends them to incident center.

3. Incident Center
- Maintains incident state machine.
- Triggers remediation and verification automatically.
- Stores timeline for post-incident review.

4. Auto Remediation Runner
- Executes predefined runbooks.
- Returns execution record for audit.

5. Verification Service
- Validates post-remediation health conditions.
- Sends binary result for incident closure.

## Data Flow

1. Device pushes telemetry to collector.
2. Collector forwards normalized event to detection.
3. Detection generates alert if anomaly is found.
4. Incident center opens incident and triggers runbook.
5. Verification checks recovery result.
6. Incident status transitions to resolved or investigating.

## Mapping to Target Vision

- Event lifecycle: state machine + timeline + postmortem hooks.
- Observability: metric path implemented; logs/traces/profiling reserved.
- Zero-intrusive: push-native design; eBPF hooks planned.
- AIOps: rule engine now, model endpoints can be added without service rewrite.
- Automation loop: full path wired in MVP.

# Event Lifecycle

## Lifecycle Stages

1. Pre-incident (Prevention)
- Validate change windows and policy constraints.
- Block unsafe config changes before rollout.

2. In-incident (Detection and Fast Recovery)
- Detect anomalies through rules/models.
- Correlate alerts and trigger rapid remediation.

3. Post-incident (Operations and Learning)
- Persist timeline and runbook outcomes.
- Feed postmortem corpus for similar-case recommendations.

## Incident State Machine (Current MVP)

1. `open`
2. `remediating`
3. `verifying`
4. `resolved` or `investigating`

## Fields to Keep for Future COE

- incident_id
- symptom metric/value
- suspected root cause
- remediation actions
- verification evidence
- postmortem summary

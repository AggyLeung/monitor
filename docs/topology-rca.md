# Topology and RCA Evolution

## Current MVP

- Incident center accepts single-alert incidents.
- Correlation uses metric-level context only.

## Next Step

1. Build topology graph entities:
- device
- service
- link
- zone
- dependency

2. Add causal propagation:
- upstream/downstream impact propagation
- suppress child alarms when parent fault is confirmed

3. RCA ranking:
- combine topology distance, temporal order, and metric deviation score

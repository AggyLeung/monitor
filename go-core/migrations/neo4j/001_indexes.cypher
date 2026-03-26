CREATE INDEX ci_id IF NOT EXISTS FOR (c:CI) ON (c.id);
CREATE INDEX ci_type IF NOT EXISTS FOR (c:CI) ON (c.type);
CREATE INDEX ci_status IF NOT EXISTS FOR (c:CI) ON (c.status);

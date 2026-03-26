CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS ci_type (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) UNIQUE NOT NULL,
  label VARCHAR(100),
  icon VARCHAR(50),
  created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS ci_type_attribute (
  id SERIAL PRIMARY KEY,
  ci_type_id INT NOT NULL REFERENCES ci_type(id) ON DELETE CASCADE,
  name VARCHAR(100) NOT NULL,
  label VARCHAR(100),
  data_type VARCHAR(50) NOT NULL,
  is_required BOOLEAN DEFAULT false,
  enum_options JSONB,
  sort_order INT
);

CREATE TABLE IF NOT EXISTS ci (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  ci_type_id INT REFERENCES ci_type(id),
  name VARCHAR(255) NOT NULL,
  status VARCHAR(50) DEFAULT 'active',
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW(),
  created_by UUID,
  updated_by UUID
);

CREATE TABLE IF NOT EXISTS ci_attribute_value (
  ci_id UUID NOT NULL REFERENCES ci(id) ON DELETE CASCADE,
  attribute_id INT NOT NULL REFERENCES ci_type_attribute(id) ON DELETE CASCADE,
  value JSONB,
  PRIMARY KEY (ci_id, attribute_id)
);

CREATE TABLE IF NOT EXISTS relation (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  source_ci_id UUID REFERENCES ci(id),
  target_ci_id UUID REFERENCES ci(id),
  type VARCHAR(50),
  properties JSONB,
  created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS audit_log (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID,
  action VARCHAR(50),
  ci_id UUID,
  old_value JSONB,
  new_value JSONB,
  created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS sync_task (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  task_type VARCHAR(50),
  status VARCHAR(20),
  params JSONB,
  result JSONB,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  completed_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS graph_sync_failed (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  entity_type VARCHAR(50) NOT NULL,
  entity_id UUID,
  payload JSONB NOT NULL,
  error_message TEXT,
  status VARCHAR(20) NOT NULL DEFAULT 'pending',
  retry_count INT NOT NULL DEFAULT 0,
  next_retry_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_ci_type_name ON ci_type(name);
CREATE INDEX IF NOT EXISTS idx_ci_ci_type_id ON ci(ci_type_id);
CREATE INDEX IF NOT EXISTS idx_ci_status ON ci(status);
CREATE INDEX IF NOT EXISTS idx_relation_source_ci_id ON relation(source_ci_id);
CREATE INDEX IF NOT EXISTS idx_relation_target_ci_id ON relation(target_ci_id);
CREATE INDEX IF NOT EXISTS idx_audit_log_ci_id ON audit_log(ci_id);
CREATE INDEX IF NOT EXISTS idx_sync_task_status ON sync_task(status);
CREATE INDEX IF NOT EXISTS idx_graph_sync_failed_status ON graph_sync_failed(status);

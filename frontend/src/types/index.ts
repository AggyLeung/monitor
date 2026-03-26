export type ResourceStatus = "active" | "deleted";

export interface Resource {
  id: string;
  name: string;
  type: string;
  owner: string;
  ip: string;
  site: string;
  status: ResourceStatus;
  attributes: Record<string, string | number>;
  updatedAt: string;
}

export interface Relation {
  source: string;
  target: string;
  label?: string;
}

export interface ChangeRecord {
  id: string;
  field: string;
  before: string;
  after: string;
  operator: string;
  at: string;
}

export interface CiType {
  id: string;
  name: string;
  attributes: CiAttributeSchema[];
  enabled: boolean;
}

export type CiAttributeType = "string" | "int" | "enum";

export interface CiAttributeSchema {
  name: string;
  label: string;
  type: CiAttributeType;
  options?: string[];
}

export interface SyncTask {
  id: string;
  name: string;
  status: "running" | "success" | "failed" | "queued";
  startedAt: string;
  durationSec: number;
  log: string;
}

export interface UserItem {
  id: string;
  username: string;
  role: "admin" | "operator" | "viewer";
  email: string;
  active: boolean;
}

import { createSlice, nanoid, PayloadAction } from "@reduxjs/toolkit";
import { Resource } from "../../types";

interface ResourceFormInput {
  name: string;
  type: string;
  owner: string;
  ip: string;
  site: string;
  attributes?: Record<string, string | number>;
}

interface ResourcesState {
  items: Resource[];
}

const initialState: ResourcesState = {
  items: [
    {
      id: "ci-001",
      name: "edge-router-01",
      type: "NetworkDevice",
      owner: "NOC-A",
      ip: "10.10.1.1",
      site: "Shanghai-A",
      status: "active",
      attributes: {
        vendor: "Cisco",
        model: "ASR-1001",
        os: "IOS-XE"
      },
      updatedAt: "2026-03-26T15:20:00Z"
    },
    {
      id: "ci-002",
      name: "wan-fw-01",
      type: "Firewall",
      owner: "SecOps",
      ip: "10.10.1.10",
      site: "Shanghai-A",
      status: "active",
      attributes: {
        vendor: "Palo Alto",
        model: "PA-3220",
        policySet: "corp-main"
      },
      updatedAt: "2026-03-26T14:02:00Z"
    }
  ]
};

const resourcesSlice = createSlice({
  name: "resources",
  initialState,
  reducers: {
    addResource(state, action: PayloadAction<ResourceFormInput>) {
      state.items.unshift({
        id: `ci-${nanoid(6)}`,
        status: "active",
        attributes: action.payload.attributes ?? {},
        updatedAt: new Date().toISOString(),
        ...action.payload
      });
    },
    updateResource(state, action: PayloadAction<{ id: string; patch: ResourceFormInput }>) {
      const item = state.items.find((r) => r.id === action.payload.id);
      if (!item) {
        return;
      }
      Object.assign(item, action.payload.patch);
      item.attributes = action.payload.patch.attributes ?? item.attributes;
      item.updatedAt = new Date().toISOString();
    },
    softDeleteResource(state, action: PayloadAction<string>) {
      const item = state.items.find((r) => r.id === action.payload);
      if (!item) {
        return;
      }
      item.status = "deleted";
      item.updatedAt = new Date().toISOString();
    }
  }
});

export const { addResource, updateResource, softDeleteResource } = resourcesSlice.actions;
export default resourcesSlice.reducer;

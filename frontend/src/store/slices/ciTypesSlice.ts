import { createSlice, nanoid, PayloadAction } from "@reduxjs/toolkit";
import { CiAttributeSchema, CiType } from "../../types";

interface CiTypesState {
  items: CiType[];
}

const initialState: CiTypesState = {
  items: [
    {
      id: "type-server",
      name: "Server",
      enabled: true,
      attributes: [
        { name: "cpu", label: "CPU Cores", type: "int" },
        { name: "memory", label: "Memory (GB)", type: "int" },
        { name: "ip", label: "IP Address", type: "string" },
        { name: "os", label: "Operating System", type: "string" }
      ]
    },
    {
      id: "type-network-device",
      name: "NetworkDevice",
      enabled: true,
      attributes: [
        { name: "vendor", label: "Vendor", type: "string" },
        { name: "model", label: "Model", type: "string" },
        { name: "os", label: "OS", type: "string" }
      ]
    },
    {
      id: "type-firewall",
      name: "Firewall",
      enabled: true,
      attributes: [
        { name: "throughput", label: "Throughput (Gbps)", type: "int" },
        { name: "policy_set", label: "Policy Set", type: "string" },
        { name: "mode", label: "Mode", type: "enum", options: ["L3", "L7"] }
      ]
    }
  ]
};

const ciTypesSlice = createSlice({
  name: "ciTypes",
  initialState,
  reducers: {
    addCiType(state, action: PayloadAction<{ name: string; attributes: CiAttributeSchema[] }>) {
      state.items.unshift({
        id: `type-${nanoid(6)}`,
        name: action.payload.name,
        attributes: action.payload.attributes,
        enabled: true
      });
    },
    toggleCiType(state, action: PayloadAction<{ id: string; enabled: boolean }>) {
      const item = state.items.find((it) => it.id === action.payload.id);
      if (!item) {
        return;
      }
      item.enabled = action.payload.enabled;
    }
  }
});

export const { addCiType, toggleCiType } = ciTypesSlice.actions;
export default ciTypesSlice.reducer;

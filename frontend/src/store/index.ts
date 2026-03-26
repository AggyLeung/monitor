import { configureStore } from "@reduxjs/toolkit";
import resourcesReducer from "./slices/resourcesSlice";
import ciTypesReducer from "./slices/ciTypesSlice";

export const store = configureStore({
  reducer: {
    resources: resourcesReducer,
    ciTypes: ciTypesReducer
  }
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

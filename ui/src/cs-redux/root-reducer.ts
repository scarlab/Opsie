import { combineReducers } from "@reduxjs/toolkit";
import GlobalSlice from "./slices/global.slice";
import AuthSlice from "./slices/auth.slice";

const CsRootReducer = combineReducers({
    global: GlobalSlice.reducer,
    auth: AuthSlice.reducer,
})

export type CsRootState = ReturnType<typeof CsRootReducer>;
export default CsRootReducer;
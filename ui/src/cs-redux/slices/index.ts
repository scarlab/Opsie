import TeamSlice from "./team.slice";
import { combineReducers } from "@reduxjs/toolkit";
import GlobalSlice from "./global.slice";
import AuthSlice from "./auth.slice";
import UserSlice from "./user.slice";

const CsRootReducer = combineReducers({
    team: TeamSlice.reducer,
    global: GlobalSlice.reducer,
    auth: AuthSlice.reducer,
    user: UserSlice.reducer,
})

export type CsRootState = ReturnType<typeof CsRootReducer>;
export default CsRootReducer;
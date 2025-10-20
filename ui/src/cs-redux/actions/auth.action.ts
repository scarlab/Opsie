import ApiManager from "@/configs/api.config";
import type { LoginPayload } from "@/types/auth";
import { createAsyncThunk } from "@reduxjs/toolkit";

export class AuthAction {

    Login = createAsyncThunk(
        "Auth/Login",
        async (payload: LoginPayload, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.post(`/auth/login`, payload);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )


    Session = createAsyncThunk(
        "Auth/Session",
        async (_, { rejectWithValue, fulfillWithValue }) => {
            try {

                const { data } = await ApiManager.get(`/auth/session`);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )


    Logout = createAsyncThunk(
        "Auth/Logout",
        async (_, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.get(`/auth/logout`);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )
}
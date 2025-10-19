import ApiManager from "@/configs/api.config";
import type { LoginPayload } from "@/types/auth";
import { createAsyncThunk } from "@reduxjs/toolkit";

export class AuthAction {

    login = createAsyncThunk(
        "auth/login",
        async (payload: LoginPayload, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.post(`/auth/login`, payload);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )


    session = createAsyncThunk(
        "auth/session",
        async (_, { rejectWithValue, fulfillWithValue }) => {
            try {

                const { data } = await ApiManager.get(`/auth/session`);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )


    logout = createAsyncThunk(
        "auth/logout",
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
import ApiManager from "@/configs/api.config";
import type { NewOwnerPayload } from "@/types/user";
import { createAsyncThunk } from "@reduxjs/toolkit";

export class UserAction {

    CreateOwnerAccount = createAsyncThunk(
        "User/CreateOwnerAccount",
        async (payload: NewOwnerPayload, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.post(`/user/owner/create`, payload);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )

    GetOwnerCount = createAsyncThunk(
        "User/GetOwnerCount",
        async (_, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.get(`/user/owner/count`);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )


    // User Account
    UpdateAccountDisplayName = createAsyncThunk(
        "User/UpdateAccountName",
        async (payload: { display_name: string }, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.patch(`/user/account/update/name`, payload);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )

    UpdateAccountPassword = createAsyncThunk(
        "User/UpdateAccountPassword",
        async (payload: { password: string; new_password: string }, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.patch(`/user/account/update/password`, payload);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )
}
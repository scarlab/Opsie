import ApiManager from "@/configs/api.config";
import type { NewOwnerPayload } from "@/types/user";
import { createAsyncThunk } from "@reduxjs/toolkit";

export class UserAction {

    createOwnerAccount = createAsyncThunk(
        "user/createOwnerAccount",
        async (payload: NewOwnerPayload, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.post(`/user/owner/create`, payload);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )

    getOwnerCount = createAsyncThunk(
        "user/getOwnerCount",
        async (_, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.get(`/user/owner/count`);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )
}
// Created: 2025/10/20 15:33:22

import ApiManager from "@/configs/api.config";
import { createAsyncThunk } from "@reduxjs/toolkit";

export class TeamAction {
    Example = createAsyncThunk(
        "Team/Example",
        async (_, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.get(`/team/login`);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )
}
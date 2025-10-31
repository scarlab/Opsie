// Created: 2025/10/20 15:33:22

import ApiManager from "@/configs/api.config";
import { createAsyncThunk } from "@reduxjs/toolkit";

export class TeamAction {
    /// _______________________________________________________________________________________________
    /// User ------------------------------------------------------------------------------------------
    /// ---

    GetAllTeamOfUser = createAsyncThunk(
        "Team/GetAllTeamOfUser",
        async (_, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.get(`/team/user/get/all`);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )

    GetDefaultTeamOfUser = createAsyncThunk(
        "Team/GetDefaultTeamOfUser",
        async (_, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.get(`/team/user/get/default`);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )

    SetDefaultTeamOfUser = createAsyncThunk(
        "Team/SetDefaultTeamOfUser",
        async (payload: { id: number }, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.get(`/team/user/set/default/${payload.id}`);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )
}
// Created: 2025/10/20 15:33:22

import ApiManager from "@/configs/api.config";
import type { NewTeamPayload } from "@/types/team.type";
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
                console.log(data);
                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )

    SetDefaultTeamOfUser = createAsyncThunk(
        "Team/SetDefaultTeamOfUser",
        async (payload: { id: string }, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.get(`/team/user/set/default/${payload.id}`);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )

    /// _______________________________________________________________________________________________
    /// Protected Section [Auth, Admin] ---------------------------------------------------------------
    /// ---

    GetTeamMembers = createAsyncThunk(
        "Team/GetTeamMembers",
        async (payload: { id: string }, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.get(`/team/get/members/${payload.id}`);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )

    CreateTeam = createAsyncThunk(
        "Team/CreateTeam",
        async (payload: NewTeamPayload, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.post(`/team/create`, payload);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )
}
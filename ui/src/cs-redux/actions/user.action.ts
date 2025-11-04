import ApiManager from "@/configs/api.config";
import type { AddUserToTeamPayload, RemoveUserToTeamPayload } from "@/types/user-team.type";
import type { NewOwnerPayload, NewUserPayload, UpdateUserPayload } from "@/types/user.type";
import { createAsyncThunk } from "@reduxjs/toolkit";

export class UserAction {

    /// _______________________________________________________________________________________________
    /// Onboarding ------------------------------------------------------------------------------------
    /// ---

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



    /// _______________________________________________________________________________________________
    /// User Account ----------------------------------------------------------------------------------
    /// ---

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


    /// _______________________________________________________________________________________________
    /// Protected Section [Auth, Admin] ---------------------------------------------------------------
    /// ---

    CreateUser = createAsyncThunk(
        "User/CreateUser",
        async (payload: NewUserPayload, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.post(`/user/create`, payload);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )
    GetAllUser = createAsyncThunk(
        "User/GetAllUser",
        async (_, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.get(`/user/get`);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )

    GetUserById = createAsyncThunk(
        "User/GetUserById",
        async (payload: { id: string }, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.get(`/user/get/${payload.id}`);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )

    UpdateUser = createAsyncThunk(
        "User/UpdateUser",
        async (payload: { id: string; data: UpdateUserPayload }, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.patch(`/user/update/${payload.id}`, payload.data);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )


    DeleteUser = createAsyncThunk(
        "User/DeleteUser",
        async (payload: { id: string; }, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.delete(`/user/delete/${payload.id}`);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )

    AddUserToTeam = createAsyncThunk(
        "User/AddUserToTeam",
        async (payload: AddUserToTeamPayload, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.post(`/user/team/add`, payload);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )

    RemoveUserToTeam = createAsyncThunk(
        "User/RemoveUserToTeam",
        async (payload: RemoveUserToTeamPayload, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.post(`/user/team/remove`, payload);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )

}
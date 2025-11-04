import { createSlice } from '@reduxjs/toolkit';
import type { UserModel } from '@/types/user.type';
import { UserAction } from '../actions/user.action';

const UserModel = new UserAction()

const initialState: {
    user: UserModel | undefined;
    users: UserModel[];

    loading: boolean;
    notFound: boolean;
} = {
    user: undefined,
    loading: false,
    notFound: false,
    users: []
};

const UserSlice = createSlice({
    name: "UserSlice",
    initialState,
    reducers: {
        restoreUserUser: (state, { payload }) => {
            state.user = payload
        },
        removeUser: (state, _) => {
            state.user = undefined
        },
    },
    extraReducers: (builder) => {

        /// _______________________________________________________________________________________________
        /// Protected Section [Auth, Admin] ---------------------------------------------------------------
        /// ---

        builder
            .addCase(UserModel.CreateUser.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(UserModel.CreateUser.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.users.push(payload.user);
            })
            .addCase(UserModel.CreateUser.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
                state.user = undefined;
            })


        builder
            .addCase(UserModel.GetAllUser.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(UserModel.GetAllUser.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.users = payload.users;
            })
            .addCase(UserModel.GetAllUser.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
                state.users = [];
            })

        builder
            .addCase(UserModel.GetUserById.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(UserModel.GetUserById.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.user = payload.user;
            })
            .addCase(UserModel.GetUserById.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
                state.user = undefined;
            })

        builder
            .addCase(UserModel.UpdateUser.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(UserModel.UpdateUser.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.user = payload.user;
            })
            .addCase(UserModel.UpdateUser.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
                state.user = undefined;
            })
    }
});



export default UserSlice;







import { createSlice } from '@reduxjs/toolkit';
import type { User } from '@/types/user.type';
import { UserAction } from '../actions/user.action';

const User = new UserAction()

const initialState: {
    user: User | undefined;
    users: User[];

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
        }
    },
    extraReducers: (builder) => {

        /// _______________________________________________________________________________________________
        /// Protected Section [Auth, Admin] ---------------------------------------------------------------
        /// ---

        builder
            .addCase(User.CreateUser.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(User.CreateUser.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.user = payload.user;

                state.users.unshift(payload.user);
            })
            .addCase(User.CreateUser.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
                state.user = undefined;
            })


        builder
            .addCase(User.GetAllUser.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(User.GetAllUser.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.users = payload.users;
            })
            .addCase(User.GetAllUser.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
                state.users = [];
            })

        builder
            .addCase(User.GetUserById.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(User.GetUserById.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.user = payload.user;
            })
            .addCase(User.GetUserById.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
                state.user = undefined;
            })

        builder
            .addCase(User.UpdateUser.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(User.UpdateUser.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.user = payload.user;
            })
            .addCase(User.UpdateUser.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
                state.user = undefined;
            })
    }
});



export default UserSlice;







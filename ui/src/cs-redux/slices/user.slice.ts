import { createSlice } from '@reduxjs/toolkit';
import type { User } from '@/types/user';
import { UserAction } from '../actions/user.action';

const User = new UserAction()

const initialState: {
    user: User | undefined;

    loading: boolean;
    notFound: boolean;
} = {
    user: undefined,
    loading: false,
    notFound: false,
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
        builder
            .addCase(User.createOwnerAccount.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(User.createOwnerAccount.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.user = payload.user;
            })
            .addCase(User.createOwnerAccount.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
                state.user = undefined;
            })
    }
});



export default UserSlice;







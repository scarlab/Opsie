import { createSlice } from '@reduxjs/toolkit';
import { AuthAction } from '../actions/auth.action';
import type { AuthUser } from '@/types/auth';
import { removeLocalAuthUser, setLocalAuthUser } from '@/helpers/auth.helper';

const Auth = new AuthAction()

const initialState: {
    authUser: AuthUser | undefined;

    loading: boolean;
    notFound: boolean;
} = {
    authUser: undefined,
    loading: false,
    notFound: false,
};

const AuthSlice = createSlice({
    name: "AuthSlice",
    initialState,
    reducers: {
        restoreAuthUser: (state, { payload }) => {
            state.authUser = payload
        }
    },
    extraReducers: (builder) => {
        builder
            .addCase(Auth.login.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Auth.login.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.authUser = payload.auth_user;

                //...
                setLocalAuthUser(payload.auth_user)
            })
            .addCase(Auth.login.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
                state.authUser = undefined;
            })

        builder
            .addCase(Auth.session.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Auth.session.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.authUser = payload.auth_user;

                //...
                setLocalAuthUser(payload.auth_user)
            })
            .addCase(Auth.session.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
                state.authUser = undefined;

                // ...
                removeLocalAuthUser()
            })


        builder
            .addCase(Auth.logout.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Auth.logout.fulfilled, (state, _) => {
                state.loading = false;
                state.authUser = undefined;

                // ...
                removeLocalAuthUser()
            })
            .addCase(Auth.logout.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
            })
    }
});



export default AuthSlice;







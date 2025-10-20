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
        },
        updateAuthUser: (state, { payload }) => {
            state.authUser = payload
        },
    },
    extraReducers: (builder) => {
        builder
            .addCase(Auth.Login.pending, () => {
                // state.loading = true;
            })
            .addCase(Auth.Login.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.authUser = payload.auth_user;

                //...
                setLocalAuthUser(payload.auth_user)
            })
            .addCase(Auth.Login.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
                state.authUser = undefined;
            })

        builder
            .addCase(Auth.Session.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Auth.Session.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.authUser = payload.auth_user;

                //...
                setLocalAuthUser(payload.auth_user)
            })
            .addCase(Auth.Session.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
                state.authUser = undefined;

                // ...
                removeLocalAuthUser()
            })


        builder
            .addCase(Auth.Logout.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Auth.Logout.fulfilled, (state, _) => {
                state.loading = false;
                state.authUser = undefined;

                // ...
                removeLocalAuthUser()
            })
            .addCase(Auth.Logout.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
            })
    }
});



export default AuthSlice;







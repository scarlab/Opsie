import { createSlice } from '@reduxjs/toolkit';
import { AuthAction } from '../actions/auth.action';

const Auth = new AuthAction()

const initialState: {
    sessionUser: undefined;

    loading: boolean;
    notFound: boolean;
} = {
    sessionUser: undefined,
    loading: false,
    notFound: false,
};

const AuthSlice = createSlice({
    name: "AuthSlice",
    initialState,
    reducers: {

    },
    extraReducers: (builder) => {
        builder
            .addCase(Auth.login.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Auth.login.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.sessionUser = payload.session_user;
            })
            .addCase(Auth.login.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
                state.sessionUser = undefined;
            })

        builder
            .addCase(Auth.session.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Auth.session.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.sessionUser = payload.session_user;
            })
            .addCase(Auth.session.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
                state.sessionUser = undefined;
            })


        builder
            .addCase(Auth.logout.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Auth.logout.fulfilled, (state, _) => {
                state.loading = false;
                state.sessionUser = undefined;
            })
            .addCase(Auth.logout.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
            })
    }
});



export default AuthSlice;







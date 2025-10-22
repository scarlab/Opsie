// Created: 2025/10/20 15:33:22

import { createSlice } from '@reduxjs/toolkit';
import { TeamAction } from '../actions/team.action';

const Team = new TeamAction()

const initialState: {
    name: string;
    loading: boolean;
    notFound: boolean;
} = {
    name: '',
    loading: false,
    notFound: false,
};

const TeamSlice = createSlice({
    name: "TeamSlice",
    initialState,
    reducers: {
        setName: (state, { payload }) => {
            state.name = payload;
        },
    },
    extraReducers: (builder) => {
        builder
            .addCase(Team.Example.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Team.Example.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.name = payload.name;
            })
            .addCase(Team.Example.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
            })
    }
});



export default TeamSlice;







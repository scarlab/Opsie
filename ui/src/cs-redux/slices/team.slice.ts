// Created: 2025/10/20 15:33:22

import { createSlice } from '@reduxjs/toolkit';
import { TeamAction } from '../actions/team.action';
import type { TeamModel } from '@/types/team.type';
import { setLocalTeam } from '@/helpers/team.helper';

const Team = new TeamAction()

const initialState: {
    team: TeamModel | undefined;
    teams: TeamModel[] | undefined;
    user_default_team: TeamModel | undefined;
    user_teams: TeamModel[] | undefined;
    loading: boolean;
    notFound: boolean;
} = {
    loading: false,
    notFound: false,
    team: undefined,
    teams: undefined,
    user_default_team: undefined,
    user_teams: undefined
};

const TeamSlice = createSlice({
    name: "TeamSlice",
    initialState,
    reducers: {
        setName: (state, { payload }) => {
            state.team = payload;
        },
    },
    extraReducers: (builder) => {
        builder
            .addCase(Team.GetAllTeamOfUser.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Team.GetAllTeamOfUser.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.user_teams = payload.teams;
            })
            .addCase(Team.GetAllTeamOfUser.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
            })


        builder
            .addCase(Team.GetDefaultTeamOfUser.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Team.GetDefaultTeamOfUser.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.user_default_team = payload.team;

                // ---
                setLocalTeam(payload.team);
            })
            .addCase(Team.GetDefaultTeamOfUser.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
            })

        builder
            .addCase(Team.SetDefaultTeamOfUser.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Team.SetDefaultTeamOfUser.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.user_default_team = payload.team;

                // ---
                setLocalTeam(payload.team);
            })
            .addCase(Team.SetDefaultTeamOfUser.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
            })
    }
});



export default TeamSlice;







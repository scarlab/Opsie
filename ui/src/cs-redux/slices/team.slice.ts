// Created: 2025/10/20 15:33:22

import { createSlice } from '@reduxjs/toolkit';
import { TeamAction } from '../actions/team.action';
import type { TeamModel, UserTeam } from '@/types/team.type';
import { setLocalTeam } from '@/helpers/team.helper';
import type { TeamMember } from '@/types/user-team.type';

const Team = new TeamAction()

const initialState: {
    team: TeamModel | undefined;
    teams: TeamModel[] | undefined;
    user_default_team: UserTeam | undefined;
    user_teams: UserTeam[] | undefined;
    team_members: TeamMember[] | undefined;
    loading: boolean;
    notFound: boolean;
} = {
    loading: false,
    notFound: false,
    team: undefined,
    teams: undefined,
    user_default_team: undefined,
    user_teams: undefined,
    team_members: undefined
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
        /// _______________________________________________________________________________________________
        /// Auth User -------------------------------------------------------------------------------------
        /// ---

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



        /// _______________________________________________________________________________________________
        /// Protected Section [Auth, Admin] ---------------------------------------------------------------
        /// ---

        builder
            .addCase(Team.GetTeamMembers.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Team.GetTeamMembers.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.team_members = payload.members;

                // ---
                setLocalTeam(payload.team);
            })
            .addCase(Team.GetTeamMembers.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
            })

        builder
            .addCase(Team.CreateTeam.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Team.CreateTeam.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.teams?.unshift(payload.team);
            })
            .addCase(Team.CreateTeam.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
            })

    }
});



export default TeamSlice;







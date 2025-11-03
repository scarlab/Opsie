// Created: 2025/10/20 15:33:22

import { createSlice } from '@reduxjs/toolkit';
import { TeamAction } from '../actions/team.action';
import type { TeamModel, UserTeam } from '@/types/team.type';
import { setLocalTeam } from '@/helpers/team.helper';
import type { TeamMember } from '@/types/user-team.type';

const Team = new TeamAction()

const initialState: {
    // Auth User
    auth_default_team: UserTeam | undefined;
    auth_teams: UserTeam[] | undefined;

    // Admin
    team: TeamModel | undefined;
    teams: TeamModel[] | undefined;
    user_teams: UserTeam[] | undefined;
    team_members: TeamMember[] | undefined;

    // Utils
    loading: boolean;
    notFound: boolean;
} = {
    loading: false,
    notFound: false,
    team: undefined,
    teams: undefined,
    auth_default_team: undefined,
    auth_teams: undefined,
    team_members: undefined,
    user_teams: undefined
};

const TeamSlice = createSlice({
    name: "TeamSlice",
    initialState,
    reducers: {
        removeTeamMembers: (state, _) => {
            state.team_members = undefined;
        },
        removeTeam: (state, _) => {
            state.team = undefined;
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
                state.auth_teams = payload.teams;
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
                state.auth_default_team = payload.team;

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
                state.auth_default_team = payload.team;

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


        builder
            .addCase(Team.GetAllTeams.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Team.GetAllTeams.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.teams = payload.teams;
            })
            .addCase(Team.GetAllTeams.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
            })


        builder
            .addCase(Team.GetTeamById.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Team.GetTeamById.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.team = payload.team;
            })
            .addCase(Team.GetTeamById.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
            })


        builder
            .addCase(Team.GetTeamMembers.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Team.GetTeamMembers.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.team_members = payload.members;
            })
            .addCase(Team.GetTeamMembers.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
            })


        builder
            .addCase(Team.GetAllTeamsOfUser.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Team.GetAllTeamsOfUser.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.user_teams = payload.teams;
            })
            .addCase(Team.GetAllTeamsOfUser.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
            })


        builder
            .addCase(Team.Update.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Team.Update.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.team = payload.team;
            })
            .addCase(Team.Update.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
            })


        builder
            .addCase(Team.Delete.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Team.Delete.fulfilled, (state, _) => {
                state.loading = false;
                state.team = undefined;
            })
            .addCase(Team.Delete.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
            })


    }
});



export default TeamSlice;







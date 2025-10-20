// Created: 2025/10/20 15:33:22

import { createSlice } from '@reduxjs/toolkit';
import { OrganizationAction } from '../actions/organization.action';

const Organization = new OrganizationAction()

const initialState: {
    name: string;
    loading: boolean;
    notFound: boolean;
} = {
    name: '',
    loading: false,
    notFound: false,
};

const OrganizationSlice = createSlice({
    name: "OrganizationSlice",
    initialState,
    reducers: {
        setName: (state, { payload }) => {
            state.name = payload;
        },
    },
    extraReducers: (builder) => {
        builder
            .addCase(Organization.Example.pending, (state, _) => {
                state.loading = true;
            })
            .addCase(Organization.Example.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.name = payload.name;
            })
            .addCase(Organization.Example.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
            })
    }
});



export default OrganizationSlice;







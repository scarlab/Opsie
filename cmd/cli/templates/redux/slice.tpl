// Created: {{.CreatedAt}}

import { createSlice } from '@reduxjs/toolkit';
import { {{.Name}}Action } from '../actions/{{.PackageName}}.action';

const {{.Name}} = new {{.Name}}Action()

const initialState: {
    name: string;
    loading: boolean;
    notFound: boolean;
} = {
    name: '',
    loading: false,
    notFound: false,
};

const {{.Name}}Slice = createSlice({
    name: "{{.Name}}Slice",
    initialState,
    reducers: {
        setName: (state, { payload }) => {
            state.name = payload;
        },
    },
    extraReducers: (builder) => {
        builder
            .addCase({{.Name}}.Example.pending, (state, _) => {
                state.loading = true;
            })
            .addCase({{.Name}}.Example.fulfilled, (state, { payload }) => {
                state.loading = false;
                state.name = payload.name;
            })
            .addCase({{.Name}}.Example.rejected, (state, _) => {
                state.loading = false;
                state.notFound = true;
            })
    }
});



export default {{.Name}}Slice;







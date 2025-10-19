// Created: {{.CreatedAt}}

import ApiManager from "@/configs/api.config";
import { createAsyncThunk } from "@reduxjs/toolkit";

export class {{.Name}}Action {
    Example = createAsyncThunk(
        "{{.Name}}/Example",
        async (_, { rejectWithValue, fulfillWithValue }) => {
            try {
                const { data } = await ApiManager.get(`/{{.PackageName}}/login`);

                return fulfillWithValue(data)
            } catch (error: any) {
                return rejectWithValue(error.response.data);
            }
        }
    )
}
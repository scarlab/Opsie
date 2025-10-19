import { createSlice } from '@reduxjs/toolkit';



const initialState: {
    query: string;
    loading: boolean;
    notFound: boolean;
} = {
    query: '',
    loading: false,
    notFound: false,
};

const GlobalSlice = createSlice({
    name: "GlobalSlice",
    initialState,
    reducers: {
        setSearchQuery: (state, { payload }) => {
            state.query = payload;
        },
    },
    extraReducers: (builder) => {
    }
});



export default GlobalSlice;







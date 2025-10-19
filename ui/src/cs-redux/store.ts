import { configureStore } from '@reduxjs/toolkit'
import CsRootReducer from './slices'
import { type TypedUseSelectorHook, useDispatch, useSelector } from 'react-redux'


export const CsCreateStore = () => {
    return configureStore({
        reducer: CsRootReducer,
    })
}

// Infer the store type from createStore()
type _Store = ReturnType<typeof CsCreateStore>;
type _Dispatch = _Store['dispatch'];
type _RootState = ReturnType<typeof CsRootReducer>;


// Hooks 
export const useCsDispatch = () => useDispatch<_Dispatch>();
export const useCsSelector: TypedUseSelectorHook<_RootState> = useSelector;
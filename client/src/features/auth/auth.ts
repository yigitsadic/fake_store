import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { RootState } from '../../app/store';
import {User} from "./user";

export interface AuthState {
    user?: User,

    loggedIn: boolean;
}

const initialState: AuthState = {
    user: {},
    loggedIn: false,
};

export const authSlice = createSlice({
    name: "auth-handler",
    initialState,
    reducers: {
        login: (state, action: PayloadAction<User>) => {
            state.loggedIn = true;
            state.user = action.payload;
        },
        logout: (state) => {
            state.loggedIn = false;
            state.user = {};
        },
    },
});

export const { login, logout } = authSlice.actions;
export const selectUser = (state: RootState) => state.auth.user;
export const selectLoggedIn = (state: RootState) => state.auth.loggedIn;

export default authSlice.reducer;

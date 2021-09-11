import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import {User} from "./user";
import {RootState} from "../store";

export interface AuthState {
    user: User,
}

const initialState: AuthState = {
    user: {
        loggedIn: false,
    },
};

export const authSlice = createSlice({
    name: "auth-handler",
    initialState,
    reducers: {
        login: (state, action: PayloadAction<User>) => {
            state.user = action.payload;
        },
        logout: (state) => {
            localStorage.removeItem("fake_store_token");
            state.user = { loggedIn: false };
        },
    },
});

export const { login, logout } = authSlice.actions;
export const selectedCurrentUser = (state: RootState) => state.auth.user;

export default authSlice.reducer;

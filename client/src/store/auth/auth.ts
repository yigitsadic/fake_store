import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import {User} from "./user";
import {RootState} from "../store";

export interface AppState {
    user: User;
    cartItemCount: number;
}

const initialState: AppState = {
    user: {
        loggedIn: false,
    },
    cartItemCount: 0,
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
        updateCartCount: (state, action: PayloadAction<number>) => {
            state.cartItemCount = action.payload;
        },
    },
});

export const { login, logout, updateCartCount } = authSlice.actions;
export const selectedCurrentUser = (state: RootState) => state.auth.user;
export const selectCartCount = (state: RootState) => state.auth.cartItemCount;

export default authSlice.reducer;

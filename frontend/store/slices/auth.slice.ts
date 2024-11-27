import { createSlice } from '@reduxjs/toolkit';

interface AuthState {
  token: string | null;
  isAuthenticated: boolean;
  role: string | null;
  email: string | null;
}

const initialState: AuthState = {
  token: null,
  isAuthenticated: false,
  role: null,
  email: null,
};

const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    login(state, action) {
      state.token = action.payload.token;
      state.isAuthenticated = action.payload.isAuthenticated;
      state.role = action.payload.role;
      state.email = action.payload.email;
    },
    logout(state) {
      state.token = null;
      state.isAuthenticated = false;
      state.role = null;
      state.email = null;
    },
    setRole(state, action) {
      state.role = action.payload;
    },
    setEmail(state, action) {
      state.email = action.payload;
    },
  },
});

export const { login, logout, setRole, setEmail } = authSlice.actions;
export default authSlice.reducer;

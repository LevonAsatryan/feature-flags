import { createSlice, PayloadAction } from '@reduxjs/toolkit';

export type LoginState = {
  isLoggedIn: boolean;
};

const initialState: LoginState = {
  isLoggedIn: true,
};

export const loginSlice = createSlice({
  name: 'login',
  initialState,
  reducers: {
    setIsLoggedIn: (state, action: PayloadAction<boolean>) => {
      state.isLoggedIn = action.payload;
    },
  },
});

export const { setIsLoggedIn } = loginSlice.actions;

export default loginSlice.reducer;

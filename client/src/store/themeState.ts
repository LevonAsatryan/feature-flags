import { createSlice } from '@reduxjs/toolkit';

export type ThemeState = {
  theme: 'light' | 'dark';
};

const initialState: ThemeState = {
  theme: 'dark',
};

type ThemeColors = {
  '--ff-primary-color': string;
  '--ff-text-color-primary': string;
  '--ff-background-color-primary': string;
};

const lightThemeColors = new Map<string, string>([
  ['--ff-primary-color', '#4096ff'],
  ['--ff-text-color-primary', '#131212'],
  ['--ff-background-color-primary', '#fff'],
]);

const darkThemeColors = new Map<string, string>([
  ['--ff-primary-color', '#4096ff'],
  ['--ff-text-color-primary', '#fff'],
  ['--ff-background-color-primary', '#131212'],
]);

const setTheme = (theme: 'light' | 'dark') => {
  const themeMap = theme === 'light' ? darkThemeColors : lightThemeColors;

  themeMap.forEach((v, k) => {
    document.documentElement.style.setProperty(k, v);
  });
};

export const themeSlice = createSlice({
  name: 'theme',
  initialState,
  reducers: {
    toggleTheme: (state) => {
      const theme = state.theme;
      setTheme(theme);
      state.theme = theme === 'light' ? 'dark' : 'light';
    },
  },
});

export const { toggleTheme } = themeSlice.actions;

export default themeSlice.reducer;

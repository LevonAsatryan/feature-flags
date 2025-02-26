import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { LoadingState } from './types';
import axios from '../http/axios';

export type FeatureFlag = {
  id: string;
  name: string;
  value: boolean;
};

export type FeatureFlagState = {
  ffs: FeatureFlag[];
  loading: LoadingState;
};

const initialState: FeatureFlagState = {
  ffs: [],
  loading: 'idle',
};

export const fetchFeatureFlags = createAsyncThunk('ffs', async (groupId: string) => {
  const response = await axios?.get(`feature-flags/${groupId}`);
  return response.data;
});

export const ffsSlice = createSlice({
  name: 'ffs',
  initialState,
  reducers: {},
  extraReducers: (builders) => {
    builders.addCase(fetchFeatureFlags.fulfilled, (state, action) => {
      state.ffs = action.payload;
      state.loading = 'succeeded';
    });
    builders.addCase(fetchFeatureFlags.pending, (state) => {
      state.loading = 'pending';
    });
  },
});

export default ffsSlice.reducer;

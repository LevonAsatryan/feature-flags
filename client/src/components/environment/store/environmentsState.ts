import { createAsyncThunk, createSlice, PayloadAction } from '@reduxjs/toolkit';
import axios from '../../../http/axios';
import { LoadingState } from '../../../store/types';

export type Environment = {
  id: string;
  name: string;
};

export type EnvironmentsState = {
  environments: Environment[];
  loading: LoadingState;
};

const initialState: EnvironmentsState = {
  environments: [],
  loading: 'idle',
};

export const fetchEnvironment = createAsyncThunk('groups', async () => {
  const response = await axios?.get('/groups');
  return response?.data;
});

export const environmentSlice = createSlice({
  name: 'environments',
  initialState,
  reducers: {
    setEnvironments: (state, action: PayloadAction<Environment[]>) => {
      state.environments = action.payload;
    },
  },
  extraReducers: (builder) => {
    builder.addCase(fetchEnvironment.fulfilled, (state, action) => {
      state.environments = action.payload;
      state.loading = 'succeeded';
    });
    builder.addCase(fetchEnvironment.pending, (state) => {
      state.loading = 'pending';
    });
  },
});

export const { setEnvironments } = environmentSlice.actions;

export default environmentSlice.reducer;

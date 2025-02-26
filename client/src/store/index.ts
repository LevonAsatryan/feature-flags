import { configureStore } from '@reduxjs/toolkit';
import loginReducer from './loginState';
import environmentReducer from './environmentsState';
import themeReducer from './themeState';
import ffsReducer from './featureFlagsState';

export const store = configureStore({
  reducer: {
    login: loginReducer,
    environments: environmentReducer,
    theme: themeReducer,
    ffs: ffsReducer,
  },
});

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;

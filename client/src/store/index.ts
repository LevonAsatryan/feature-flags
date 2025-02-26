import { configureStore } from '@reduxjs/toolkit';
import loginReducer from './loginState';
import environmentReducer from '../components/environment/store/environmentsState';

export const store = configureStore({
  reducer: {
    login: loginReducer,
    environments: environmentReducer,
  },
});

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;

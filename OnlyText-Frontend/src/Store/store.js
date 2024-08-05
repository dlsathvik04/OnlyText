import { combineReducers, configureStore, getDefaultMiddleware } from '@reduxjs/toolkit';

const reducer = combineReducers({
  // Reducers here
});

export default function () {
  return configureStore({
    reducer,
    middleware: [...getDefaultMiddleware(),]
  });
}
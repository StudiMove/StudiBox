import { combineReducers } from '@reduxjs/toolkit';
import authReducer from './slices/auth.slice'; // On créera ce fichier ensuite
// import userReducer from './slices/user.slice'; // On créera ce fichier ensuite

const rootReducer = combineReducers({
  auth: authReducer,
  //   user: userReducer,
  // Ajoute d'autres reducers ici si nécessaire
});

export default rootReducer;

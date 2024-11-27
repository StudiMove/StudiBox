// src/store/index.ts
import { configureStore, combineReducers } from '@reduxjs/toolkit';
import { persistStore, persistReducer } from 'redux-persist';
import storage from 'redux-persist/lib/storage'; // Utilisation du local storage pour persister
import authReducer from './slices/auth.slice'; // Le reducer de ton slice d'auth
import profileImageReducer from './slices/profileImage.slice'; // Importer le slice

// Configuration de Redux Persist
const persistConfig = {
  key: 'root',
  storage, // On utilise le local storage pour la persistance
};

// Combinaison des reducers (tu pourras ajouter d'autres slices ici plus tard)
const rootReducer = combineReducers({
  auth: authReducer,
  profileImage: profileImageReducer, // Ajouter le nouveau slice

  // D'autres reducers peuvent être ajoutés ici
});

// Création du reducer persisté
const persistedReducer = persistReducer(persistConfig, rootReducer);

// Configuration du store
const store = configureStore({
  reducer: persistedReducer,
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false, // Désactivation de la vérification de sérialisation pour Redux Persist
    }),
});

// Persistance du store
export const persistor = persistStore(store);

// Définition du type RootState
export type RootState = ReturnType<typeof rootReducer>; // Le type global de l'état

// Exportation du store
export default store;

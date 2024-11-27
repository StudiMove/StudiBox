// utils/authHelpers.ts

import axios from 'axios';

// Stocke le token d'authentification dans localStorage

export const storeToken = (token: string) => {
  localStorage.setItem('authToken', token);
};

// Supprime le token d'authentification de localStorage
export const removeToken = () => {
  localStorage.removeItem('authToken');
};

// Définit ou supprime l'en-tête Authorization pour les requêtes axios

export const setAuthToken = (token: string | null) => {
  if (token) {
    axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
  } else {
    delete axios.defaults.headers.common['Authorization'];
  }
};

export const getTokenFromStorage = (): string | null => {
  return localStorage.getItem('authToken');
};

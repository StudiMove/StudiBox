import axios from 'axios';
import { LoginForm, SignupForm } from '../types/api/auth/AuthForm';
import { SignupResponse, LoginResponse } from '../types/api/auth/AuthResponse';
import { storeToken, setAuthToken, removeToken } from '../utils/authHelpers';

import {
  API_URL_LOGIN_USER,
  API_URL_GET_USER_ROLE,
  API_URL_REGISTER_ORGANISATION_USER,
} from '../config/apiEndpoints';

// Récupération du token depuis le localStorage au chargement
const token = localStorage.getItem('authToken');
if (token) {
  setAuthToken(token);
}

// Fonction pour récupérer le rôle de l'utilisateur
const fetchUserRole = async (token: string): Promise<string> => {
  const response = await axios.get(API_URL_GET_USER_ROLE, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return response.data.role;
};

// Fonction pour se déconnecter
export const logoutUser = () => {
  removeToken();
  setAuthToken(null);
};

// Fonction pour gérer l'inscription avec un formulaire plus détaillé
export const signupUser = async (
  signupData: SignupForm,
): Promise<SignupResponse> => {
  try {
    const response = await axios.post<SignupResponse>(
      API_URL_REGISTER_ORGANISATION_USER,
      signupData,
      {
        headers: {
          'Content-Type': 'application/json',
        },
      },
    );

    const token = response.data.token;
    if (token) {
      storeToken(token);
      setAuthToken(token); // Définit l'en-tête Authorization pour les futures requêtes
    }

    return response.data;
  } catch (error) {
    console.error('Error signing up user:', error);

    if (axios.isAxiosError(error) && error.response?.status === 400) {
      throw new Error("Erreur lors de l'inscription. Vérifiez vos données.");
    }

    throw new Error("Erreur lors de l'inscription. Veuillez réessayer.");
  }
};

// Fonction pour gérer la connexion utilisateur
export const loginUser = async (
  loginData: LoginForm,
): Promise<LoginResponse> => {
  try {
    const response = await axios.post(API_URL_LOGIN_USER, loginData);

    const token = response.data.token;
    const profileImage = { url: response.data.profile_image };

    if (token) {
      storeToken(token);
      setAuthToken(token); // Définit l'en-tête Authorization

      const userRole = await fetchUserRole(token);

      return { token, role: userRole, profileImage, isAuthenticated: true };
    }

    throw new Error('Token manquant');
  } catch (error) {
    console.error("Erreur lors de la connexion de l'utilisateur :", error);

    if (axios.isAxiosError(error)) {
      if (error.response?.status === 401) {
        throw new Error('Email ou mot de passe incorrect.');
      }
    }

    throw new Error('Erreur lors de la connexion. Veuillez réessayer.');
  }
};

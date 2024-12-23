// src/config/apiEndpoints.ts

const API_URL = 'https://api.studibox.fr/api/v1'; // Utilise le nom du service Docker "backend"
const BASE = 'http://localhost:8080'; 
// const API_URL = process.env.API_URL || 'http://localhost:8080/api/v1'; // Utilise le nom du service Docker "backend"

// EVENTS

//Il SONT PAS UTILISER OU PLUS A  JOURS
export const API_URL_GET_USER_ONLINE_EVENTS = `${API_URL}/events/userlist/online`; // Récupérer les événements en ligne de l'utilisateur
export const API_URL_GET_ALL_EVENTS = `${API_URL}/events/all`; // Récupérer tous les événements
export const API_URL_GET_ALL_ONLINE_EVENTS = `${API_URL}/events/all/online`; // Récupérer tous les événements
export const API_URL_UPDATE_EVENT = `${API_URL}/events/update`; // Mettre à jour un événement

export const API_URL_DELETE_EVENT = `${API_URL}/events/delete`; // Supprimer un événement
export const API_URL_HARD_DELETE_EVENT = `${API_URL}/events/delete/hard`; // Supprimer définitivement un événement

export const API_URL_GET_EVENT_CATEGORIES_BY_ID = (eventId: number): string =>
  `${API_URL}/events/categories/${eventId}`;

// REFACTO

export const API_URL_LOGIN_USER = `${API_URL}/auth/login`;
export const API_URL_REGISTER_USER = `${API_URL}/auth/register`;
export const API_URL_REGISTER_ORGANISATION_USER = `${API_URL}/auth/register/organisation`;
export const API_URL_GET_USER_ROLE = `${API_URL}/profil/user/role`; // Utilisation du nouveau préfixe

export const API_URL_UPDATE_PROFILE_DATA = `${API_URL}/profil/organisation/profile/update`;
export const API_URL_UPDATE_TARGET_PROFILE_DATA = `${API_URL}/profil/organisation/profile/update/targetId`;

export const API_URL_REQUEST_PASSWORD_RESET = `${API_URL}/password/request-reset`; // Envoi d'un code de réinitialisation
export const API_URL_VERIFY_RESET_CODE = `${API_URL}/password/verify-reset-code`; // Vérification du code de réinitialisation
export const API_URL_UPDATE_PASSWORD = `${API_URL}/password/update`; // Mise à jour du mot de passe

export const API_URL_GET_PROFIL_DATA = `${API_URL}/profil/organisation/profile`; // Utilisation du nouveau préfixe

export const API_URL_ALL_GET_ORGANISATION_DATA = `${API_URL}/organisations/all`;
export const API_URL_ALL_ACTIVE_GET_ORGANISATION_DATA = `${API_URL}/organisations/all/active`;
export const API_URL_ALL_INACTIVE_GET_ORGANISATION_DATA = `${API_URL}/organisations/all/inactive`;
export const API_URL_ALL_IS_PENDING_ORGANISATION_DATA = `${API_URL}/organisations/all/pending`;
export const API_URL_ALL_SUSPENDED_ORGANISATION_DATA = `${API_URL}/organisations/all/suspended`;

export const API_URL_GET_PROFIL_DATA_BY_ID = `${API_URL}/profil/organisation/profile/targetId`; // Modifier selon votre API

export const API_URL_GET_EVENT_BY_ID = `${API_URL}/events/get`; // Récupérer un événement spécifique par ID
export const API_URL_CREATE_EVENT = `${API_URL}/events/create`; // Créer un nouvel événement
export const API_URL_UPLOAD_EVENT = `${API_URL}/events/upload`; // Créer un nouvel événement

export const API_URL_UPDATE_PROFILE_IMAGE = `${API_URL}/profil/organisation/profile/upload-image`;
export const API_URL_UPDATE_TARGET_PROFILE_IMAGE = `${API_URL}/profil/organisation/profile/upload-image/targetId`;

export const API_URL_GET_ALL_EVENT_TAGS = `${API_URL}/events/tags`;
export const API_URL_GET_ALL_EVENT_CATEGORIES = `${API_URL}/events/categories`;

export const API_URL_GET_USER_EVENTS = `${API_URL}/events/list`; // Récupérer les événements de l'utilisateur
export const API_URL_GET_TARGETLIST_EVENT = `${API_URL}/events/list/target`;
export const API_URL_UPLOAD_EVENT_IMAGE = `${API_URL}/events/upload-image`; // Mettre à jour un événement

export const API_URL_GET_TICKET_DETAILS = `${API_URL}/ticket`;

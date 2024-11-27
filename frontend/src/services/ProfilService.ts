import axios from 'axios';
import {
  API_URL_GET_PROFIL_DATA,
  API_URL_UPDATE_PROFILE_IMAGE,
  API_URL_GET_PROFIL_DATA_BY_ID,
  API_URL_UPDATE_PROFILE_DATA,
  API_URL_UPDATE_TARGET_PROFILE_DATA,
  API_URL_UPDATE_TARGET_PROFILE_IMAGE,
} from '../config/apiEndpoints';
import { ProfileForm } from '../types/api/profile/ProfilForm/ProfilForm';
import { ProfileResponse } from '../types/api/profile/ProfileResponse';

import {
  ProfileUpdateData,
  TargetProfileUpdateData,
} from '../types/api/profile/ProfileRequest';
import {
  getTokenFromStorage,
  setAuthToken,
  removeToken,
} from '../utils/authHelpers';

export const getProfileData = async (): Promise<ProfileForm> => {
  try {
    const token = getTokenFromStorage(); // Récupérez le token depuis le stockage
    setAuthToken(token); // Définissez l'authentification avec le token

    // Effectuez la requête en utilisant le token dans l'en-tête
    const response = await axios.get<ProfileResponse>(API_URL_GET_PROFIL_DATA, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    // Mappage de la réponse API en objet ProfileForm
    const data = response.data;
    const mappedData: ProfileForm = {
      userId: data.user_id,
      email: data.email,
      phone: data.phone,
      profileImage: data.profile_image,
      roles: data.roles,
      organisationName: data.organisation.name,
      organisationAddress: data.organisation.address,
      organisationCity: data.organisation.city,
      organisationPostcode: data.organisation.postcode,
      organisationRegion: data.organisation.region,
      organisationCountry: data.organisation.country,
      organisationDescription: data.organisation.description,
      status: data.organisation.status,
      siret: data.organisation.siret || '', // Valeur par défaut
      type: data.roles[0] || '', // Valeur par défaut si le rôle est vide
      is_validated: data.is_validated || false, // Valeur par défaut
      is_pending: data.is_pending || true, // Valeur par défaut
      is_activated: data.is_activated || false, // Valeur par défaut
    };

    return mappedData; // Retourne les données mappées au format ProfileForm
  } catch (error) {
    console.error(
      'Erreur lors de la récupération des données du profil:',
      error,
    );

    // Supprimez le token si une erreur d'authentification est rencontrée
    if (axios.isAxiosError(error) && error.response?.status === 401) {
      removeToken();
    }

    throw error; // Propagation de l'erreur pour gestion en aval
  }
};

// Fonction pour récupérer les données d'un profil par son ID
export const getProfileDataById = async (
  targetId: number,
): Promise<ProfileForm> => {
  try {
    const response = await axios.post<ProfileResponse>(
      API_URL_GET_PROFIL_DATA_BY_ID,
      { targetId },
    );

    // Mappage de la réponse API en objet ProfileForm
    const data = response.data;
    const mappedData: ProfileForm = {
      userId: data.user_id,
      email: data.email,
      phone: data.phone,
      profileImage: data.profile_image,
      roles: data.roles,
      organisationName: data.organisation.name,
      organisationAddress: data.organisation.address,
      organisationCity: data.organisation.city,
      organisationPostcode: data.organisation.postcode,
      organisationCountry: data.organisation.country,
      organisationRegion: data.organisation.region,
      organisationDescription: data.organisation.description,
      status: data.organisation.status,
      siret: data.organisation.siret || '', // Valeur par défaut
      type: data.roles[0] || '', // Valeur par défaut si le rôle est vide
      is_validated: data.is_validated || false, // Valeur par défaut
      is_pending: data.is_pending || true, // Valeur par défaut
      is_activated: data.is_activated || false, // Valeur par défaut
    };

    return mappedData; // Retourne les données mappées au format ProfileForm
  } catch (error) {
    console.error(
      'Erreur lors de la récupération des données du profil par ID:',
      error,
    );
    throw new Error('Impossible de récupérer les données du profil.');
  }
};

// Mise à jour du profil de l'utilisateur courant
export const updateProfileData = async (
  profileData: ProfileUpdateData,
): Promise<void> => {
  try {
    const token = getTokenFromStorage();

    await axios.put(API_URL_UPDATE_PROFILE_DATA, profileData, {
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    });
  } catch (error) {
    console.error(
      'Erreur lors de la mise à jour des données du profil:',
      error,
    );
    throw error;
  }
};

// Fonction pour mettre à jour le profil d'un utilisateur cible
export const updateTargetProfileData = async (
  dataToSend: TargetProfileUpdateData,
): Promise<void> => {
  try {
    const token = getTokenFromStorage();
    if (!token)
      throw new Error('Token introuvable, veuillez vous reconnecter.');

    await axios.put(API_URL_UPDATE_TARGET_PROFILE_DATA, dataToSend, {
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    });
  } catch (error) {
    console.error(
      'Erreur lors de la mise à jour des données du profil cible:',
      error,
    );
    throw new Error('Impossible de mettre à jour les données du profil cible.');
  }
};

export const uploadTargetProfileImage = async (
  targetId: number,
  formData: FormData,
): Promise<string | undefined> => {
  try {
    const token = getTokenFromStorage();
    formData.append('targetId', targetId.toString());

    const response = await axios.post(
      API_URL_UPDATE_TARGET_PROFILE_IMAGE,
      formData,
      {
        headers: {
          Authorization: `Bearer ${token}`,
          'Content-Type': 'multipart/form-data',
        },
      },
    );
    console.log('Image de profil mise à jour pour targetId:', targetId);
    return response.data.url; // Retourne l'URL de l'image depuis la réponse du serveur
  } catch (error) {
    console.error(
      "Erreur lors de l'upload de l'image de profil pour targetId:",
      error,
    );
    throw error;
  }
};

// Fonction pour uploader l'image de profil de l'utilisateur connecté
export const updateProfileImage = async (
  formData: FormData,
): Promise<string> => {
  try {
    const token = getTokenFromStorage();
    const response = await axios.post(API_URL_UPDATE_PROFILE_IMAGE, formData, {
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'multipart/form-data',
      },
    });

    console.log("Réponse du serveur après l'envoi de l'image:", response.data);
    return response.data.url; // Retourne l'URL de l'image
  } catch (error) {
    console.error("Erreur lors de la mise à jour de l'image de profil:", error);
    throw error;
  }
};

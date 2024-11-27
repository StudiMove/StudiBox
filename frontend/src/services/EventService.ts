import axios from 'axios';
import {
  API_URL_CREATE_EVENT,
  API_URL_GET_USER_EVENTS,
  API_URL_GET_USER_ONLINE_EVENTS,
  API_URL_GET_ALL_ONLINE_EVENTS,
  API_URL_GET_ALL_EVENTS,
  API_URL_GET_EVENT_BY_ID,
  API_URL_UPDATE_EVENT,
  API_URL_DELETE_EVENT,
  API_URL_HARD_DELETE_EVENT,
  API_URL_GET_TARGETLIST_EVENT,
  API_URL_UPLOAD_EVENT_IMAGE,
  API_URL_GET_ALL_EVENT_TAGS,
  API_URL_GET_ALL_EVENT_CATEGORIES,
  API_URL_GET_EVENT_CATEGORIES_BY_ID,
} from '../config/apiEndpoints';
import { EventCreateData } from '../types/api/event/EventCreateData';
import { EventData } from '../types/api/event/EventTypes';
import { EventUpdateData } from '../types/api/event/EventTypes';
// Fonction pour récupérer le token du localStorage
const getTokenFromStorage = (): string | null =>
  localStorage.getItem('authToken');

// Récupérer tous les événements
export const getAllEvents = async (): Promise<Event[]> => {
  try {
    const token = getTokenFromStorage();
    if (!token) throw new Error('Token manquant. Veuillez vous reconnecter.');

    const response = await axios.get<Event[]>(API_URL_GET_ALL_EVENTS, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return response.data;
  } catch (error) {
    console.error(
      'Erreur lors de la récupération de tous les événements:',
      error,
    );
    throw error;
  }
};

// Récupérer les événements de l'utilisateur
export const getUserEvents = async (): Promise<Event[]> => {
  try {
    const token = getTokenFromStorage();
    if (!token) throw new Error('Token manquant. Veuillez vous reconnecter.');

    const response = await axios.get<Event[]>(API_URL_GET_USER_EVENTS, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return response.data;
  } catch (error) {
    console.error(
      'Erreur lors de la récupération des événements utilisateur:',
      error,
    );
    throw error;
  }
};

// Récupérer les événements en ligne de l'utilisateur
export const getUserOnlineEvents = async (): Promise<Event[]> => {
  try {
    const token = getTokenFromStorage();
    if (!token) throw new Error('Token manquant. Veuillez vous reconnecter.');

    const response = await axios.get<Event[]>(API_URL_GET_USER_ONLINE_EVENTS, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return response.data;
  } catch (error) {
    console.error(
      'Erreur lors de la récupération des événements en ligne:',
      error,
    );
    throw error;
  }
};
export const getAllOnlineEvents = async (): Promise<Event[]> => {
  try {
    const token = getTokenFromStorage();
    if (!token) throw new Error('Token manquant. Veuillez vous reconnecter.');

    const response = await axios.get<Event[]>(API_URL_GET_ALL_ONLINE_EVENTS, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return response.data;
  } catch (error) {
    console.error(
      'Erreur lors de la récupération des événements en ligne:',
      error,
    );
    throw error;
  }
};

// // Service pour mettre à jour un événement spécifique
// export const updateEvent = async (
//   eventData: EventUpdateData,
// ): Promise<EventUpdateData> => {
//   try {
//     if (Array.isArray(eventData.images)) {
//       eventData.images = eventData.images.join(', ');
//     } else if (typeof eventData.images === 'string') {
//       eventData.images = eventData.images;
//     } else {
//       eventData.images = '';
//     }

//     const token = getTokenFromStorage();
//     if (!token) throw new Error('Token manquant. Veuillez vous reconnecter.');

//     // Effectuer une requête PUT pour mettre à jour les informations de l'événement
//     const response = await axios.put<EventUpdateData>(
//       API_URL_UPDATE_EVENT,
//       eventData,
//       {
//         headers: {
//           Authorization: `Bearer ${token}`,
//         },
//       },
//     );
//     return response.data;
//   } catch (error) {
//     console.error("Erreur lors de la mise à jour de l'événement :", error);
//     throw error;
//   }
// };
// Service pour mettre à jour un événement spécifique

export const updateEvent = async (
  eventData: EventUpdateData,
  existingImages: string[] = [], // Ajout d'un paramètre pour les images existantes
): Promise<EventUpdateData> => {
  try {
    // Vérifier s'il y a des nouvelles images; sinon, utiliser les images existantes
    if (Array.isArray(eventData.images) && eventData.images.length > 0) {
      eventData.images = eventData.images.join(', ');
    } else if (typeof eventData.images === 'string' && eventData.images) {
      eventData.images = eventData.images;
    } else {
      // Utiliser les images existantes si aucune nouvelle image n'est fournie
      eventData.images = existingImages.join(', ');
    }

    const token = getTokenFromStorage();
    if (!token) throw new Error('Token manquant. Veuillez vous reconnecter.');

    // Effectuer une requête PUT pour mettre à jour les informations de l'événement
    const response = await axios.put<EventUpdateData>(
      API_URL_UPDATE_EVENT,
      eventData,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      },
    );
    return response.data;
  } catch (error) {
    console.error("Erreur lors de la mise à jour de l'événement :", error);
    throw error;
  }
};
export const getEventById = async (eventId: number): Promise<EventData> => {
  try {
    const token = getTokenFromStorage();
    if (!token) throw new Error('Token manquant. Veuillez vous reconnecter.');

    // Effectuer une requête POST avec `event_id` dans le corps
    const response = await axios.post<EventData>(
      API_URL_GET_EVENT_BY_ID,
      { event_id: eventId }, // Passer `event_id` dans le corps de la requête
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      },
    );

    return response.data;
  } catch (error) {
    console.error("Erreur lors de la récupération de l'événement:", error);
    throw error;
  }
};

// Créer un nouvel événement
export const createEvent = async (
  eventData: EventCreateData,
): Promise<EventCreateData> => {
  try {
    // Vérifiez si images est un tableau, puis joignez les éléments en une seule chaîne
    if (Array.isArray(eventData.images)) {
      eventData.images = eventData.images.join(', ');
    }

    const token = getTokenFromStorage();
    if (!token) throw new Error('Token manquant. Veuillez vous reconnecter.');
    console.log(eventData);
    const response = await axios.post<EventCreateData>(
      API_URL_CREATE_EVENT,
      eventData,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      },
    );
    return response.data;
  } catch (error) {
    console.error("Erreur lors de la création de l'événement:", error);
    throw error;
  }
};

// Récupérer tous les tags
export const getAllEventTags = async (): Promise<string[]> => {
  try {
    const token = getTokenFromStorage();
    if (!token) throw new Error('Token manquant. Veuillez vous reconnecter.');

    const response = await axios.get<string[]>(API_URL_GET_ALL_EVENT_TAGS, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return response.data; // Retourner les noms des tags
  } catch (error) {
    console.error('Erreur lors de la récupération de tous les tags:', error);
    throw error;
  }
};

// Récupérer toutes les catégories
export const getAllEventCategories = async (): Promise<string[]> => {
  try {
    const token = getTokenFromStorage();
    if (!token) throw new Error('Token manquant. Veuillez vous reconnecter.');

    const response = await axios.get<string[]>(
      API_URL_GET_ALL_EVENT_CATEGORIES,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      },
    );
    return response.data; // Retourner les noms des catégories
  } catch (error) {
    console.error(
      'Erreur lors de la récupération de toutes les catégories:',
      error,
    );
    throw error;
  }
};

export const deleteEvent = async (eventId: number): Promise<void> => {
  try {
    const token = getTokenFromStorage();
    if (!token) throw new Error('Token manquant. Veuillez vous reconnecter.');

    const response = await axios.delete<void>(
      `${API_URL_DELETE_EVENT}/${eventId}`,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      },
    );
    return response.data;
  } catch (error) {
    console.error("Erreur lors de la suppression de l'événement:", error);
    throw error;
  }
};

// Supprimer définitivement un événement
export const hardDeleteEvent = async (eventId: number): Promise<void> => {
  try {
    const token = getTokenFromStorage();
    if (!token) throw new Error('Token manquant. Veuillez vous reconnecter.');

    const response = await axios.delete<void>(
      `${API_URL_HARD_DELETE_EVENT}/${eventId}`,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      },
    );
    return response.data;
  } catch (error) {
    console.error(
      "Erreur lors de la suppression définitive de l'événement:",
      error,
    );
    throw error;
  }
};
export const getUserEventsByOrganisation = async (
  user_id: number,
): Promise<Event[]> => {
  try {
    const token = getTokenFromStorage();
    if (!token) throw new Error('Token manquant. Veuillez vous reconnecter.');

    // Envoyer la requête avec targetUserID dans le corps de la requête
    const response = await axios.post<Event[]>(
      API_URL_GET_TARGETLIST_EVENT, // URL de l'API
      { user_id: user_id }, // Le corps de la requête contenant userTargetID
      {
        headers: {
          Authorization: `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
      },
    );

    return response.data;
  } catch (error) {
    console.error(
      'Erreur lors de la récupération des événements par organisation:',
      error,
    );
    throw error;
  }
};
// Fonction pour uploader les images

export const uploadEventImages = async (files: File[]): Promise<string[]> => {
  try {
    console.log('Étape 1: Récupération du token');
    const token = getTokenFromStorage();
    if (!token) throw new Error('Token manquant. Veuillez vous reconnecter.');

    console.log('Étape 2: Initialisation de FormData');
    const formData = new FormData();
    files.forEach((file, index) => {
      formData.append('file', file);
      console.log(`Fichier ajouté au FormData: ${file.name}, index: ${index}`);
    });

    console.log('Étape 3: Envoi de la requête POST à l’API');
    const response = await axios.post<{ urls: string[] }>(
      API_URL_UPLOAD_EVENT_IMAGE,
      formData,
      {
        headers: {
          Authorization: `Bearer ${token}`,
          'Content-Type': 'multipart/form-data',
        },
      },
    );

    console.log('Étape 4: Réponse reçue de l’API', response.data);

    return response.data.urls; // Retourner les URLs des images uploadées
  } catch (error) {
    console.error("Erreur lors de l'upload des images:", error);
    throw error;
  }
};

export const getEventCategoriesById = async (eventId: number): Promise<any> => {
  // Déclarez le type de retour
  try {
    const response = await axios.get(
      API_URL_GET_EVENT_CATEGORIES_BY_ID(eventId),
    );
    return response.data; // Renvoie les catégories
  } catch (error) {
    console.error(
      "Erreur lors de la récupération des catégories de l'événement:",
      error,
    );
    // Si vous utilisez TypeScript, vous pouvez gérer l'erreur comme ceci
    if (axios.isAxiosError(error)) {
      // Gérer les erreurs Axios ici
      throw new Error(
        error.response?.data ||
          "Une erreur s'est produite lors de la récupération des catégories",
      );
    } else {
      // Gérer les autres types d'erreurs ici
      throw new Error("Une erreur inconnue s'est produite");
    }
  }
};

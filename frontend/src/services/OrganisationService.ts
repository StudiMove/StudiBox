// src/services/OrganisationService.ts

import axios from 'axios';
import {
  API_URL_ALL_GET_ORGANISATION_DATA,
  API_URL_ALL_ACTIVE_GET_ORGANISATION_DATA,
  API_URL_ALL_INACTIVE_GET_ORGANISATION_DATA,
  API_URL_ALL_IS_PENDING_ORGANISATION_DATA,
  API_URL_ALL_SUSPENDED_ORGANISATION_DATA,
} from '../config/apiEndpoints';
import { Organisation } from '../types/api/organisation/OrganisationTypes';
import { OrganisationsResponse } from '../types/api/organisation/OrganisationsResponse';
// Fonction générique pour récupérer des organisations selon une URL spécifique
const fetchOrganisations = async (url: string): Promise<Organisation[]> => {
  try {
    const response = await axios.get<OrganisationsResponse>(url);
    return response.data.organisations;
  } catch (error) {
    console.error('Erreur lors de la récupération des organisations:', error);
    throw new Error('Impossible de récupérer les organisations');
  }
};

// Récupérer toutes les organisations
export const getAllOrganisations = async (): Promise<Organisation[]> => {
  return fetchOrganisations(API_URL_ALL_GET_ORGANISATION_DATA);
};

// Récupérer les organisations actives
export const getActiveOrganisations = async (): Promise<Organisation[]> => {
  return fetchOrganisations(API_URL_ALL_ACTIVE_GET_ORGANISATION_DATA);
};

// Récupérer les organisations inactives
export const getInactiveOrganisations = async (): Promise<Organisation[]> => {
  return fetchOrganisations(API_URL_ALL_INACTIVE_GET_ORGANISATION_DATA);
};

// Récupérer les organisations en attente
export const getPendingOrganisations = async (): Promise<Organisation[]> => {
  return fetchOrganisations(API_URL_ALL_IS_PENDING_ORGANISATION_DATA);
};

// Récupérer les organisations suspendues
export const getSuspendedOrganisations = async (): Promise<Organisation[]> => {
  return fetchOrganisations(API_URL_ALL_SUSPENDED_ORGANISATION_DATA);
};

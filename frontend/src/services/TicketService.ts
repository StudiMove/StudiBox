import axios from 'axios';
import { TicketDetailsResponse } from '../types/api/ticket/TicketResponse';
import { API_URL_GET_TICKET_DETAILS } from '../config/apiEndpoints';

// Fonction pour récupérer les détails d'un ticket
export const getTicketDetails = async (
    uuid: string,
  ): Promise<TicketDetailsResponse> => {
    try {
      const response = await axios.get<TicketDetailsResponse>(
        `${API_URL_GET_TICKET_DETAILS}/${uuid}`,
      );
      return response.data;
    } catch (error) {
      console.error('Erreur lors de la récupération des détails du ticket:', error);
      throw error;
    }
  };
  
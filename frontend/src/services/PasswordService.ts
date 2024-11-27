// services/PasswordService.ts
import axios from 'axios';
import {
  RequestPasswordReset,
  VerifyResetCodeRequest,
  UpdatePasswordRequest,
} from '../types/api/password/PasswordRequest';
import {
  PasswordResetResponse,
  VerifyResetCodeResponse,
} from '../types/api/password/PasswordResponse';

import {
  API_URL_REQUEST_PASSWORD_RESET,
  API_URL_VERIFY_RESET_CODE,
  API_URL_UPDATE_PASSWORD,
} from '../config/apiEndpoints';

export const requestPasswordReset = async (
  email: string,
): Promise<PasswordResetResponse> => {
  try {
    const response = await axios.post<PasswordResetResponse>(
      API_URL_REQUEST_PASSWORD_RESET,
      { email } as RequestPasswordReset,
    );
    return response.data;
  } catch (error) {
    // Gérer l'erreur, par exemple, en la journalisant ou en renvoyant une valeur par défaut
    console.error(
      'Erreur lors de la demande de réinitialisation de mot de passe:',
      error,
    );
    throw error;
  }
};

export const verifyResetCode = async (
  email: string,
  resetCode: number,
): Promise<VerifyResetCodeResponse> => {
  try {
    const response = await axios.post<VerifyResetCodeResponse>(
      API_URL_VERIFY_RESET_CODE,
      { email, reset_code: resetCode } as VerifyResetCodeRequest,
    );
    return response.data;
  } catch (error) {
    console.error(
      'Erreur lors de la vérification du code de réinitialisation:',
      error,
    );
    throw error;
  }
};

export const updatePassword = async (
  email: string,
  newPassword: string,
): Promise<PasswordResetResponse> => {
  try {
    const response = await axios.put<PasswordResetResponse>(
      API_URL_UPDATE_PASSWORD,
      {
        email,
        new_password: newPassword,
      } as UpdatePasswordRequest,
    );
    return response.data;
  } catch (error) {
    console.error('Erreur lors de la mise à jour du mot de passe:', error);
    throw error;
  }
};

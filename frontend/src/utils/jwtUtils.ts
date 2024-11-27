// src/utils/jwtUtils.ts
import { jwtDecode } from 'jwt-decode';

export type DecodedToken = {
  exp: number; // Date d'expiration du token
  iat: number; // Date de création du token
  user_id: number; // ID de l'utilisateur
};

export const decodeJWT = (token: string): DecodedToken | null => {
  try {
    return jwtDecode<DecodedToken>(token); // Assurez-vous que jwtDecode est correctement utilisé ici
  } catch (error) {
    console.error('Erreur lors du décodage du token JWT :', error);
    return null;
  }
};

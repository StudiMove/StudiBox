// src/types/DataTypes.ts
export interface User {
  id: number;
  profileImage: string;
  pseudo: string;
  firstName: string;
  email: string;
  type: string;
}
export interface Company {
  user_id: number; // Identifiant unique de l'utilisateur ou de l'organisation
  profile_image: string; // URL de l'image de profil de l'organisation
  name: string; // Nom de l'organisation
  role_name: string; // Rôle de l'organisation (admin, utilisateur, etc.)
  status: string; // Statut de l'organisation (validé, en attente, suspendu, etc.)
  is_validated: boolean; // Indicateur si l'organisation est validée
  is_activated: boolean; // Indicateur si l'organisation est activée
  is_pending: boolean; // Indicateur si l'organisation est en attente
}

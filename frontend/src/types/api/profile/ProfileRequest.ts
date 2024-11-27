// types/api/profil/ProfileRequest.ts

export interface ProfileUpdateData {
  name?: string;
  address?: string;
  city?: string;
  postcode?: string;
  region?: string;
  phone?: string;
  country?: string;
  email?: string;
  siret?: string;
  description?: string;
  status?: string;
  type?: string;
  is_validated?: boolean;
  is_pending?: boolean;
  is_activated?: boolean;
  profileImage?: string; // Nouveau champ pour l'URL de l'image de profil
}

export interface TargetProfileUpdateData {
  targetId: number;
  updateData: ProfileUpdateData;
}

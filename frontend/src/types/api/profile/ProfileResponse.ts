// src/types/api/profile/ProfileResponse.ts
export interface ProfileResponse {
  user_id: number;
  email: string;
  phone: string;
  profile_image: string;
  roles: string[];
  is_validated: boolean;
  is_pending: boolean;
  is_activated: boolean;
  organisation: {
    name: string;
    address: string;
    city: string;
    postcode: string;
    region: string;
    country: string;
    description: string;
    status: string;
    siret: string; // Ajouté
    type: string; // Ajouté
  };
}

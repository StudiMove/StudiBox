// src/types/ProfilForm/ProfilForm.ts
export interface ProfileForm {
  userId: number;
  email: string;
  phone: string;
  profileImage: string;
  roles: string[];
  organisationName: string;
  organisationAddress: string;
  organisationCity: string;
  organisationPostcode: string;
  organisationCountry: string;
  organisationRegion: string;
  organisationDescription: string;
  status: string;
  siret: string; // Ajouté
  type: string; // Ajouté
  is_validated: boolean; // Ajouté
  is_pending: boolean; // Ajouté
  is_activated: boolean; // Ajouté
}

import { ProfileImage } from '../profile/Profile';

export interface SignupResponse {
  token: any;
  userId: number;
  message: string;
  success: boolean;
}

export interface LoginResponse {
  token: string;
  role: string;
  profileImage: ProfileImage; // Profil image selon le format de modèle
  isAuthenticated: boolean;
}

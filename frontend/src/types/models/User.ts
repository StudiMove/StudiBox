import { Role } from './Role';

export interface User {
  id: number;
  firstName?: string;
  lastName?: string;
  pseudo?: string;
  email: string;
  password: string;
  phone?: string;
  profileImage?: string;
  birthDate?: string;
  city?: string;
  profileType?: string;
  associationId?: number;
  schoolId?: number;
  studiboxCoins: number;
  roles: Role[];
  createdAt: string;
  updatedAt: string;
}

import { User } from './User';

export interface SchoolUser {
  userId: number;
  user: User;
  name: string;
  address?: string;
  city?: string;
  postcode?: string;
  region?: string;
  country?: string;
  description?: string;
  status?: string;
  isValidated: boolean;
  isActivated: boolean;
  isPending: boolean;
  memberCount: number;
}

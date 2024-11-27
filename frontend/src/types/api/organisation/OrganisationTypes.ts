export interface Organisation {
  user_id: number;
  name: string;
  is_validated: boolean;
  is_activated: boolean;
  is_pending: boolean;
  role_name: string;
  profile_image: string;
  status: string;
}

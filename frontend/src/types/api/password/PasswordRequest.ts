// types/api/password/PasswordRequest.ts
export interface RequestPasswordReset {
  email: string;
}

export interface VerifyResetCodeRequest {
  email: string;
  reset_code: number;
}

export interface UpdatePasswordRequest {
  email: string;
  reset_code: string;
  new_password: string;
}

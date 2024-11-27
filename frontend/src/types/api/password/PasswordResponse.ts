// types/api/password/PasswordResponse.ts
export interface PasswordResetResponse {
  success: boolean;
  message: string;
}

export interface VerifyResetCodeResponse {
  success: boolean;
  message: string;
}

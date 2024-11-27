export interface SignupForm {
  email: string;
  password: string;
  organisationName: string;
  address: string;
  postalCode: string;
  city: string;
  country: string;
  phone: string;
  description: string;
  organisationType: 'business' | 'association' | 'school';
}

export interface LoginForm {
  email: string;
  password: string;
}

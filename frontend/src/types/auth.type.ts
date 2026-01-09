export type RegistrationRequest = {
  name: string;
  email: string;
  phone_number: string;
  password: string;
};

export type LoginRequest = {
  email: string;
  password: string;
};

export type ChangePasswordRequest = {
  old_password: string;
  new_password: string;
};

export type ResetPasswordRequest = {
  email: string;
  otp: string;
  new_password: string;
};

export type RegistrationResponse = RegistrationRequest & {
  id: string;
};

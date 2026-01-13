export type PatientMinimumResponse = {
  id: string;
  name: string;
  email: string;
  phone_number: string;
  avatar_url?: string;
};

export type PatientResponse = {
  id: string;
  birth_date: string;
  gender: string;
  identity_number: string;
  blood_type: string;
  weight_kg: number;
  height_cm: number;
  allergies?: string | null;
  history_chronic_diseases?: string | null;
  emergency_contact?: string | null;
  insurance_provider?: string | null;
  insurance_number?: string | null;
  occupation?: string | null;
};

export type PatientUpdateRequest = Omit<
  PatientResponse,
  "id" | "identity_number"
>;

export type OnBoardPatientRequest = {
  birth_date: string;
  gender: string;
  identity_number: string;
  blood_type: string;
  weight_kg: number;
  height_cm: number;
  allergies?: string | null;
  history_chronic_diseases?: string | null;
  emergency_contact?: string | null;
  insurance_provider?: string | null;
  insurance_number?: string | null;
  occupation?: string | null;
};

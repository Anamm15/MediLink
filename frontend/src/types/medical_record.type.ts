export type MedicalRecordResponse = {
  id: string;
  appointment_id: string;
  patient_id: string;
  doctor_id: string;
  title: string;
  date: string;
  subjective?: string;
  objective?: string;
  assessment?: string;
  plan?: string;
  created_at: string;
};

export type MedicalRecordCreateRequest = {
  patient_id: string;
  appointment_id: string;
  title: string;
  date: string;
  subjective?: string;
  objective?: string;
  assessment?: string;
  plan?: string;
};

export type MedicalRecordUpdateRequest = {
  title?: string;
  subjective?: string;
  objective?: string;
  assessment?: string;
  plan?: string;
};

export type MedicalRecordDeleteRequest = {
  doctor_id: string;
};

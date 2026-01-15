import { DoctorMinimumResponse } from "./doctor.type";
import { MedicineResponse } from "./medicine.type";
import { Metadata } from "./metadata.type";
import { PatientMinimumResponse } from "./patient.type";

export type PrescriptionMedicinesCreate = {
  medicine_id: string;
  quantity: number;
};

export type PrescriptionResponse = {
  id: string;
  doctor: DoctorMinimumResponse;
  patient: PatientMinimumResponse;
  medical_record_id: string;
  prescription: string;
  notes: string;
  is_redeemed: boolean;
  medicines: (MedicineResponse & { quantity: number })[];
  created_at: string;
};

export type PrescriptionPaginateResponse = {
  data: PrescriptionResponse[];
  metadata: Metadata;
};

export type PrescriptionCreateRequest = {
  patient_id: string;
  medical_record_id: string;
  notes: string;
  medicines: PrescriptionMedicinesCreate[];
};

export type PrescriptionUpdateRequest = {
  notes: string;
  is_redeemed: boolean;
};

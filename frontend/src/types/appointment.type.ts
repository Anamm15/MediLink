import { DoctorMinimumResponse } from "./doctor.type";
import { PatientMinimumResponse } from "./patient.type";

export type AppointmentDetailResponse = {
  id: string;
  doctor: DoctorMinimumResponse;
  patient: PatientMinimumResponse;
  appointment_date: string;
  start_time: string;
  end_time: string;
  type: string;
  status: string;
  consultation_fee_snapshot: number;
  queue_number: number;
  meeting_link: string;
  symptom_complaint: string;
  doctor_notes: string;
};

export type CreateBookingRequest = {
  doctor_id: string;
  schedule_id: string;
  appointment_date: string;
  symptom_complaint?: string;
};

export type BookingResponse = {
  appointment_id: string;
  payment_url: string;
};

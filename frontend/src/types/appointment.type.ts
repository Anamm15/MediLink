export type AppointmentDetailResponse = {
  id: string;
  doctor_name: string;
  doctor_specialization: string;
  doctor_phone_number: string;
  patient_name: string;
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
};

export type BookingResponse = {
  appointment_id: string;
  payment_url: string;
};

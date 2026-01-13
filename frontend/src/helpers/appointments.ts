import { AppointmentDetailResponse } from "@/types/appointment.type";

export type GroupedAppointments = {
  upcoming: AppointmentDetailResponse[];
  past: AppointmentDetailResponse[];
};

export function groupAppointmentsByTime(
  appointments: AppointmentDetailResponse[]
): GroupedAppointments {
  const now = new Date();

  const upcoming: AppointmentDetailResponse[] = [];
  const past: AppointmentDetailResponse[] = [];

  appointments.forEach((appointment) => {
    const appointmentDateTime = new Date(
      `${appointment.appointment_date}T${appointment.start_time}`
    );

    if (appointmentDateTime >= now) {
      upcoming.push(appointment);
    } else {
      past.push(appointment);
    }
  });

  return {
    upcoming,
    past,
  };
}

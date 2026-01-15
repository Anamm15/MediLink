import { AppointmentResponse } from "@/types/appointment.type";

export type GroupedAppointments = {
  upcoming: AppointmentResponse[];
  past: AppointmentResponse[];
};

export function groupAppointmentsByTime(
  appointments: AppointmentResponse[]
): GroupedAppointments {
  const now = new Date();

  const upcoming: AppointmentResponse[] = [];
  const past: AppointmentResponse[] = [];

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

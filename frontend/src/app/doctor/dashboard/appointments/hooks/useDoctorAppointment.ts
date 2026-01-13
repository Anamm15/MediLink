import { getDoctorAppointments } from "@/services/appointment.service";
import { useQuery } from "@tanstack/react-query";

export function useDoctorAppointmentsQuery() {
  return useQuery({
    queryKey: ["doctor-appointments"],
    queryFn: getDoctorAppointments,
  });
}

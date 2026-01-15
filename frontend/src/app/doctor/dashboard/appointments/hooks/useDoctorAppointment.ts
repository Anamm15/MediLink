import { getDoctorAppointments } from "@/services/appointment.service";
import { useQuery } from "@tanstack/react-query";

export function useDoctorAppointmentsQuery(page: number, limit: number) {
  return useQuery({
    queryKey: ["doctor-appointments", page],
    queryFn: () => getDoctorAppointments(page, limit),
    staleTime: 3 * 60 * 1000,
  });
}

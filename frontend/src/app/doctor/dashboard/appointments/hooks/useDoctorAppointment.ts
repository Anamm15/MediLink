import { getDoctorAppointments } from "@/services/appointment.service";
import { useQuery } from "@tanstack/react-query";

export function useDoctorAppointmentsQuery(
  id: string,
  page: number,
  limit: number,
  status?: string
) {
  return useQuery({
    queryKey: ["doctor-appointments", id, page, status],
    queryFn: () => getDoctorAppointments(id, page, limit, status),
    staleTime: 3 * 60 * 1000,
  });
}

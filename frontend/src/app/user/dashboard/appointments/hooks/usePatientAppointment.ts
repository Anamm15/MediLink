import { useQuery } from "@tanstack/react-query";
import { getPatientAppointments } from "@/services/appointment.service";

export function usePatientAppointmentsQuery(
  id: string,
  page: number,
  limit: number,
  status?: string
) {
  return useQuery({
    queryKey: ["patient-appointments", page, status],
    queryFn: () => getPatientAppointments(id, page, limit, status),
    enabled: !!id,
    staleTime: 3 * 60 * 1000,
  });
}

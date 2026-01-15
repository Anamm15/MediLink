import { useQuery } from "@tanstack/react-query";
import { getPatientAppointments } from "@/services/appointment.service";

export function usePatientAppointmentsQuery(page: number, limit: number) {
  return useQuery({
    queryKey: ["patient-appointments", page],
    queryFn: () => getPatientAppointments(page, limit),
  });
}

import { useQuery } from "@tanstack/react-query";
import { getPatientAppointments } from "@/services/appointment.service";

export function usePatientAppointmentsQuery() {
  return useQuery({
    queryKey: ["patient-appointments"],
    queryFn: getPatientAppointments,
  });
}

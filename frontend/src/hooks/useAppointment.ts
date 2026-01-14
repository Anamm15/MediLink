import { useQuery } from "@tanstack/react-query";

import { getAppointment } from "@/services/appointment.service";

export function useAppointment(id: string) {
  return useQuery({
    queryKey: ["appointment", id],
    queryFn: () => getAppointment(id),
    staleTime: 5 * 60 * 1000,
  });
}

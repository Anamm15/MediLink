import { getAvailableSchedules } from "@/services/doctor.service";
import { useQuery } from "@tanstack/react-query";

export function useAvailableSchedulesQuery(
  doctor_id: string,
  date: string,
  day: string
) {
  return useQuery({
    queryKey: ["schedule", doctor_id, day, date],
    queryFn: () => getAvailableSchedules(doctor_id, date, day),
    staleTime: 3 * 60 * 1000,
  });
}

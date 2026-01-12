import { getAvailableSchedules } from "@/services/doctor.service";
import { useQuery } from "@tanstack/react-query";

export function useAvailableSchedulesQuery(
  doctor_id: string,
  date: string,
  day: string
) {
  return useQuery({
    queryKey: ["schedule", doctor_id],
    queryFn: () => getAvailableSchedules(doctor_id, date, day),
  });
}

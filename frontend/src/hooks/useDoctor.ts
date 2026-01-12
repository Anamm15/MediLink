import {
  getDoctor,
  getScheduleByID,
  searchDoctor,
} from "@/services/doctor.service";
import { useQuery } from "@tanstack/react-query";

export function useSearchDoctor(name: string, page: number) {
  return useQuery({
    queryKey: ["doctors", name, page],
    queryFn: () => searchDoctor(name, page),
  });
}

export function useDoctorQuery(id: string) {
  return useQuery({
    queryKey: ["doctor", id],
    queryFn: () => getDoctor(id),
  });
}

export function useScheduleQuery(id: string) {
  return useQuery({
    queryKey: ["schedule", id],
    queryFn: () => getScheduleByID(id),
  });
}

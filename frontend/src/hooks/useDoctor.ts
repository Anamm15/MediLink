import {
  doctorMe,
  getDoctor,
  getScheduleByID,
  searchDoctor,
} from "@/services/doctor.service";
import { useQuery } from "@tanstack/react-query";

export function useDoctorIdQuery() {
  return useQuery({
    queryKey: ["doctor-id"],
    queryFn: async () => {
      const data = await doctorMe();
      return data.id;
    },
    staleTime: Infinity,
  });
}

export function useSearchDoctor(name: string, page: number) {
  return useQuery({
    queryKey: ["doctors", name, page],
    queryFn: () => searchDoctor(name, page),
    staleTime: 60 * 1000,
  });
}

export function useDoctorQuery(id: string) {
  return useQuery({
    queryKey: ["doctor", id],
    queryFn: () => getDoctor(id),
    staleTime: 5 * 60 * 1000,
  });
}

export function useScheduleQuery(id: string) {
  return useQuery({
    queryKey: ["schedule", id],
    queryFn: () => getScheduleByID(id),
    staleTime: 5 * 60 * 1000,
  });
}

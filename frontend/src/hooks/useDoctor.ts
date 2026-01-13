import {
  doctorMe,
  getDoctor,
  getScheduleByID,
  searchDoctor,
} from "@/services/doctor.service";
import { useQuery } from "@tanstack/react-query";

export function useDoctorIdQuery() {
  return useQuery({
    queryKey: ["doctor-me"],
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

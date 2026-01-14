import { patientMe } from "@/services/patient.service";
import { useQuery } from "@tanstack/react-query";

export function usePatientIdQuery() {
  return useQuery({
    queryKey: ["patient-id"],
    queryFn: async () => {
      const data = await patientMe();
      return data.id;
    },
    staleTime: Infinity,
  });
}

import { getPatientPrescriptions } from "@/services/prescription.service";
import { useQuery } from "@tanstack/react-query";

export function usePatientPrescriptionQuery(id: string) {
  return useQuery({
    queryKey: ["patient-prescriptions"],
    queryFn: () => getPatientPrescriptions(id),
    enabled: !!id,
  });
}

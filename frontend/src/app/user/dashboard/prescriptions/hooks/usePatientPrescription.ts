import { getPatientPrescriptions } from "@/services/prescription.service";
import { useQuery } from "@tanstack/react-query";

export function usePatientPrescriptionQuery(
  id: string,
  page: number,
  limit: number,
  isRedeemed?: string
) {
  return useQuery({
    queryKey: ["patient-prescriptions", page, isRedeemed],
    queryFn: () => getPatientPrescriptions(id, page, limit, isRedeemed),
    enabled: !!id,
    staleTime: 3 * 60 * 1000,
  });
}

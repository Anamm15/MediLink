import { getDoctorPrescriptions } from "@/services/prescription.service";
import { useQuery } from "@tanstack/react-query";

export function useDoctorPrescriptionQuery(id: string) {
  return useQuery({
    queryKey: ["doctor-prescriptions", id],
    queryFn: () => getDoctorPrescriptions(id),
    enabled: !!id,
  });
}

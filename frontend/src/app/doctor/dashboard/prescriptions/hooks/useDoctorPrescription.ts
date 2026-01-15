import { getDoctorPrescriptions } from "@/services/prescription.service";
import { useQuery } from "@tanstack/react-query";

export function useDoctorPrescriptionQuery(
  doctor_id: string,
  page: number,
  limit: number
) {
  return useQuery({
    queryKey: ["doctor-prescriptions", doctor_id, page],
    queryFn: () => getDoctorPrescriptions(doctor_id, page, limit),
    enabled: !!doctor_id,
    staleTime: 3 * 60 * 1000,
  });
}

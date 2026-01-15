import { getDoctorPrescriptions } from "@/services/prescription.service";
import { useQuery } from "@tanstack/react-query";

export function useDoctorPrescriptionQuery(doctor_id: string) {
  return useQuery({
    queryKey: ["doctor-prescriptions", doctor_id],
    queryFn: () => getDoctorPrescriptions(doctor_id),
    enabled: !!doctor_id,
  });
}

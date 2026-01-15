import { getDoctorMedicalRecords } from "@/services/medical_record.service";
import { useQuery } from "@tanstack/react-query";

export function useDoctorMedicalRecord(
  doctor_id: string,
  page: number,
  limit: number
) {
  return useQuery({
    queryKey: ["doctor-medical-records", doctor_id, page],
    queryFn: () => getDoctorMedicalRecords(doctor_id, page, limit),
    enabled: !!doctor_id,
    staleTime: 3 * 60 * 1000,
  });
}

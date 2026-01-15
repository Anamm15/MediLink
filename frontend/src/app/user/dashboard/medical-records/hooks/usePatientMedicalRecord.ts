import { getPatientMedicalRecords } from "@/services/medical_record.service";
import { useQuery } from "@tanstack/react-query";

export function usePatientMedicalRecord(
  patient_id: string,
  page: number,
  limit: number
) {
  return useQuery({
    queryKey: ["patient-medical-records", page],
    queryFn: () => getPatientMedicalRecords(patient_id, page, limit),
    enabled: !!patient_id,
    staleTime: 3 * 60 * 1000,
  });
}

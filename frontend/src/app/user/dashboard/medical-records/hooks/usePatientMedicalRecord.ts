import { getPatientMedicalRecords } from "@/services/medical_record.service";
import { useQuery } from "@tanstack/react-query";

export function usePatientMedicalRecord(patient_id: string) {
  return useQuery({
    queryKey: ["patient-medical-records"],
    queryFn: () => getPatientMedicalRecords(patient_id),
    enabled: !!patient_id,
  });
}

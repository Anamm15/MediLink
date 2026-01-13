import { getDoctorMedicalRecords } from "@/services/medical_record.service";
import { useQuery } from "@tanstack/react-query";

export function useDoctorMedicalRecord(doctor_id: string) {
  return useQuery({
    queryKey: ["doctor-medical-records"],
    queryFn: () => getDoctorMedicalRecords(doctor_id),
    enabled: !!doctor_id,
  });
}

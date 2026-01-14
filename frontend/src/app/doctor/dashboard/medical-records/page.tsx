"use client";

import { useDoctorIdQuery } from "@/hooks/useDoctor";
import { MedicalRecordItem } from "./components/MedicalRecordItem";
import { useDoctorMedicalRecord } from "./hooks/useDoctorMedicalRecord";
import { useState } from "react";
import PrescriptionCreateModal from "./components/PrescriptionModal";

export default function MedicalRecordsPage() {
  const { data: doctorId } = useDoctorIdQuery();
  const { data: records } = useDoctorMedicalRecord(doctorId!);
  const [isPrescriptionModalOpen, setIsPrescriptionModalOpen] = useState(false);
  const [selectedField, setSelectedField] = useState({
    patient_id: "",
    medical_record_id: "",
  });

  return (
    <div className="space-y-6">
      <header className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <div>
          <h1 className="text-3xl font-bold text-gray-800">
            My Medical Record
          </h1>
          <p className="mt-1 text-gray-500">
            Your health journey timeline and medical documents.
          </p>
        </div>
      </header>

      <div className="relative">
        <div className="absolute left-6 top-0 h-full w-0.5 bg-gray-200"></div>
        <div className="space-y-8">
          {records &&
            records.map((record) => (
              <MedicalRecordItem
                key={record.id}
                record={record}
                setSelectedField={setSelectedField}
                setIsPrescriptionModalOpen={setIsPrescriptionModalOpen}
              />
            ))}
        </div>
      </div>

      <div className="text-center text-sm text-gray-400 pt-8">
        <p>Your medical record data is encrypted and stored securely.</p>
      </div>

      {isPrescriptionModalOpen && (
        <PrescriptionCreateModal
          isOpen={true}
          setIsOpen={setIsPrescriptionModalOpen}
          patient_id={selectedField.patient_id}
          medical_record_id={selectedField.medical_record_id}
        />
      )}
    </div>
  );
}

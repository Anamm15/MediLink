"use client";

import { useDoctorIdQuery } from "@/hooks/useDoctor";
import { MedicalRecordItem } from "./components/MedicalRecordItem";
import { useDoctorMedicalRecord } from "./hooks/useDoctorMedicalRecord";

export default function MedicalRecordsPage() {
  const { data: patientId } = useDoctorIdQuery();
  const { data: records } = useDoctorMedicalRecord(patientId!);

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

      {/* Filter (Opsional, bisa ditambahkan di sini) */}

      {/* Linimasa / Timeline */}
      <div className="relative">
        {/* Garis Vertikal Linimasa */}
        <div className="absolute left-6 top-0 h-full w-0.5 bg-gray-200"></div>

        <div className="space-y-8">
          {records &&
            records.map((record) => (
              <MedicalRecordItem key={record.id} record={record} />
            ))}
        </div>
      </div>

      <div className="text-center text-sm text-gray-400 pt-8">
        <p>Your medical record data is encrypted and stored securely.</p>
      </div>
    </div>
  );
}

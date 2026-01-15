"use client";

import { usePatientIdQuery } from "@/hooks/usePatient";
import { MedicalRecordItem } from "./components/MedicalRecordItem";
import { usePatientMedicalRecord } from "./hooks/usePatientMedicalRecord";
import { DEFAULT_LIMIT_QUERY, DEFAULT_PAGE_QUERY } from "@/helpers/constant";
import { useEffect, useState } from "react";
import { MedicalRecordResponse } from "@/types/medical_record.type";
import { DefaultPagination } from "@/components/ui/pagination/DefaultPagination";

export default function MedicalRecordsPage() {
  const { data: patientId } = usePatientIdQuery();
  const [page, setPage] = useState(DEFAULT_PAGE_QUERY);
  const [records, setRecords] = useState<MedicalRecordResponse[]>([]);
  const [totalPages, setTotalPages] = useState(0);
  const { data: recordsWithMetadata } = usePatientMedicalRecord(
    patientId!,
    page,
    DEFAULT_LIMIT_QUERY
  );

  useEffect(() => {
    if (!recordsWithMetadata) return;
    const { data: records, metadata } = recordsWithMetadata;
    setTotalPages(metadata.total_pages);
    setRecords(records);
  }, [recordsWithMetadata]);

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
              <MedicalRecordItem key={record.id} record={record} />
            ))}
        </div>

        {totalPages > 1 && (
          <DefaultPagination
            page={page}
            totalPages={totalPages}
            onPageChange={setPage}
          />
        )}
      </div>

      <div className="text-center text-sm text-gray-400 pt-8">
        <p>Your medical record data is encrypted and stored securely.</p>
      </div>
    </div>
  );
}

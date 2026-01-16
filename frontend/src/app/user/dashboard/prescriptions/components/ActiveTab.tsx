"use client";

import { DEFAULT_LIMIT_QUERY, DEFAULT_PAGE_QUERY } from "@/helpers/constant";
import { useState } from "react";
import { usePatientPrescriptionQuery } from "../hooks/usePatientPrescription";
import { DefaultPagination } from "@/components/ui/pagination/DefaultPagination";
import PrescriptionCard from "./PrescriptionCard";
import { Spinner } from "@/components/ui/Spinner";

interface ActiveTabProps {
  patientId: string;
}

export default function ActiveTab({ patientId }: ActiveTabProps) {
  const [page, setPage] = useState(DEFAULT_PAGE_QUERY);
  const { data, isLoading } = usePatientPrescriptionQuery(
    patientId!,
    page,
    DEFAULT_LIMIT_QUERY,
    "false"
  );

  const prescriptions = data?.data ?? [];
  const metadata = data?.metadata;

  if (isLoading) {
    return (
      <div className="flex h-full w-full justify-center items-center mt-40">
        <Spinner />
      </div>
    );
  }

  return (
    <div className="space-y-4">
      {prescriptions &&
        prescriptions.map((prescription) => (
          <PrescriptionCard key={prescription.id} prescription={prescription} />
        ))}

      {prescriptions.length === 0 && (
        <div className="text-center py-12">
          <p className="text-gray-500">
            You don not have any active prescription.
          </p>
          <p className="text-sm text-gray-400 mt-1">
            New prescription will be appear here after consultation.
          </p>
        </div>
      )}

      {metadata && metadata.total_pages > 1 && prescriptions.length > 0 && (
        <DefaultPagination
          page={page}
          onPageChange={setPage}
          totalPages={metadata.total_pages}
        />
      )}
    </div>
  );
}

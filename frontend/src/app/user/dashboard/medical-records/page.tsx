"use client";

import { usePatientIdQuery } from "@/hooks/usePatient";
import { MedicalRecordItem } from "./components/MedicalRecordItem";
import { usePatientMedicalRecord } from "./hooks/usePatientMedicalRecord";
import { DEFAULT_LIMIT_QUERY, DEFAULT_PAGE_QUERY } from "@/helpers/constant";
import { useState } from "react";
import { DefaultPagination } from "@/components/ui/pagination/DefaultPagination";
import { Spinner } from "@/components/ui/Spinner";
import { AnimatePresence, motion } from "framer-motion";

const listVariants = {
  hidden: { opacity: 0 },
  visible: { opacity: 1, transition: { staggerChildren: 0.15 } },
};

export default function MedicalRecordsPage() {
  const { data: patientId } = usePatientIdQuery();
  const [page, setPage] = useState(DEFAULT_PAGE_QUERY);
  const { data, isLoading } = usePatientMedicalRecord(
    patientId!,
    page,
    DEFAULT_LIMIT_QUERY
  );

  const records = data?.data ?? [];
  const metadata = data?.metadata;

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

      {isLoading ? (
        <div className="flex h-full w-full justify-center items-center mt-40">
          <Spinner />
        </div>
      ) : (
        <AnimatePresence>
          <motion.div
            initial="hidden"
            animate="visible"
            exit={{ opacity: 0 }}
            variants={listVariants}
            className="space-y-4"
          >
            {records &&
              records.map((item) => (
                <MedicalRecordItem key={item.id} record={item} />
              ))}

            {metadata && metadata.total_pages > 1 && (
              <DefaultPagination
                page={page}
                totalPages={metadata.total_pages}
                onPageChange={setPage}
                siblingCount={3}
              />
            )}
          </motion.div>
        </AnimatePresence>
      )}

      <div className="text-center text-sm text-gray-400 pt-8">
        <p>Your medical record data is encrypted and stored securely.</p>
      </div>
    </div>
  );
}

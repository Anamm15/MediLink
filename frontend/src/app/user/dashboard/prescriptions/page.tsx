"use client";
import { useEffect, useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { usePatientIdQuery } from "@/hooks/usePatient";
import { usePatientPrescriptionQuery } from "./hooks/usePatientPrescription";
import { PrescriptionResponse } from "@/types/prescription.type";
import PrescriptionCard from "./components/PrescriptionCard";
import { DefaultPagination } from "@/components/ui/pagination/DefaultPagination";
import { DEFAULT_LIMIT_QUERY, DEFAULT_PAGE_QUERY } from "@/helpers/constant";

const listVariants = {
  hidden: { opacity: 0 },
  visible: { opacity: 1, transition: { staggerChildren: 0.15 } },
};

export default function UserPrescriptionPage() {
  const [activeTab, setActiveTab] = useState("active");
  const [page, setPage] = useState(DEFAULT_PAGE_QUERY);
  const [totalPages, setTotalPages] = useState(0);
  const { data: patientId } = usePatientIdQuery();
  const { data: prescriptionsWithMetadata } = usePatientPrescriptionQuery(
    patientId!,
    page,
    DEFAULT_LIMIT_QUERY
  );
  const [redeemedPrescriptions, setRedeemedPrescriptions] = useState<
    PrescriptionResponse[]
  >([]);
  const [unredeemedPrescriptions, setUnredeemedPrescriptions] = useState<
    PrescriptionResponse[]
  >([]);

  useEffect(() => {
    if (!prescriptionsWithMetadata) return;
    const { data: prescriptions, metadata } = prescriptionsWithMetadata;
    setTotalPages(metadata.total_pages);

    if (prescriptions) {
      const redeemed = prescriptions.filter(
        (prescription) => prescription.is_redeemed === true
      );
      const unredeemed = prescriptions.filter(
        (prescription) => prescription.is_redeemed === false
      );
      setRedeemedPrescriptions(redeemed);
      setUnredeemedPrescriptions(unredeemed);
    }
  }, [prescriptionsWithMetadata]);

  return (
    <div className="space-y-6">
      <header>
        <h1 className="text-3xl font-bold text-gray-800">My Prescriptions</h1>
        <p className="mt-1 text-gray-500">
          View and redeem prescriptions given by your doctor.
        </p>
      </header>

      {/* Tombol Tab */}
      <div className="flex border-b border-gray-200">
        <button
          onClick={() => setActiveTab("active")}
          className={`px-4 py-2 text-sm font-semibold transition-colors ${
            activeTab === "active"
              ? "border-b-2 border-cyan-500 text-cyan-600"
              : "text-gray-500 hover:text-gray-800"
          }`}
        >
          Active ({unredeemedPrescriptions.length})
        </button>
        <button
          onClick={() => setActiveTab("history")}
          className={`px-4 py-2 text-sm font-semibold transition-colors ${
            activeTab === "history"
              ? "border-b-2 border-cyan-500 text-cyan-600"
              : "text-gray-500 hover:text-gray-800"
          }`}
        >
          History ({redeemedPrescriptions.length})
        </button>
      </div>

      {/* Konten Resep */}
      <AnimatePresence mode="wait">
        <motion.div
          key={activeTab}
          initial="hidden"
          animate="visible"
          exit={{ opacity: 0 }}
          variants={listVariants}
          className="space-y-4"
        >
          {(activeTab === "active"
            ? unredeemedPrescriptions
            : redeemedPrescriptions
          ).map((rx) => (
            <PrescriptionCard key={rx.id} prescription={rx} />
          ))}

          {activeTab === "active" && unredeemedPrescriptions.length === 0 && (
            <div className="text-center py-12">
              <p className="text-gray-500">
                You don not have any active prescription.
              </p>
              <p className="text-sm text-gray-400 mt-1">
                New prescription will be appear here after consultation.
              </p>
            </div>
          )}

          {activeTab === "history" && redeemedPrescriptions.length === 0 && (
            <div className="text-center py-12">
              <p className="text-gray-500">
                You don not have any prescription history yet.
              </p>
            </div>
          )}

          {activeTab === "active" &&
            totalPages > 1 &&
            unredeemedPrescriptions.length > 0 && (
              <DefaultPagination
                page={page}
                onPageChange={setPage}
                totalPages={totalPages}
              />
            )}

          {activeTab === "history" &&
            totalPages > 1 &&
            redeemedPrescriptions.length > 0 && (
              <DefaultPagination
                page={page}
                onPageChange={setPage}
                totalPages={totalPages}
              />
            )}
        </motion.div>
      </AnimatePresence>
    </div>
  );
}

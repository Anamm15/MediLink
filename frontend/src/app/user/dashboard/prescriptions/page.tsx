"use client";
import { useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { usePatientIdQuery } from "@/hooks/usePatient";
import ActiveTab from "./components/ActiveTab";
import HistoryTab from "./components/HistoryTab";

const listVariants = {
  hidden: { opacity: 0 },
  visible: { opacity: 1, transition: { staggerChildren: 0.15 } },
};

export default function UserPrescriptionPage() {
  const [activeTab, setActiveTab] = useState("active");
  const { data: patientId } = usePatientIdQuery();

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
          Active
        </button>

        <button
          onClick={() => setActiveTab("history")}
          className={`px-4 py-2 text-sm font-semibold transition-colors ${
            activeTab === "history"
              ? "border-b-2 border-cyan-500 text-cyan-600"
              : "text-gray-500 hover:text-gray-800"
          }`}
        >
          History
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
          {activeTab === "active" ? (
            <ActiveTab patientId={patientId!} />
          ) : (
            <HistoryTab patientId={patientId!} />
          )}
        </motion.div>
      </AnimatePresence>
    </div>
  );
}

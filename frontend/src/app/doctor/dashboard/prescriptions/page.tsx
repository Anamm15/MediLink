"use client";
import { useState } from "react";
import { Search } from "lucide-react";
import { CreatePrescriptionModal } from "./components/CreatePrescriptionModal";
import { AnimatePresence, motion } from "framer-motion";
import { Input } from "@/components/ui/form/Input";
import { useDoctorIdQuery } from "@/hooks/useDoctor";
import { useDoctorPrescriptionQuery } from "./hooks/useDoctorPrescription";
import PrescriptionCard from "./components/PrescriptionCard";

const listVariants = {
  hidden: { opacity: 0 },
  visible: { opacity: 1, transition: { staggerChildren: 0.15 } },
};

export default function DoctorPrescriptionPage() {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [search, setSearch] = useState("");
  const { data: doctorId } = useDoctorIdQuery();
  const { data: prescriptions } = useDoctorPrescriptionQuery(doctorId!);

  return (
    <>
      <div className="flex flex-col h-full">
        <header className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 mb-6">
          <div>
            <h1 className="text-3xl font-bold text-gray-800">
              Prescription Management
            </h1>
            <p className="mt-1 text-gray-500">
              Create or view digital prescription history for your patients.
            </p>
          </div>
        </header>

        <Input
          type="text"
          className="w-full rounded-lg mb-6 py-6 text-md"
          placeholder="Search for prescriptions by patient name or ID..."
          value={search}
          startIcon={<Search className="w-5 h-5 text-gray-400" />}
          onChange={(e) => setSearch(e.target.value)}
        />

        <AnimatePresence>
          <motion.div
            initial="hidden"
            animate="visible"
            exit={{ opacity: 0 }}
            variants={listVariants}
            className="space-y-4"
          >
            {prescriptions &&
              prescriptions.map((item) => (
                <PrescriptionCard key={item.id} prescription={item} />
              ))}
          </motion.div>
        </AnimatePresence>
      </div>

      <AnimatePresence>
        {isModalOpen && (
          <CreatePrescriptionModal onClose={() => setIsModalOpen(false)} />
        )}
      </AnimatePresence>
    </>
  );
}

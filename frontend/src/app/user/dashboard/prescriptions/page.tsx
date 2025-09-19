"use client";
import { useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { UserPrescriptionCard } from "@/components/cards/UserPrescriptionCard";

// --- DATA DUMMY ---
const mockActivePrescriptions = [
  {
    id: "RX-USER-01",
    doctor: {
      name: "Dr. Adinda Melati, Sp.A",
      specialty: "Dokter Anak",
      avatarUrl: "https://i.pravatar.cc/150?u=adinda",
    },
    date: "19 Sep 2025",
    status: "Aktif",
    medicines: [
      { name: "Paracetamol Sirup", quantity: "1 botol" },
      { name: "Cetirizine 10mg", quantity: "1 strip" },
    ],
  },
];

const mockPastPrescriptions = [
  {
    id: "RX-USER-02",
    doctor: {
      name: "Dr. Citra Ayu, Sp.KK",
      specialty: "Kulit & Kelamin",
      avatarUrl: "https://i.pravatar.cc/150?u=citra",
    },
    date: "15 Agustus 2025",
    status: "Sudah Ditebus",
    medicines: [{ name: "Hydrocortisone Cream 1%", quantity: "1 tube" }],
  },
  {
    id: "RX-USER-03",
    doctor: {
      name: "Dr. Budi Santoso, Sp.PD",
      specialty: "Penyakit Dalam",
      avatarUrl: "https://i.pravatar.cc/150?u=budi",
    },
    date: "10 Juli 2025",
    status: "Sudah Ditebus",
    medicines: [
      { name: "Amoxicillin 500mg", quantity: "15 tablet" },
      { name: "Ibuprofen 400mg", quantity: "10 tablet" },
    ],
  },
];

const listVariants = {
  hidden: { opacity: 0 },
  visible: { opacity: 1, transition: { staggerChildren: 0.15 } },
};

export default function UserPrescriptionPage() {
  const [activeTab, setActiveTab] = useState("active");

  return (
    <div className="space-y-6">
      <header>
        <h1 className="text-3xl font-bold text-gray-800">Resep Digital Saya</h1>
        <p className="mt-1 text-gray-500">
          Lihat dan tebus resep yang diberikan oleh dokter Anda.
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
          Resep Aktif ({mockActivePrescriptions.length})
        </button>
        <button
          onClick={() => setActiveTab("history")}
          className={`px-4 py-2 text-sm font-semibold transition-colors ${
            activeTab === "history"
              ? "border-b-2 border-cyan-500 text-cyan-600"
              : "text-gray-500 hover:text-gray-800"
          }`}
        >
          Riwayat ({mockPastPrescriptions.length})
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
            ? mockActivePrescriptions
            : mockPastPrescriptions
          ).map((rx) => (
            <UserPrescriptionCard key={rx.id} prescription={rx} />
          ))}

          {activeTab === "active" && mockActivePrescriptions.length === 0 && (
            <div className="text-center py-12">
              <p className="text-gray-500">Anda tidak memiliki resep aktif.</p>
              <p className="text-sm text-gray-400 mt-1">
                Resep baru akan muncul di sini setelah konsultasi.
              </p>
            </div>
          )}

          {activeTab === "history" && mockPastPrescriptions.length === 0 && (
            <div className="text-center py-12">
              <p className="text-gray-500">
                Anda belum memiliki riwayat resep.
              </p>
            </div>
          )}
        </motion.div>
      </AnimatePresence>
    </div>
  );
}

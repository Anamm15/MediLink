"use client";
import { useState } from "react";
import { Plus, Search } from "lucide-react";
import { PrescriptionListItem } from "./components/PrescriptionListItem";
import { CreatePrescriptionModal } from "./components/CreatePrescriptionModal";
import { motion, AnimatePresence } from "framer-motion";

// --- DATA DUMMY ---
const mockHistory = [
  {
    id: "RX-001",
    patientName: "Budi Setiawan",
    date: "19 Sep 2025",
    status: "Ditebus",
  },
  {
    id: "RX-002",
    patientName: "Sarah L.",
    date: "18 Sep 2025",
    status: "Ditebus",
  },
  {
    id: "RX-003",
    patientName: "Rian D.",
    date: "17 Sep 2025",
    status: "Belum Ditebus",
  },
  {
    id: "RX-004",
    patientName: "Anita P.",
    date: "16 Sep 2025",
    status: "Ditebus",
  },
];

export default function DoctorPrescriptionPage() {
  const [isModalOpen, setIsModalOpen] = useState(false);

  return (
    <>
      <div className="flex flex-col h-full">
        {/* Header Halaman */}
        <header className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 mb-6">
          <div>
            <h1 className="text-3xl font-bold text-gray-800">
              Manajemen Resep
            </h1>
            <p className="mt-1 text-gray-500">
              Buat atau lihat riwayat resep digital untuk pasien Anda.
            </p>
          </div>
          <button
            onClick={() => setIsModalOpen(true)}
            className="flex items-center justify-center gap-2 px-5 py-2.5 font-semibold text-white bg-slate-800 rounded-lg shadow-sm hover:bg-slate-700"
          >
            <Plus className="w-5 h-5" />
            Buat Resep Baru
          </button>
        </header>

        {/* Search Bar */}
        <div className="mb-6 relative">
          <input
            type="text"
            placeholder="Cari resep berdasarkan nama pasien atau ID..."
            className="w-full p-3 pl-10 border rounded-lg"
          />
          <Search className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-gray-400" />
        </div>

        {/* Daftar Riwayat Resep */}
        <div className="space-y-3">
          {mockHistory.map((item) => (
            <PrescriptionListItem key={item.id} prescription={item} />
          ))}
        </div>
      </div>

      {/* Modal Pembuatan Resep */}
      <AnimatePresence>
        {isModalOpen && (
          <CreatePrescriptionModal onClose={() => setIsModalOpen(false)} />
        )}
      </AnimatePresence>
    </>
  );
}

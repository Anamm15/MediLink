"use client";
import { useState } from "react";
import { Plus } from "lucide-react";
import { motion, AnimatePresence } from "framer-motion";
import { UserAppointmentCard } from "@/components/cards/UserAppointmentCard";
import Link from "next/link";

// --- DATA DUMMY ---
const mockUpcomingAppointments = [
  {
    id: "apt-user-01",
    doctor: {
      name: "Dr. Adinda Melati, Sp.A",
      specialty: "Dokter Anak",
      avatarUrl: "https://i.pravatar.cc/150?u=adinda",
    },
    date: `Besok, ${new Date(
      new Date().setDate(new Date().getDate() + 1)
    ).toLocaleDateString("id-ID", { day: "numeric", month: "long" })}`,
    time: "10:00",
    type: "Online",
    status: "Dikonfirmasi",
  },
  {
    id: "apt-user-02",
    doctor: {
      name: "Dr. Budi Santoso, Sp.PD",
      specialty: "Penyakit Dalam",
      avatarUrl: "https://i.pravatar.cc/150?u=budi",
      clinicAddress: "RS Harapan Kita, Lt. 3",
    },
    date: `Lusa, ${new Date(
      new Date().setDate(new Date().getDate() + 2)
    ).toLocaleDateString("id-ID", { day: "numeric", month: "long" })}`,
    time: "14:00",
    type: "Onsite",
    status: "Dikonfirmasi",
  },
];

const mockPastAppointments = [
  {
    id: "apt-user-03",
    doctor: {
      name: "Dr. Citra Ayu, Sp.KK",
      specialty: "Kulit & Kelamin",
      avatarUrl: "https://i.pravatar.cc/150?u=citra",
    },
    date: `15 September 2025`,
    time: "11:00",
    type: "Online",
    status: "Selesai",
  },
];

const listVariants = {
  hidden: { opacity: 0 },
  visible: { opacity: 1, transition: { staggerChildren: 0.1 } },
};

export default function UserAppointmentPage() {
  const [activeTab, setActiveTab] = useState("upcoming");

  return (
    <div className="space-y-6">
      <header className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <div>
          <h1 className="text-3xl font-bold text-gray-800">Janji Temu Saya</h1>
          <p className="mt-1 text-gray-500">
            Lihat dan kelola semua jadwal konsultasi Anda.
          </p>
        </div>
        <Link
          href="/doctors"
          className="flex items-center justify-center gap-2 px-5 py-2.5 font-semibold text-white bg-cyan-500 rounded-lg shadow-sm hover:bg-cyan-600"
        >
          <Plus className="w-5 h-5" />
          Booking Janji Baru
        </Link>
      </header>

      {/* Tombol Tab */}
      <div className="flex border-b border-gray-200">
        <button
          onClick={() => setActiveTab("upcoming")}
          className={`px-4 py-2 text-sm font-semibold transition-colors ${
            activeTab === "upcoming"
              ? "border-b-2 border-cyan-500 text-cyan-600"
              : "text-gray-500 hover:text-gray-800"
          }`}
        >
          Akan Datang ({mockUpcomingAppointments.length})
        </button>
        <button
          onClick={() => setActiveTab("history")}
          className={`px-4 py-2 text-sm font-semibold transition-colors ${
            activeTab === "history"
              ? "border-b-2 border-cyan-500 text-cyan-600"
              : "text-gray-500 hover:text-gray-800"
          }`}
        >
          Riwayat ({mockPastAppointments.length})
        </button>
      </div>

      {/* Konten Janji Temu */}
      <AnimatePresence mode="wait">
        <motion.div
          key={activeTab}
          initial="hidden"
          animate="visible"
          exit={{ opacity: 0 }}
          variants={listVariants}
          className="space-y-4"
        >
          {(activeTab === "upcoming"
            ? mockUpcomingAppointments
            : mockPastAppointments
          ).map((app) => (
            <UserAppointmentCard key={app.id} appointment={app} />
          ))}

          {activeTab === "upcoming" &&
            mockUpcomingAppointments.length === 0 && (
              <p className="text-center text-gray-500 pt-8">
                Anda tidak memiliki janji temu yang akan datang.
              </p>
            )}

          {activeTab === "history" && mockPastAppointments.length === 0 && (
            <p className="text-center text-gray-500 pt-8">
              Anda belum memiliki riwayat konsultasi.
            </p>
          )}
        </motion.div>
      </AnimatePresence>
    </div>
  );
}

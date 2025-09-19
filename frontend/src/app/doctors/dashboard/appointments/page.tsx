"use client";
import { useState } from "react";
import { AppointmentList } from "./components/AppointmentList";
import { motion, AnimatePresence } from "framer-motion";

// --- DATA DUMMY ---
// Fungsi helper untuk mendapatkan tanggal dinamis
const getFormattedDate = (offset: number) => {
  const date = new Date();
  date.setDate(date.getDate() + offset);
  return date.toLocaleDateString("id-ID", {
    weekday: "long",
    day: "numeric",
    month: "long",
  });
};

const DATES = {
  TODAY: `Hari Ini, ${getFormattedDate(0)}`,
  TOMORROW: `Besok, ${getFormattedDate(1)}`,
  PAST: getFormattedDate(-1),
};

const mockUpcomingAppointments = [
  {
    id: "apt-001",
    date: DATES.TODAY,
    time: "09:00",
    type: "Video Call",
    complaint: "Demam dan batuk kering selama 2 hari.",
    patient: {
      name: "Budi Setiawan",
      age: 34,
      avatarUrl: "https://i.pravatar.cc/150?u=budi",
    },
    isUpcoming: true,
    isActionable: true,
  },
  {
    id: "apt-002",
    date: DATES.TODAY,
    time: "10:00",
    type: "Chat",
    complaint: "Konsultasi hasil lab darah.",
    patient: {
      name: "Sarah L.",
      age: 28,
      avatarUrl: "https://i.pravatar.cc/150?u=sarah",
    },
    isUpcoming: true,
    isActionable: false,
  },
  {
    id: "apt-003",
    date: DATES.TODAY,
    time: "11:00",
    type: "Video Call",
    complaint: "Ruam merah pada kulit lengan.",
    patient: {
      name: "Anita P.",
      age: 22,
      avatarUrl: "https://i.pravatar.cc/150?u=anita",
    },
    isUpcoming: true,
    isActionable: false,
  },
  {
    id: "apt-004",
    date: DATES.TOMORROW,
    time: "14:00",
    type: "Video Call",
    complaint: "Kontrol pasca-pemulihan.",
    patient: {
      name: "Rian D.",
      age: 45,
      avatarUrl: "https://i.pravatar.cc/150?u=rian",
    },
    isUpcoming: true,
    isActionable: false,
  },
];

const mockPastAppointments = [
  {
    id: "apt-005",
    date: DATES.PAST,
    time: "10:00",
    type: "Chat",
    complaint: "Konsultasi jerawat.",
    patient: {
      name: "Citra Ayu",
      age: 19,
      avatarUrl: "https://i.pravatar.cc/150?u=citra",
    },
    isUpcoming: false,
  },
];

export default function DoctorAppointmentPage() {
  const [activeTab, setActiveTab] = useState("upcoming");

  return (
    <div className="flex flex-col h-full">
      <header className="mb-6">
        <h1 className="text-3xl font-bold text-gray-800">Janji Temu Pasien</h1>
        <p className="mt-1 text-gray-500">
          Kelola jadwal konsultasi Anda dengan efisien.
        </p>
      </header>

      {/* Tombol Tab */}
      <div className="flex border-b border-gray-200 mb-6">
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

      {/* Konten Tab */}
      <AnimatePresence mode="wait">
        <motion.div
          key={activeTab}
          initial={{ opacity: 0, y: 10 }}
          animate={{ opacity: 1, y: 0 }}
          exit={{ opacity: 0, y: -10 }}
          transition={{ duration: 0.2 }}
        >
          {activeTab === "upcoming" ? (
            <AppointmentList appointments={mockUpcomingAppointments} />
          ) : (
            <AppointmentList appointments={mockPastAppointments} />
          )}
        </motion.div>
      </AnimatePresence>
    </div>
  );
}

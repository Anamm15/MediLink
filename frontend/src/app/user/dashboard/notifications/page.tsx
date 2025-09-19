"use client";
import { useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { NotificationItem } from "./components/NotificationItem";
import { CheckCheck } from "lucide-react";

// --- DATA DUMMY ---
const mockNotifications = [
  {
    id: 1,
    type: "appointment_reminder",
    title: "Pengingat Janji Temu",
    description:
      "Anda memiliki jadwal konsultasi dengan Dr. Adinda Melati besok pukul 10:00.",
    timestamp: "5 menit lalu",
    isRead: false,
    action: { text: "Lihat Detail Janji Temu", href: "/appointment" },
  },
  {
    id: 2,
    type: "prescription_ready",
    title: "Resep Baru Telah Terbit",
    description: "Dr. Adinda Melati telah menerbitkan resep baru untuk Anda.",
    timestamp: "1 jam lalu",
    isRead: false,
    action: { text: "Lihat & Tebus Resep", href: "/prescriptions" },
  },
  {
    id: 3,
    type: "payment_success",
    title: "Pembayaran Berhasil",
    description:
      "Pembayaran sebesar Rp 150.000 untuk konsultasi telah berhasil.",
    timestamp: "3 jam lalu",
    isRead: true,
    action: { text: "Lihat Invoice", href: "/payment-history" },
  },
  {
    id: 4,
    type: "promo",
    title: "Diskon Spesial Untuk Anda!",
    description:
      "Dapatkan diskon 30% untuk konsultasi dengan dokter spesialis kulit minggu ini.",
    timestamp: "1 hari lalu",
    isRead: true,
    action: { text: "Lihat Promo", href: "#" },
  },
  {
    id: 5,
    type: "appointment_reminder",
    title: "Janji Temu Dibatalkan",
    description:
      "Jadwal konsultasi Anda dengan Dr. Budi Santoso telah dibatalkan.",
    timestamp: "2 hari lalu",
    isRead: true,
  },
];

const importantTypes = ["appointment_reminder", "prescription_ready"];
const infoTypes = ["payment_success", "promo"];

export default function NotificationsPage() {
  const [activeTab, setActiveTab] = useState("all");

  const filteredNotifications = mockNotifications.filter((n) => {
    if (activeTab === "important") return importantTypes.includes(n.type);
    if (activeTab === "info") return infoTypes.includes(n.type);
    return true; // 'all'
  });

  return (
    <div className="space-y-6">
      <header className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <div>
          <h1 className="text-3xl font-bold text-gray-800">Notifikasi</h1>
          <p className="mt-1 text-gray-500">
            Semua pembaruan penting Anda ada di sini.
          </p>
        </div>
        <button className="flex items-center justify-center gap-2 px-4 py-2 text-sm font-semibold text-gray-600 bg-white border rounded-lg shadow-sm hover:bg-gray-100">
          <CheckCheck className="w-4 h-4" /> Tandai semua sudah dibaca
        </button>
      </header>

      {/* Tombol Tab */}
      <div className="flex border-b border-gray-200">
        <button
          onClick={() => setActiveTab("all")}
          className={`px-4 py-2 text-sm font-semibold transition-colors ${
            activeTab === "all"
              ? "border-b-2 border-cyan-500 text-cyan-600"
              : "text-gray-500 hover:text-gray-800"
          }`}
        >
          Semua
        </button>
        <button
          onClick={() => setActiveTab("important")}
          className={`px-4 py-2 text-sm font-semibold transition-colors ${
            activeTab === "important"
              ? "border-b-2 border-cyan-500 text-cyan-600"
              : "text-gray-500 hover:text-gray-800"
          }`}
        >
          Penting
        </button>
        <button
          onClick={() => setActiveTab("info")}
          className={`px-4 py-2 text-sm font-semibold transition-colors ${
            activeTab === "info"
              ? "border-b-2 border-cyan-500 text-cyan-600"
              : "text-gray-500 hover:text-gray-800"
          }`}
        >
          Info & Promosi
        </button>
      </div>

      {/* Konten Notifikasi */}
      <AnimatePresence>
        <motion.div
          key={activeTab}
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ duration: 0.3 }}
          className="space-y-3"
        >
          {filteredNotifications.length > 0 ? (
            filteredNotifications.map((n) => (
              <NotificationItem key={n.id} notification={n} />
            ))
          ) : (
            <div className="text-center py-16">
              <p className="font-semibold text-gray-700">
                Tidak ada notifikasi di kategori ini
              </p>
              <p className="text-sm text-gray-400 mt-1">
                Semua notifikasi Anda yang relevan akan muncul di sini.
              </p>
            </div>
          )}
        </motion.div>
      </AnimatePresence>
    </div>
  );
}

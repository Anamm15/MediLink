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
    title: "Appointment Reminder",
    description:
      "You have a scheduled consultation with Dr. Adinda Melati tomorrow at 10:00 AM.",
    timestamp: "5 minutes ago",
    isRead: false,
    action: { text: "View Appointment Details", href: "/appointment" },
  },
  {
    id: 2,
    type: "prescription_ready",
    title: "New Prescription Issued",
    description: "Dr. Adinda Melati has issued a new prescription for you.",
    timestamp: "1 hour ago",
    isRead: false,
    action: { text: "View & Redeem Prescription", href: "/prescriptions" },
  },
  {
    id: 3,
    type: "payment_success",
    title: "Payment Successful",
    description:
      "A payment of IDR 150,000 for the consultation was completed successfully.",
    timestamp: "3 hours ago",
    isRead: true,
    action: { text: "View Invoice", href: "/payment-history" },
  },
  {
    id: 4,
    type: "promo",
    title: "Special Discount Just for You!",
    description:
      "Enjoy a 30% discount on consultations with dermatology specialists this week.",
    timestamp: "1 day ago",
    isRead: true,
    action: { text: "View Promotion", href: "#" },
  },
  {
    id: 5,
    type: "appointment_reminder",
    title: "Appointment Cancelled",
    description:
      "Your consultation appointment with Dr. Budi Santoso has been cancelled.",
    timestamp: "2 days ago",
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
          <h1 className="text-3xl font-bold text-gray-800">Notification</h1>
          <p className="mt-1 text-gray-500">
            All your important updates are here.
          </p>
        </div>
        <button className="flex items-center justify-center gap-2 px-4 py-2 text-sm font-semibold text-gray-600 bg-white border rounded-lg shadow-sm hover:bg-gray-100">
          <CheckCheck className="w-4 h-4" /> Mark all as read
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
          All
        </button>
        <button
          onClick={() => setActiveTab("important")}
          className={`px-4 py-2 text-sm font-semibold transition-colors ${
            activeTab === "important"
              ? "border-b-2 border-cyan-500 text-cyan-600"
              : "text-gray-500 hover:text-gray-800"
          }`}
        >
          Important
        </button>
        <button
          onClick={() => setActiveTab("info")}
          className={`px-4 py-2 text-sm font-semibold transition-colors ${
            activeTab === "info"
              ? "border-b-2 border-cyan-500 text-cyan-600"
              : "text-gray-500 hover:text-gray-800"
          }`}
        >
          Info & Promotions
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
                There are no notifications in this category
              </p>
              <p className="text-sm text-gray-400 mt-1">
                All your relevant notifications will appear here.
              </p>
            </div>
          )}
        </motion.div>
      </AnimatePresence>
    </div>
  );
}

"use client";
import { useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { AppointmentResponse } from "@/types/appointment.type";
import { useDoctorIdQuery } from "@/hooks/useDoctor";
import UpcomingTab from "./components/UpcomingTab";
import HistoryTab from "./components/HistoryTab";

export const groupAppointmentsByDate = (
  appointments: AppointmentResponse[]
) => {
  const groups = appointments.reduce((acc, app) => {
    (acc[app.appointment_date] = acc[app.appointment_date] || []).push(app);
    return acc;
  }, {} as Record<string, AppointmentResponse[]>);
  return groups;
};

export default function DoctorAppointmentPage() {
  const [activeTab, setActiveTab] = useState("upcoming");
  const { data: doctorId } = useDoctorIdQuery();

  return (
    <div className="flex flex-col h-full">
      <header className="mb-6">
        <h1 className="text-3xl font-bold text-gray-800">
          Patient Appointment
        </h1>
        <p className="mt-1 text-gray-500">
          Manage your consultation schedule efficiently.
        </p>
      </header>

      <div className="flex border-b border-gray-200 mb-6">
        <button
          onClick={() => setActiveTab("upcoming")}
          className={`px-4 py-2 text-sm font-semibold transition-colors ${
            activeTab === "upcoming"
              ? "border-b-2 border-cyan-500 text-cyan-600"
              : "text-gray-500 hover:text-gray-800"
          }`}
        >
          Upcoming
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

      <AnimatePresence mode="wait">
        <motion.div
          key={activeTab}
          initial={{ opacity: 0, y: 10 }}
          animate={{ opacity: 1, y: 0 }}
          exit={{ opacity: 0, y: -10 }}
          transition={{ duration: 0.2 }}
        >
          {activeTab === "upcoming" ? (
            <UpcomingTab doctorId={doctorId!} />
          ) : (
            <HistoryTab doctorId={doctorId!} />
          )}
        </motion.div>
      </AnimatePresence>
    </div>
  );
}

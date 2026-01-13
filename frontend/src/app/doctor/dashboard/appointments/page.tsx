"use client";
import { useEffect, useState } from "react";
import { AppointmentList } from "./components/AppointmentList";
import { motion, AnimatePresence } from "framer-motion";
import { useDoctorAppointmentsQuery } from "./hooks/useDoctorAppointment";
import { AppointmentDetailResponse } from "@/types/appointment.type";
import { groupAppointmentsByTime } from "@/helpers/appointments";

export default function DoctorAppointmentPage() {
  const [activeTab, setActiveTab] = useState("upcoming");
  const { data: appointments } = useDoctorAppointmentsQuery();
  const [upComingAppointments, setUpComingAppointments] = useState<
    AppointmentDetailResponse[]
  >([]);
  const [pastAppointments, setPastAppointments] = useState<
    AppointmentDetailResponse[]
  >([]);

  useEffect(() => {
    if (!appointments) return;
    const { upcoming, past } = groupAppointmentsByTime(appointments);
    setUpComingAppointments(upcoming);
    setPastAppointments(past);
  }, [appointments]);

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
          Upcoming ({upComingAppointments.length})
        </button>
        <button
          onClick={() => setActiveTab("history")}
          className={`px-4 py-2 text-sm font-semibold transition-colors ${
            activeTab === "history"
              ? "border-b-2 border-cyan-500 text-cyan-600"
              : "text-gray-500 hover:text-gray-800"
          }`}
        >
          History ({pastAppointments.length})
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
            <AppointmentList appointments={upComingAppointments} />
          ) : (
            <AppointmentList appointments={pastAppointments} />
          )}
        </motion.div>
      </AnimatePresence>
    </div>
  );
}

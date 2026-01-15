"use client";
import { useEffect, useState } from "react";
import { Plus } from "lucide-react";
import { motion, AnimatePresence } from "framer-motion";
import { UserAppointmentCard } from "./components/UserAppointmentCard";
import Link from "next/link";
import { usePatientAppointmentsQuery } from "./hooks/usePatientAppointment";
import { AppointmentResponse } from "@/types/appointment.type";
import { groupAppointmentsByTime } from "@/helpers/appointments";
import { DefaultPagination } from "@/components/ui/pagination/DefaultPagination";
import { Metadata } from "@/types/metadata.type";
import { DEFAULT_LIMIT_QUERY } from "@/helpers/constant";

const listVariants = {
  hidden: { opacity: 0 },
  visible: { opacity: 1, transition: { staggerChildren: 0.1 } },
};

export default function UserAppointmentPage() {
  const [activeTab, setActiveTab] = useState("upcoming");
  const [page, setPage] = useState(1);
  const { data: appointmentsWithMetadata } = usePatientAppointmentsQuery(
    page,
    DEFAULT_LIMIT_QUERY
  );
  const [totalPages, setTotalPages] = useState(0);
  const [upComingAppointments, setUpComingAppointments] = useState<
    AppointmentResponse[]
  >([]);
  const [pastAppointments, setPastAppointments] = useState<
    AppointmentResponse[]
  >([]);

  useEffect(() => {
    if (!appointmentsWithMetadata) return;
    const { data: appointments, metadata } = appointmentsWithMetadata;
    setTotalPages(metadata.total_pages);

    const { upcoming, past } = groupAppointmentsByTime(appointments);
    setUpComingAppointments(upcoming);
    setPastAppointments(past);
  }, [appointmentsWithMetadata]);

  return (
    <div className="space-y-6">
      <header className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <div>
          <h1 className="text-3xl font-bold text-gray-800">My Appointment</h1>
          <p className="mt-1 text-gray-500">
            View and manage all your consultation schedules.
          </p>
        </div>
        <Link
          href="/doctor"
          className="flex items-center justify-center gap-2 px-5 py-2.5 font-semibold text-white bg-cyan-500 rounded-lg shadow-sm hover:bg-cyan-600"
        >
          <Plus className="w-5 h-5" />
          Book New Appointment
        </Link>
      </header>

      <div className="flex border-b border-gray-200">
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
          initial="hidden"
          animate="visible"
          exit={{ opacity: 0 }}
          variants={listVariants}
          className="space-y-4"
        >
          {(activeTab === "upcoming"
            ? upComingAppointments
            : pastAppointments
          ).map((app) => (
            <UserAppointmentCard key={app.id} appointment={app} />
          ))}

          {activeTab === "upcoming" && upComingAppointments.length === 0 && (
            <p className="text-center text-gray-500 pt-8">
              You don't have any upcoming appointments.
            </p>
          )}

          {activeTab === "history" && pastAppointments.length === 0 && (
            <p className="text-center text-gray-500 pt-8">
              You have no consultation history yet.
            </p>
          )}

          {activeTab === "upcoming" &&
            totalPages > 1 &&
            upComingAppointments.length > 0 && (
              <DefaultPagination
                page={page}
                onPageChange={setPage}
                totalPages={totalPages}
              />
            )}

          {activeTab === "history" &&
            totalPages > 1 &&
            pastAppointments.length > 0 && (
              <DefaultPagination
                page={page}
                onPageChange={setPage}
                totalPages={totalPages}
              />
            )}
        </motion.div>
      </AnimatePresence>
    </div>
  );
}

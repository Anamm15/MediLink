"use client";
import Image from "next/image";
import { motion } from "framer-motion";
import {
  Calendar,
  Clock,
  Video,
  Building,
  MoreVertical,
  FileText,
  Repeat,
} from "lucide-react";
import { AppointmentResponse } from "@/types/appointment.type";
import { DEFAULT_PROFILE } from "@/helpers/constant";
import Link from "next/link";
import { formatIDDate } from "@/helpers/datetime";

interface CardProps {
  appointment: AppointmentResponse;
}

export const UserAppointmentCard = ({ appointment }: CardProps) => {
  const { doctor, appointment_date, start_time, end_time, type, status } =
    appointment;
  const isUpcoming = status === "Dikonfirmasi";
  const TypeIcon = type === "Online" ? Video : Building;

  const cardVariants = {
    hidden: { opacity: 0, y: 20 },
    visible: { opacity: 1, y: 0, transition: { duration: 0.5 } },
  };

  return (
    <motion.div
      variants={cardVariants}
      className="bg-white rounded-xl shadow-sm border border-gray-200 flex flex-col sm:flex-row hover:border-cyan-200"
    >
      {/* Kiri: Info Dokter */}
      <div className="flex-shrink-0 w-full sm:w-1/3 md:w-1/4 p-4 flex items-center gap-4 sm:flex-col sm:justify-center sm:text-center sm:border-r">
        <Image
          src={doctor.avatar_url || DEFAULT_PROFILE}
          alt={doctor.name}
          width={72}
          height={72}
          className="rounded-full"
        />
        <div>
          <h3 className="font-bold text-gray-800">{doctor.name}</h3>
          <p className="text-sm text-gray-500">{doctor.specialization}</p>
        </div>
      </div>

      {/* Tengah: Detail Jadwal */}
      <div className="flex-grow p-4 border-t sm:border-t-0">
        <div className="flex items-center gap-4 text-sm text-gray-600">
          <span className="flex items-center gap-1.5">
            <Calendar className="w-4 h-4 text-gray-400" />{" "}
            {formatIDDate(appointment_date)}
          </span>
          <span className="flex items-center gap-1.5">
            <Clock className="w-4 h-4 text-gray-400" /> {start_time} -{" "}
            {end_time} WIB
          </span>
        </div>
        <div className="mt-3 flex items-center gap-2">
          <TypeIcon className="w-5 h-5 text-cyan-600" />
          <p className="font-semibold text-gray-800">Consultation {type}</p>
        </div>
        {/* {type === "Onsite" && (
          <p className="text-xs text-gray-500 mt-1">{doctor.clinicAddress}</p>
        )} */}
      </div>

      {/* Kanan: Aksi */}
      <div className="flex-shrink-0 w-full sm:w-auto p-4 flex items-center justify-end gap-2 bg-slate-50 border-t sm:border-t-0 sm:border-l">
        {isUpcoming ? (
          <>
            <button className="px-4 py-2 text-sm font-semibold text-white bg-cyan-500 rounded-lg hover:bg-cyan-600">
              {type === "Online" ? "Join Consultation" : "View Detail Location"}
            </button>
            <button className="p-2 text-gray-500 rounded-md hover:bg-gray-200">
              <MoreVertical className="w-5 h-5" />
            </button>
          </>
        ) : (
          <>
            <button className="px-4 py-2 text-sm font-semibold text-gray-700 bg-gray-200 rounded-lg hover:bg-gray-300 flex items-center gap-2">
              <FileText className="w-4 h-4" /> View Prescription
            </button>
            <Link
              href={`/user/dashboard/appointments/${appointment.id}`}
              className="px-4 py-2 text-sm font-semibold text-white bg-slate-800 rounded-lg hover:bg-slate-700 flex items-center gap-2"
            >
              <Repeat className="w-4 h-4" /> View Detail
            </Link>
          </>
        )}
      </div>
    </motion.div>
  );
};

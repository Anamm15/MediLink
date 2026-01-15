"use client";

import { motion } from "framer-motion";
import Image from "next/image";
import Link from "next/link";
import {
  Calendar,
  Video,
  MapPin,
  ChevronLeft,
  ClipboardList,
  Stethoscope,
  CreditCard,
  Link as LinkIcon,
  Info,
} from "lucide-react";
import { AppointmentDetailResponse } from "@/types/appointment.type";
import { statusConfig } from "@/helpers/badge";
import { DEFAULT_PROFILE } from "@/helpers/constant";
import { Button } from "@/components/ui/Button";
import {
  useCancelAppointment,
  useCompleteAppointment,
} from "./hooks/useStatusAppointment";
import { useState } from "react";

type AppointmentDetailProps = {
  data: AppointmentDetailResponse;
  urlBack: string;
  isDoctor?: boolean;
};

export default function AppointmentDetail({
  data,
  urlBack,
  isDoctor = false,
}: AppointmentDetailProps) {
  const [appointment, setAppointment] = useState(data);
  const status =
    statusConfig[appointment.status as string] || statusConfig.PENDING;
  const StatusIcon = status.icon;
  const { mutate: cancelAppointment } = useCancelAppointment(
    appointment,
    setAppointment
  );
  const { mutate: completeAppointment } = useCompleteAppointment(
    appointment,
    setAppointment
  );

  return (
    <motion.div
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      className="max-w-5xl mx-auto space-y-6 pb-12"
    >
      <Link
        href={urlBack}
        className="inline-flex items-center gap-2 text-sm font-semibold text-gray-500 hover:text-cyan-600 transition-colors"
      >
        <ChevronLeft className="w-4 h-4" /> Back to Appoinments
      </Link>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <div className="lg:col-span-2 space-y-6">
          <div className="bg-white rounded-2xl border border-gray-200 shadow-sm overflow-hidden">
            <div
              className={`p-4 flex items-center justify-between border-b ${status.color}`}
            >
              <div className="flex items-center gap-2">
                <StatusIcon className="w-5 h-5" />
                <span className="font-bold text-sm uppercase tracking-wider">
                  {status.label}
                </span>
              </div>
              <span className="text-xs font-mono font-medium">
                ID: {appointment.id}
              </span>
            </div>

            <div className="p-8 space-y-8">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div className="space-y-3">
                  <h3 className="text-sm font-bold text-gray-400 uppercase flex items-center gap-2">
                    <Calendar className="w-4 h-4" /> Date
                  </h3>
                  <div>
                    <p className="text-xl font-bold text-gray-800">
                      {appointment.appointment_date}
                    </p>
                    <p className="text-gray-500 font-medium">
                      {appointment.start_time} - {appointment.end_time} WIB
                    </p>
                  </div>
                </div>
                <div className="space-y-3">
                  <h3 className="text-sm font-bold text-gray-400 uppercase flex items-center gap-2">
                    <Info className="w-4 h-4" /> Queue Number
                  </h3>
                  <div className="bg-slate-50 w-fit px-6 py-2 rounded-xl border border-slate-200">
                    <p className="text-2xl font-black text-slate-800">
                      #{appointment.queue_number}
                    </p>
                  </div>
                </div>
              </div>

              {/* Sesi 2: Keluhan & Catatan */}
              <div className="space-y-6 pt-6 border-t border-gray-100">
                <div className="space-y-3">
                  <h3 className="text-sm font-bold text-gray-400 uppercase flex items-center gap-2">
                    <ClipboardList className="w-4 h-4" /> Complaint
                  </h3>
                  <p className="text-gray-700 leading-relaxed bg-gray-50 p-4 rounded-xl italic">
                    {appointment.symptom_complaint || "No Complaint"}
                  </p>
                </div>

                {appointment.doctor_notes && (
                  <div className="space-y-3">
                    <h3 className="text-sm font-bold text-gray-400 uppercase flex items-center gap-2">
                      <Stethoscope className="w-4 h-4" /> Doctor Notes
                    </h3>
                    <div className="p-4 rounded-xl border-l-4 border-cyan-500 bg-cyan-50/30">
                      <p className="text-gray-700 whitespace-pre-line">
                        {appointment.doctor_notes}
                      </p>
                    </div>
                  </div>
                )}
              </div>
            </div>
          </div>

          <div className="bg-white rounded-2xl border border-gray-200 p-6 shadow-sm">
            <h3 className="text-sm font-bold text-gray-400 uppercase mb-4">
              Patient Profile
            </h3>
            <div className="flex items-center gap-4">
              <Image
                src={appointment.patient.avatar_url || DEFAULT_PROFILE}
                alt="avatar"
                width={56}
                height={56}
                className="rounded-full bg-slate-100"
              />
              <div>
                <p className="font-bold text-gray-800">
                  {appointment.patient.name}
                </p>
                <p className="text-sm text-gray-500">
                  {appointment.patient.email} •{" "}
                  {appointment.patient.phone_number}
                </p>
              </div>
            </div>
          </div>
        </div>

        <div className="space-y-6">
          <div className="bg-white rounded-2xl border border-gray-200 p-6 shadow-sm text-center">
            <Image
              src={appointment.doctor.avatar_url || DEFAULT_PROFILE}
              alt="doctor"
              width={80}
              height={80}
              className="rounded-full mx-auto border-4 border-cyan-50"
            />
            <h3 className="mt-4 font-bold text-lg text-gray-800">
              {appointment.doctor.name}
            </h3>
            <p className="text-cyan-600 font-medium text-sm">
              {appointment.doctor.specialization}
            </p>
            <div className="mt-6 pt-6 border-t border-gray-100">
              <p className="text-xs text-gray-400 uppercase font-bold mb-1">
                Consultation Type
              </p>
              <div className="flex items-center justify-center gap-2 text-gray-700 font-semibold">
                {appointment.type === "video_call" ||
                appointment.type === "chat" ? (
                  <>
                    <Video className="w-4 h-4 text-purple-500" /> Video Call
                    (Online)
                  </>
                ) : (
                  <>
                    <MapPin className="w-4 h-4 text-orange-500" /> Onsite
                  </>
                )}
              </div>
            </div>
            {isDoctor && appointment.status === "PENDING" && (
              <div className="grid grid-cols-1 md:grid-cols-2 mt-4 gap-2">
                <Button
                  onClick={() => cancelAppointment(appointment.id)}
                  variant="secondary"
                  className="cursor-pointer rounded-xl"
                >
                  Cancel
                </Button>
                <Button
                  onClick={() => completeAppointment(appointment.id)}
                  variant="primary"
                  className="cursor-pointer rounded-xl"
                >
                  Complete
                </Button>
              </div>
            )}
          </div>

          {appointment.status === "CONFIRMED" &&
            appointment.type === "ONLINE_VIDEO" && (
              <div className="bg-slate-800 rounded-2xl p-6 text-white shadow-xl shadow-slate-200">
                <h3 className="font-bold flex items-center gap-2 mb-4">
                  <LinkIcon className="w-5 h-5" /> Meeting Link
                </h3>
                <p className="text-sm text-slate-400 mb-6">
                  The consultation will begin online via video conferencing.
                </p>
                <a
                  href={appointment.meeting_link}
                  target="_blank"
                  className="block w-full py-3 bg-cyan-500 hover:bg-cyan-600 text-center font-bold rounded-xl transition-all"
                >
                  Join Now
                </a>
              </div>
            )}

          {/* Financial Snapshot */}
          <div className="bg-white rounded-2xl border border-gray-200 p-6 shadow-sm">
            <h3 className="text-sm font-bold text-gray-400 uppercase mb-4 flex items-center gap-2">
              <CreditCard className="w-4 h-4" /> Bill Details
            </h3>
            <div className="flex justify-between items-center">
              <span className="text-gray-500">Cost</span>
              <span className="font-bold text-gray-800 text-lg">
                Rp{" "}
                {appointment.consultation_fee_snapshot.toLocaleString("id-ID")}
              </span>
            </div>
            <div className="mt-4 pt-4 border-t border-dashed flex items-center gap-2 text-[10px] text-gray-400">
              <Info className="w-3 h-3" />
              <p>This price is a fixed fee at the time the order is placed.</p>
            </div>
          </div>
        </div>
      </div>
    </motion.div>
  );
}

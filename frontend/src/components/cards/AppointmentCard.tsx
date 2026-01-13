import Image from "next/image";
import { Video, MessageSquare, Clock } from "lucide-react";
import { AppointmentDetailResponse } from "@/types/appointment.type";
import Link from "next/link";
import { useState } from "react";
import MedicalRecordModal from "@/app/doctor/dashboard/appointments/components/MedicalRecordModal";

interface AppointmentCardProps {
  appointment: AppointmentDetailResponse;
  isUpcoming: boolean;
  isActionable: boolean;
}

export const AppointmentCard = ({
  appointment,
  isUpcoming,
  isActionable,
}: AppointmentCardProps) => {
  const { patient, start_time, type, symptom_complaint } = appointment;
  const TypeIcon = type === "Video Call" ? Video : MessageSquare;
  const [isMedicalRecordModalOpen, setIsMedicalRecordModalOpen] =
    useState(false);

  return (
    <div className="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden flex transition-all hover:shadow-lg hover:border-cyan-300">
      {/* Left Column */}
      <div className="w-1/3 md:w-1/4 p-4 flex flex-col items-center justify-center text-center border-r">
        <Image
          src={"https://i.pravatar.cc/150?u=citra"}
          alt={patient.name}
          width={72}
          height={72}
          className="rounded-full"
        />
        <h3 className="mt-3 font-bold text-gray-800">{patient.name}</h3>
        {/* <p className="text-sm text-gray-500">{patient.age} tahun</p> */}
      </div>

      {/* Center Column */}
      <div className="w-2/3 md:w-1/2 p-4 flex flex-col justify-center">
        <div className="flex items-center gap-4">
          <p className="text-2xl font-bold text-gray-800 flex items-center gap-2">
            <Clock className="w-6 h-6 text-gray-400" /> {start_time}
          </p>
          <span
            className={`flex items-center gap-1.5 text-xs font-semibold px-2.5 py-1 rounded-full ${
              type === "Video Call"
                ? "bg-purple-100 text-purple-700"
                : "bg-cyan-100 text-cyan-700"
            }`}
          >
            <TypeIcon className="w-3 h-3" /> {type}
          </span>
        </div>
        <p className="mt-3 text-sm text-gray-600">
          <span className="font-semibold">Complaint:</span> {symptom_complaint}
        </p>
      </div>

      {/* Right Column */}
      <div className="hidden md:flex w-1/4 p-4 items-center justify-end bg-slate-50 border-l">
        {isUpcoming ? (
          <div className="flex flex-col items-end gap-2 w-full">
            <button
              disabled={!isActionable}
              className={`w-full font-semibold py-2.5 rounded-lg text-sm transition-colors ${
                isActionable
                  ? "bg-cyan-500 text-white hover:bg-cyan-600"
                  : "bg-gray-200 text-gray-500 cursor-not-allowed"
              }`}
            >
              Start Consultation
            </button>
          </div>
        ) : (
          <div className="flex flex-col items-end gap-2 w-full">
            <Link
              href={`/doctor/dashboard/appointments/${appointment.id}`}
              className="w-full font-semibold py-2.5 rounded-lg text-sm bg-cyan-500 text-white hover:bg-cyan-600 transition-colors text-center"
            >
              View Detail
            </Link>
            <button
              onClick={() => setIsMedicalRecordModalOpen(true)}
              className="text-xs font-semibold text-gray-500 hover:text-gray-800"
            >
              Write Medical Notes
            </button>
          </div>
        )}
      </div>

      {isMedicalRecordModalOpen && (
        <MedicalRecordModal
          isOpen={isMedicalRecordModalOpen}
          setIsOpen={setIsMedicalRecordModalOpen}
          patient_id={patient.id}
          appointment_id={appointment.id}
        />
      )}
    </div>
  );
};

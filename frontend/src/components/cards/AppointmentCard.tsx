import Image from "next/image";
import {
  Video,
  MessageSquare,
  Clock,
  FileText,
  MoreHorizontal,
} from "lucide-react";

// Tipe data untuk satu janji temu
type Appointment = {
  id: string;
  patient: {
    name: string;
    age: number;
    avatarUrl: string;
  };
  time: string;
  type: "Video Call" | "Chat";
  complaint: string;
  isUpcoming: boolean;
  isActionable?: boolean; // Untuk menandakan tombol 'Mulai' aktif
};

interface AppointmentCardProps {
  appointment: Appointment;
}

export const AppointmentCard = ({ appointment }: AppointmentCardProps) => {
  const { patient, time, type, complaint, isUpcoming, isActionable } =
    appointment;
  const TypeIcon = type === "Video Call" ? Video : MessageSquare;

  return (
    <div className="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden flex transition-all hover:shadow-lg hover:border-cyan-300">
      {/* Kolom Kiri: Info Pasien */}
      <div className="w-1/3 md:w-1/4 p-4 flex flex-col items-center justify-center text-center border-r">
        <Image
          src={patient.avatarUrl}
          alt={patient.name}
          width={72}
          height={72}
          className="rounded-full"
        />
        <h3 className="mt-3 font-bold text-gray-800">{patient.name}</h3>
        <p className="text-sm text-gray-500">{patient.age} tahun</p>
      </div>

      {/* Kolom Tengah: Detail Sesi */}
      <div className="w-2/3 md:w-1/2 p-4 flex flex-col justify-center">
        <div className="flex items-center gap-4">
          <p className="text-2xl font-bold text-gray-800 flex items-center gap-2">
            <Clock className="w-6 h-6 text-gray-400" /> {time}
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
          <span className="font-semibold">Keluhan:</span> {complaint}
        </p>
      </div>

      {/* Kolom Kanan: Aksi */}
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
              Mulai Konsultasi
            </button>
            <button className="text-xs font-semibold text-gray-500 hover:text-gray-800">
              Lihat Rekam Medis
            </button>
          </div>
        ) : (
          <div className="flex flex-col items-end gap-2 w-full">
            <button className="w-full font-semibold py-2.5 rounded-lg text-sm bg-slate-800 text-white hover:bg-slate-700">
              Lihat Detail
            </button>
            <button className="text-xs font-semibold text-gray-500 hover:text-gray-800">
              Tulis Catatan Medis
            </button>
          </div>
        )}
      </div>
    </div>
  );
};

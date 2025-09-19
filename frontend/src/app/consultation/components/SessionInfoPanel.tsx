"use client";
import { useState, useEffect } from "react";
import Image from "next/image";
import { Clock, FileText } from "lucide-react";
import { Doctor } from "@/types"; // Gunakan tipe Doctor yang sudah ada

interface SessionInfoPanelProps {
  doctor: Doctor;
  durationInMinutes: number;
}

export const SessionInfoPanel = ({
  doctor,
  durationInMinutes,
}: SessionInfoPanelProps) => {
  const [timeLeft, setTimeLeft] = useState(durationInMinutes * 60);

  useEffect(() => {
    if (timeLeft <= 0) return;
    const timerId = setInterval(() => setTimeLeft(timeLeft - 1), 1000);
    return () => clearInterval(timerId);
  }, [timeLeft]);

  const minutes = Math.floor(timeLeft / 60);
  const seconds = timeLeft % 60;
  const progress = (timeLeft / (durationInMinutes * 60)) * 100;

  return (
    <aside className="w-80 flex-shrink-0 bg-white border-l border-gray-200 p-6 flex flex-col">
      <h2 className="text-lg font-bold text-gray-800 mb-4">Detail Sesi</h2>

      {/* Doctor Info */}
      <div className="flex items-center gap-4 p-4 bg-slate-50 rounded-lg">
        <Image
          src={doctor.avatarUrl}
          alt={doctor.name}
          width={56}
          height={56}
          className="rounded-full"
        />
        <div>
          <h3 className="font-bold text-gray-900">{doctor.name}</h3>
          <p className="text-sm text-gray-500">{doctor.specialty}</p>
        </div>
      </div>

      {/* Session Timer */}
      <div className="mt-8">
        <h3 className="font-semibold text-gray-700 mb-3 flex items-center">
          <Clock className="w-5 h-5 mr-2 text-gray-400" /> Sisa Waktu
        </h3>
        <div className="w-full bg-gray-200 rounded-full h-2.5">
          <div
            className="bg-cyan-500 h-2.5 rounded-full"
            style={{ width: `${progress}%` }}
          ></div>
        </div>
        <p className="text-center text-3xl font-bold text-gray-800 mt-3">
          {String(minutes).padStart(2, "0")}:{String(seconds).padStart(2, "0")}
        </p>
      </div>

      <div className="mt-auto space-y-3">
        <button className="w-full flex items-center justify-center gap-2 py-2.5 px-4 text-sm font-semibold rounded-lg border border-cyan-500 text-cyan-500 hover:bg-cyan-50">
          <FileText className="w-4 h-4" /> Minta Ringkasan Medis
        </button>
        <button className="w-full py-2.5 px-4 text-sm font-semibold rounded-lg bg-red-500 text-white hover:bg-red-600">
          Akhiri Konsultasi
        </button>
      </div>
    </aside>
  );
};

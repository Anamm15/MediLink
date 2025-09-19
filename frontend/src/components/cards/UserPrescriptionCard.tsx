"use client";
import Image from "next/image";
import { motion } from "framer-motion";
import { User, Calendar, Pill, ShoppingCart, Eye } from "lucide-react";

// Tipe data untuk resep dari perspektif user
type UserPrescription = any; // Ganti dengan tipe data spesifik Anda

interface CardProps {
  prescription: UserPrescription;
}

export const UserPrescriptionCard = ({ prescription }: CardProps) => {
  const { doctor, date, status, medicines } = prescription;
  const isActive = status === "Aktif";

  const cardVariants = {
    hidden: { opacity: 0, scale: 0.95 },
    visible: { opacity: 1, scale: 1, transition: { duration: 0.4 } },
  };

  return (
    <motion.div
      variants={cardVariants}
      className="bg-white rounded-xl shadow-sm border border-gray-200"
    >
      <div className="p-4 flex flex-col sm:flex-row sm:items-center sm:justify-between border-b">
        {/* Info Dokter & Tanggal */}
        <div className="flex items-center gap-4">
          <Image
            src={doctor.avatarUrl}
            alt={doctor.name}
            width={48}
            height={48}
            className="rounded-full"
          />
          <div>
            <p className="font-semibold text-gray-800">{doctor.name}</p>
            <p className="text-xs text-gray-500">{doctor.specialty}</p>
          </div>
        </div>
        <div className="flex items-center gap-2 text-sm text-gray-500 mt-3 sm:mt-0">
          <Calendar className="w-4 h-4" />
          <span>Diterbitkan: {date}</span>
        </div>
      </div>

      {/* Detail Obat */}
      <div className="p-4">
        <h4 className="text-sm font-semibold text-gray-700 mb-2 flex items-center gap-2">
          <Pill className="w-4 h-4" /> Detail Obat
        </h4>
        <div className="space-y-1 text-sm text-gray-600">
          {medicines.map((med: any, index: number) => (
            <p key={index}>
              • {med.name} ({med.quantity})
            </p>
          ))}
        </div>
      </div>

      {/* Aksi & Status */}
      <div className="p-4 flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3 bg-slate-50 rounded-b-xl">
        <span
          className={`text-xs font-bold px-3 py-1.5 rounded-full ${
            isActive
              ? "bg-yellow-100 text-yellow-800"
              : "bg-green-100 text-green-800"
          }`}
        >
          {isActive ? "Aktif - Belum Ditebus" : "Sudah Ditebus"}
        </span>
        {isActive ? (
          <button className="flex items-center justify-center gap-2 w-full sm:w-auto px-4 py-2 text-sm font-semibold text-white bg-cyan-500 rounded-lg hover:bg-cyan-600">
            <ShoppingCart className="w-4 h-4" /> Lihat & Tebus Resep
          </button>
        ) : (
          <button className="flex items-center justify-center gap-2 w-full sm:w-auto px-4 py-2 text-sm font-semibold text-gray-700 bg-gray-200 rounded-lg hover:bg-gray-300">
            <Eye className="w-4 h-4" /> Lihat Detail Riwayat
          </button>
        )}
      </div>
    </motion.div>
  );
};

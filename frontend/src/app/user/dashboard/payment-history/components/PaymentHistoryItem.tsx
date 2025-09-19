"use client";
import { motion } from "framer-motion";
import { Stethoscope, Pill, Download } from "lucide-react";

// Tipe data untuk transaksi
type Transaction = any;

interface ItemProps {
  transaction: Transaction;
}

const statusStyles = {
  Berhasil: "bg-green-100 text-green-800",
  Tertunda: "bg-yellow-100 text-yellow-800",
  Gagal: "bg-red-100 text-red-800",
};

const typeIcons = {
  Konsultasi: <Stethoscope className="w-6 h-6 text-cyan-600" />,
  Apotek: <Pill className="w-6 h-6 text-purple-600" />,
};

export const PaymentHistoryItem = ({ transaction }: ItemProps) => {
  const { id, type, description, date, amount, status } = transaction;

  return (
    <motion.div
      initial={{ opacity: 0, y: 10 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.3 }}
      className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 p-4 bg-white rounded-lg border hover:border-gray-300"
    >
      <div className="flex items-center gap-4">
        <div
          className={`p-3 rounded-full ${
            type === "Konsultasi" ? "bg-cyan-50" : "bg-purple-50"
          }`}
        >
          {typeIcons[type as keyof typeof typeIcons]}
        </div>
        <div>
          <p className="font-semibold text-gray-800">{description}</p>
          <p className="text-sm text-gray-500">
            {date} • ID: {id}
          </p>
        </div>
      </div>

      <div className="flex items-center justify-between sm:justify-end gap-4 sm:gap-6">
        <p className="font-bold text-gray-800">
          Rp {amount.toLocaleString("id-ID")}
        </p>
        <span
          className={`text-xs font-bold px-3 py-1.5 rounded-full ${
            statusStyles[status as keyof typeof statusStyles]
          }`}
        >
          {status}
        </span>
        <button className="flex items-center gap-2 px-3 py-2 text-xs font-semibold text-gray-600 bg-gray-100 rounded-md hover:bg-gray-200">
          <Download className="w-3 h-3" /> Invoice
        </button>
      </div>
    </motion.div>
  );
};

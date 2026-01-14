"use client";
import { motion } from "framer-motion";
import { PaymentHistoryHeader } from "./components/PaymentHistoryHeader";
import { PaymentHistoryItem } from "./components/PaymentHistoryItem";

// --- DATA DUMMY ---
const mockPaymentHistory = [
  {
    id: "TRX-001",
    type: "Consultation",
    description: "Consultation with Dr. Adinda Melati",
    date: "19 Sep 2025, 10:30",
    amount: 150000,
    status: "Success",
  },
  {
    id: "TRX-002",
    type: "Pharmacy",
    description: "Purchase #RX-USER-01",
    date: "19 Sep 2025, 11:00",
    amount: 48000,
    status: "Success",
  },
  {
    id: "TRX-003",
    type: "Consultation",
    description: "Consultation with Dr. Budi Santoso",
    date: "12 Sep 2025, 14:00",
    amount: 125000,
    status: "Success",
  },
  {
    id: "TRX-004",
    type: "Pharmacy",
    description: "Purchase Vitamin C",
    date: "05 Sep 2025, 09:15",
    amount: 75000,
    status: "Success",
  },
  {
    id: "TRX-005",
    type: "Consultation",
    description: "Consultation with Dr. Citra Ayu",
    date: "01 Sep 2025, 16:45",
    amount: 152000,
    status: "Failed",
  },
];

export default function PaymentHistoryPage() {
  return (
    <div className="space-y-6">
      <header>
        <h1 className="text-3xl font-bold text-gray-800">Payment History</h1>
        <p className="mt-1 text-gray-500">
          Track all transactions and download your invoices here.
        </p>
      </header>

      <PaymentHistoryHeader />

      {/* Daftar Transaksi */}
      <motion.div
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        transition={{ duration: 0.5 }}
        className="space-y-3"
      >
        <div className="hidden sm:flex text-xs font-semibold text-gray-500 px-4">
          <div className="w-1/2">DESCRIPTION</div>
          <div className="w-1/2 flex justify-end items-center gap-12 pr-28">
            <span>AMOUNT</span>
            <span>STATUS</span>
          </div>
        </div>
        {mockPaymentHistory.map((trx) => (
          <PaymentHistoryItem key={trx.id} transaction={trx} />
        ))}
      </motion.div>
    </div>
  );
}

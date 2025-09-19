"use client";
import Image from "next/image";
import { Lock } from "lucide-react";

interface PaymentFlowProps {
  consultationFee: number;
  onPay: () => void; // Fungsi yang akan dipanggil saat tombol bayar diklik
}

// Daftar logo untuk ditampilkan
const supportedPayments = [
  { name: "Visa", logo: "/payment/visa.svg" },
  { name: "Mastercard", logo: "/payment/mastercard.svg" },
  { name: "BCA", logo: "/payment/bca.svg" },
  { name: "Mandiri", logo: "/payment/mandiri.svg" },
  { name: "GoPay", logo: "/payment/gopay.svg" },
  { name: "QRIS", logo: "/payment/qris.svg" },
];

export const PaymentFlow = ({ consultationFee, onPay }: PaymentFlowProps) => {
  const adminFee = 2500;
  const total = consultationFee + adminFee;

  return (
    <div className="bg-white p-6 rounded-xl border border-gray-200 shadow-sm flex flex-col h-full">
      {/* 1. Rincian Tagihan */}
      <div>
        <h2 className="text-xl font-bold text-gray-800 mb-4">
          Rincian Tagihan
        </h2>
        <div className="space-y-3 text-sm">
          <div className="flex justify-between">
            <p className="text-gray-500">Biaya Konsultasi</p>
            <p className="font-medium text-gray-800">
              Rp {consultationFee.toLocaleString("id-ID")}
            </p>
          </div>
          <div className="flex justify-between">
            <p className="text-gray-500">Biaya Layanan & Admin</p>
            <p className="font-medium text-gray-800">
              Rp {adminFee.toLocaleString("id-ID")}
            </p>
          </div>
          <div className="border-t border-dashed my-3"></div>
          <div className="flex justify-between text-base">
            <p className="font-semibold text-gray-800">Total Pembayaran</p>
            <p className="font-bold text-cyan-600 text-lg">
              Rp {total.toLocaleString("id-ID")}
            </p>
          </div>
        </div>
      </div>

      {/* 2. Seksi Metode Pembayaran (Midtrans Gateway) */}
      <div className="mt-8">
        <h2 className="text-xl font-bold text-gray-800">Metode Pembayaran</h2>
        <div className="mt-4 bg-slate-50 border border-slate-200 rounded-lg p-4">
          <div className="flex items-center justify-between">
            <p className="text-sm font-semibold text-gray-700">
              Pembayaran aman melalui:
            </p>
            <Image
              src="/payment/midtrans.svg"
              alt="Midtrans Logo"
              width={100}
              height={25}
            />
          </div>
          <div className="mt-4 flex flex-wrap items-center gap-x-4 gap-y-2">
            {supportedPayments.map((p) => (
              <Image
                key={p.name}
                src={p.logo}
                alt={p.name}
                width={40}
                height={25}
                className="object-contain"
              />
            ))}
          </div>
        </div>
      </div>

      {/* 3. Tombol Aksi Pembayaran */}
      <div className="mt-auto pt-8">
        <button
          onClick={onPay}
          className="w-full bg-slate-800 text-white font-bold py-3.5 rounded-lg text-lg hover:bg-slate-700 transition-colors duration-200 flex items-center justify-center gap-2"
        >
          <Lock className="w-5 h-5" />
          Lanjutkan Pembayaran
        </button>
        <p className="text-xs text-gray-400 text-center mt-3">
          Anda akan diarahkan ke halaman pembayaran yang aman.
        </p>
      </div>
    </div>
  );
};

"use client"; // Jadikan Client Component untuk menangani state dan event
import Script from "next/script";
import { BookingSummaryCard } from "./components/BookingSummaryCard";
import { PaymentFlow } from "./components/PaymentFlow";
import { DoctorDetail } from "@/types/index.type";
import { User } from "lucide-react";

// Declare Midtrans Snap type for TypeScript
declare global {
  interface Window {
    snap: any;
  }
}

// Data mock (sama seperti sebelumnya)
const mockDoctorDetail: DoctorDetail = {
  id: "1",
  name: "Dr. Adinda Melati, Sp.A",
  specialty: "Dokter Anak",
  avatarUrl: "https://i.pravatar.cc/150?u=adinda",
  rating: 4.9,
  reviews: 128,
  isOnline: true,
  nextAvailable: "Besok, 10:00 WIB",
  yearsOfExperience: 12,
  patientCount: 5000,
  bio: "Seorang Dokter Spesialis Anak.",
  consultationFee: 150000,
  clinic: { name: "Klinik Sehat Ceria", address: "Jl. Merdeka No. 123" },
  education: [],
  experience: [],
  reviewsList: [],
};
const mockBookingDetails = {
  date: "Rabu, 20 September 2025",
  time: "09:00",
  type: "Video Call" as const,
};
const mockUser = { name: "Budi Setiawan", email: "budi.setiawan@email.com" };

export default function OrderPage() {
  const doctor = mockDoctorDetail;
  const booking = mockBookingDetails;

  const handlePayment = async () => {
    // LANGKAH 1: Minta token dari backend Anda
    // const response = await fetch('/api/create-transaction', { method: 'POST', body: JSON.stringify({ orderId: 'UNIQUE_ORDER_ID', amount: 152500 }) });
    // const { token } = await response.json();

    // Placeholder untuk token (ganti dengan token asli dari backend)
    const token = "GANTI_DENGAN_TOKEN_DARI_BACKEND_ANDA";
    console.log("Memulai pembayaran dengan token:", token);

    // LANGKAH 2: Panggil Midtrans Snap
    if (window.snap && token) {
      window.snap.pay(token, {
        onSuccess: function (result: any) {
          console.log("success", result);
          // Redirect ke halaman sukses
          window.location.href = "/booking/success";
        },
        onPending: function (result: any) {
          console.log("pending", result);
        },
        onError: function (result: any) {
          console.log("error", result);
        },
        onClose: function () {
          console.log(
            "customer closed the popup without finishing the payment"
          );
        },
      });
    }
  };

  return (
    <>
      {/* Load script Midtrans Snap di sini */}
      <Script
        src="https://app.sandbox.midtrans.com/snap/snap.js"
        data-client-key={process.env.NEXT_PUBLIC_MIDTRANS_CLIENT_KEY}
        strategy="afterInteractive"
      />
      <main className="bg-slate-50 min-h-screen">
        <div className="container mx-auto px-4 py-8 md:py-12">
          <header className="mb-8">
            <h1 className="text-3xl md:text-4xl font-bold text-gray-800">
              Konfirmasi & Pembayaran
            </h1>
            <p className="mt-2 text-md text-gray-500">
              Periksa kembali detail janji temu Anda sebelum melanjutkan ke
              pembayaran.
            </p>
          </header>

          <div className="grid grid-cols-1 lg:grid-cols-3 gap-8 lg:gap-12 items-start">
            {/* Kolom Kiri: Detail Pasien & Alur Pembayaran */}
            <div className="lg:col-span-2 flex flex-col gap-6">
              <div className="bg-white p-6 rounded-xl border border-gray-200 shadow-sm">
                <h2 className="text-xl font-bold text-gray-800 mb-4 flex items-center">
                  <User className="w-5 h-5 mr-2" /> Detail Pasien
                </h2>
                <p>
                  <span className="text-gray-500">Nama:</span>{" "}
                  <span className="font-semibold">{mockUser.name}</span>
                </p>
                <p>
                  <span className="text-gray-500">Email:</span>{" "}
                  <span className="font-semibold">{mockUser.email}</span>
                </p>
              </div>
              <PaymentFlow
                consultationFee={doctor.consultationFee}
                onPay={handlePayment}
              />
            </div>

            {/* Kolom Kanan: Ringkasan Booking (Sticky) */}
            <div className="lg:sticky top-24">
              <BookingSummaryCard doctor={doctor} booking={booking} />
            </div>
          </div>
        </div>
      </main>
    </>
  );
}

"use client";

import { useState } from "react";
import { PrescriptionCard } from "@/components/cards/PrescriptionCard";
import { Prescription } from "@/types/index.type";
import { Search } from "lucide-react";
import { Navbar } from "@/components/layout/Navbar";
import { Input } from "@/components/ui/form/Input";

// --- DATA MOCK (Ganti dengan data asli dari API) ---
const mockPrescriptions: Prescription[] = [
  {
    id: "RX-001",
    doctorName: "Dr. Adinda Melati, Sp.A",
    doctorSpecialty: "Dokter Anak",
    date: "19 September 2025",
    isRedeemed: false,
    medicines: [
      {
        id: "med-1",
        name: "Paracetamol Sirup",
        dosage: "120 mg/5 ml",
        imageUrl: "/obat/paracetamol.jpg",
        price: 15000,
        quantity: 1,
      },
      {
        id: "med-2",
        name: "Cetirizine Tablet",
        dosage: "10 mg",
        imageUrl: "/obat/cetirizine.jpg",
        price: 8000,
        quantity: 10,
      },
    ],
  },
  {
    id: "RX-002",
    doctorName: "Dr. Citra Ayu, Sp.KK",
    doctorSpecialty: "Kulit & Kelamin",
    date: "15 September 2025",
    isRedeemed: false,
    medicines: [
      {
        id: "med-3",
        name: "Hydrocortisone 1% Cream",
        dosage: "10 gr",
        imageUrl: "/obat/hydrocortisone.jpg",
        price: 25000,
        quantity: 1,
      },
    ],
  },
];

export default function PharmacyPage() {
  const [selectedPrescription, setSelectedPrescription] =
    useState<Prescription | null>(null);

  const activePrescriptions = mockPrescriptions.filter((p) => !p.isRedeemed);

  return (
    <>
      <Navbar />

      <main className="bg-slate-50 min-h-screen">
        <div className="container mx-auto px-4 py-8 md:py-12">
          {/* Hero Section */}
          <div className="text-center bg-cyan-600 text-white p-10 rounded-xl shadow-lg">
            <h1 className="text-4xl font-bold">HealthApp Online Pharmacy</h1>
            <p className="mt-2 max-w-2xl mx-auto">
              Fill your doctor's digital prescription quickly, safely, and
              easily. Get genuine, complete medications, and have them delivered
              directly to your door.
            </p>
          </div>

          <section className="mt-12">
            <h2 className="text-2xl font-bold text-gray-800 mb-4">
              Your Unredeemed Prescriptions
            </h2>
            {activePrescriptions.length > 0 ? (
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                {activePrescriptions.map((p) => (
                  <PrescriptionCard
                    key={p.id}
                    prescription={p}
                    onSelect={() => setSelectedPrescription(p)}
                  />
                ))}
              </div>
            ) : (
              <div className="text-center p-8 bg-white rounded-lg border border-dashed">
                <p className="text-gray-500">
                  You do not have an unredeemed prescription at this time.
                </p>
              </div>
            )}
          </section>

          {/* Seksi Pencarian Obat Bebas */}
          <section className="mt-12">
            <h2 className="text-2xl font-bold text-gray-800 mb-4">
              Search What You Need
            </h2>
            <Input
              type="text"
              placeholder="search vitamin, antibiotics, medicines, and more..."
              startIcon={<Search className="w-5 h-5 text-gray-400" />}
              className="w-full text-md py-6 rounded-lg"
            />
          </section>
        </div>
      </main>

      {/* Modal Detail Resep */}
      {/* {selectedPrescription && (
        <PrescriptionDetailModal
          prescription={selectedPrescription}
          onClose={() => setSelectedPrescription(null)}
        />
      )} */}
    </>
  );
}

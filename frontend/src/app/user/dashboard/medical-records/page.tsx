"use client";
import { Upload } from "lucide-react";
import { MedicalRecordItem } from "./components/MedicalRecordItem";

// --- DATA DUMMY ---
const mockMedicalRecords = [
  {
    id: 1,
    type: "consultation",
    title: "Konsultasi Demam & Batuk",
    date: "19 September 2025",
    details: {
      doctor: "Dr. Adinda Melati, Sp.A",
      summary:
        "Pasien didiagnosis dengan infeksi saluran pernapasan atas (ISPA). Diberikan resep untuk paracetamol dan anjuran istirahat.",
    },
  },
  {
    id: 2,
    type: "prescription",
    title: "Resep #RX-USER-01",
    date: "19 September 2025",
    details: {
      doctor: "Dr. Adinda Melati, Sp.A",
      summary:
        "Resep untuk Paracetamol Sirup dan Cetirizine 10mg telah diterbitkan dan ditebus.",
    },
    attachment: "resep-rx-user-01.pdf",
  },
  {
    id: 3,
    type: "lab_result",
    title: "Hasil Tes Darah Lengkap",
    date: "10 September 2025",
    details: {
      doctor: "Lab Prodia",
      summary:
        "Hasil tes darah menunjukkan kadar leukosit normal. Tidak ada indikasi infeksi bakteri.",
    },
    attachment: "hasil-lab-darah-100925.pdf",
  },
  {
    id: 4,
    type: "user_upload",
    title: "Rontgen Dada",
    date: "05 September 2025",
    details: {
      doctor: "Diunggah oleh Anda",
      summary: "Hasil rontgen dada dari RS Mitra Keluarga.",
    },
    attachment: "rontgen-dada-050925.jpeg",
  },
];

export default function MedicalRecordsPage() {
  return (
    <div className="space-y-6">
      <header className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <div>
          <h1 className="text-3xl font-bold text-gray-800">Rekam Medis Saya</h1>
          <p className="mt-1 text-gray-500">
            Linimasa perjalanan kesehatan dan dokumen medis Anda.
          </p>
        </div>
        <button className="flex items-center justify-center gap-2 px-5 py-2.5 font-semibold text-white bg-cyan-500 rounded-lg shadow-sm hover:bg-cyan-600">
          <Upload className="w-5 h-5" />
          Unggah Dokumen Baru
        </button>
      </header>

      {/* Filter (Opsional, bisa ditambahkan di sini) */}

      {/* Linimasa / Timeline */}
      <div className="relative">
        {/* Garis Vertikal Linimasa */}
        <div className="absolute left-6 top-0 h-full w-0.5 bg-gray-200"></div>

        <div className="space-y-8">
          {mockMedicalRecords.map((record) => (
            <MedicalRecordItem key={record.id} record={record} />
          ))}
        </div>
      </div>

      <div className="text-center text-sm text-gray-400 pt-8">
        <p>Data rekam medis Anda dienkripsi dan disimpan dengan aman.</p>
      </div>
    </div>
  );
}

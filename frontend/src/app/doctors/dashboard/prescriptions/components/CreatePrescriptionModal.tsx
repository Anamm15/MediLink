"use client";
import { useState } from "react";
import { motion } from "framer-motion";
import { X, UserSearch, Pill, Plus, Trash2 } from "lucide-react";

// Data dummy untuk autocomplete
const mockPatients = [
  { id: 1, name: "Budi Setiawan" },
  { id: 2, name: "Sarah L." },
];
const mockMedicines = [
  { id: "med-1", name: "Paracetamol 500mg" },
  { id: "med-2", name: "Amoxicillin 250mg" },
];

interface ModalProps {
  onClose: () => void;
}

export const CreatePrescriptionModal = ({ onClose }: ModalProps) => {
  const [selectedPatient, setSelectedPatient] = useState<string | null>(null);
  const [medicines, setMedicines] = useState<
    { id: string; name: string; qty: string; sig: string }[]
  >([]);

  const addMedicine = () => {
    // Di aplikasi nyata, ini akan hasil dari pencarian
    const newMed = mockMedicines[0];
    setMedicines([...medicines, { ...newMed, qty: "", sig: "" }]);
  };

  const removeMedicine = (id: string) => {
    setMedicines(medicines.filter((med) => med.id !== id));
  };

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/60 backdrop-blur-sm">
      <motion.div
        initial={{ opacity: 0, scale: 0.9 }}
        animate={{ opacity: 1, scale: 1 }}
        exit={{ opacity: 0, scale: 0.9 }}
        className="bg-slate-50 rounded-xl shadow-2xl w-full max-w-2xl max-h-[90vh] flex flex-col"
      >
        {/* Header */}
        <div className="flex items-center justify-between p-4 border-b bg-white rounded-t-xl">
          <h2 className="text-xl font-bold text-gray-900">
            Buat Resep Digital Baru
          </h2>
          <button
            onClick={onClose}
            className="p-1 rounded-full hover:bg-gray-100"
          >
            <X className="w-6 h-6 text-gray-500" />
          </button>
        </div>

        {/* Konten Form */}
        <div className="p-6 space-y-6 overflow-y-auto">
          {/* 1. Pilih Pasien */}
          <div className="p-4 bg-white rounded-lg border">
            <label className="text-sm font-semibold text-gray-700 flex items-center gap-2">
              <UserSearch className="w-5 h-5" />
              Pilih Pasien
            </label>
            <input
              type="text"
              placeholder="Ketik nama pasien..."
              className="mt-2 w-full p-2 border rounded-md"
            />
            {/* Di sini akan ada logic autocomplete */}
          </div>

          {/* 2. Daftar Obat */}
          <div className="p-4 bg-white rounded-lg border">
            <div className="flex items-center justify-between">
              <h3 className="text-sm font-semibold text-gray-700 flex items-center gap-2">
                <Pill className="w-5 h-5" />
                Daftar Obat
              </h3>
              <button
                onClick={addMedicine}
                className="flex items-center gap-1 text-xs font-semibold text-cyan-600 hover:text-cyan-800"
              >
                <Plus className="w-4 h-4" /> Tambah Obat
              </button>
            </div>
            <div className="mt-4 space-y-3">
              {medicines.map((med) => (
                <div
                  key={med.id}
                  className="grid grid-cols-12 gap-2 p-2 border rounded-md bg-slate-50"
                >
                  <p className="col-span-12 font-semibold">{med.name}</p>
                  <input
                    type="text"
                    placeholder="Kuantitas"
                    className="col-span-4 p-1.5 border rounded-md text-sm"
                  />
                  <input
                    type="text"
                    placeholder="Aturan Pakai (e.g., 3x1 sesudah makan)"
                    className="col-span-7 p-1.5 border rounded-md text-sm"
                  />
                  <button
                    onClick={() => removeMedicine(med.id)}
                    className="col-span-1 flex items-center justify-center text-gray-400 hover:text-red-500"
                  >
                    <Trash2 className="w-4 h-4" />
                  </button>
                </div>
              ))}
              {medicines.length === 0 && (
                <p className="text-center text-xs text-gray-400 py-4">
                  Belum ada obat ditambahkan
                </p>
              )}
            </div>
          </div>
        </div>

        {/* Footer Aksi */}
        <div className="flex items-center justify-end gap-3 p-4 border-t bg-white rounded-b-xl">
          <button
            onClick={onClose}
            className="px-4 py-2 text-sm font-semibold text-gray-700 bg-gray-100 rounded-lg hover:bg-gray-200"
          >
            Batal
          </button>
          <button className="px-6 py-2 text-sm font-semibold text-white bg-cyan-500 rounded-lg hover:bg-cyan-600">
            Terbitkan Resep
          </button>
        </div>
      </motion.div>
    </div>
  );
};

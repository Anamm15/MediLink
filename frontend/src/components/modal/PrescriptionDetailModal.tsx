import { X, ShoppingCart } from "lucide-react";
import Image from "next/image";
import { Prescription } from "@/types";

interface ModalProps {
  prescription: Prescription;
  onClose: () => void;
}

export const PrescriptionDetailModal = ({
  prescription,
  onClose,
}: ModalProps) => {
  const total = prescription.medicines.reduce(
    (sum, item) => sum + item.price * item.quantity,
    0
  );

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/60 backdrop-blur-sm">
      <div className="bg-white rounded-xl shadow-2xl w-full max-w-2xl max-h-[90vh] flex flex-col">
        {/* Header */}
        <div className="flex items-start justify-between p-6 border-b">
          <div>
            <h2 className="text-xl font-bold text-gray-900">
              Detail Resep Digital
            </h2>
            <p className="text-sm text-gray-500">
              Dari {prescription.doctorName} - {prescription.date}
            </p>
          </div>
          <button
            onClick={onClose}
            className="p-1 rounded-full hover:bg-gray-100"
          >
            <X className="w-6 h-6 text-gray-500" />
          </button>
        </div>

        {/* Daftar Obat */}
        <div className="flex-grow p-6 space-y-4 overflow-y-auto">
          {prescription.medicines.map((med) => (
            <div key={med.id} className="flex items-center gap-4">
              <Image
                src={med.imageUrl}
                alt={med.name}
                width={64}
                height={64}
                className="rounded-md border object-contain"
              />
              <div className="flex-grow">
                <p className="font-semibold text-gray-800">{med.name}</p>
                <p className="text-sm text-gray-500">
                  {med.dosage} • {med.quantity} unit
                </p>
              </div>
              <p className="font-semibold text-gray-800">
                Rp {(med.price * med.quantity).toLocaleString("id-ID")}
              </p>
            </div>
          ))}
        </div>

        {/* Footer */}
        <div className="p-6 border-t bg-slate-50 rounded-b-xl">
          <div className="flex justify-between items-center mb-4">
            <span className="text-gray-600 font-medium">Total Harga</span>
            <span className="text-xl font-bold text-cyan-600">
              Rp {total.toLocaleString("id-ID")}
            </span>
          </div>
          <button className="w-full flex items-center justify-center gap-2 bg-slate-800 text-white font-bold py-3 rounded-lg hover:bg-slate-700">
            <ShoppingCart className="w-5 h-5" />
            Tambahkan Semua ke Keranjang
          </button>
        </div>
      </div>
    </div>
  );
};

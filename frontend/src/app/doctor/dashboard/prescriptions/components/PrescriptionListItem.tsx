import { User, Calendar, Hash, Eye } from "lucide-react";

// Tipe data untuk satu resep di daftar
type Prescription = any; // Ganti dengan tipe data Anda

interface ListItemProps {
  prescription: Prescription;
}

export const PrescriptionListItem = ({ prescription }: ListItemProps) => {
  const isRedeemed = prescription.status === "Ditebus";

  return (
    <div className="bg-white rounded-xl shadow-sm border p-4 grid grid-cols-1 md:grid-cols-4 items-center gap-4 transition-all hover:bg-slate-50">
      <div className="flex items-center gap-3">
        <User className="w-5 h-5 text-gray-400" />
        <span className="font-semibold text-gray-800">
          {prescription.patientName}
        </span>
      </div>
      <div className="flex items-center gap-3 text-sm text-gray-500">
        <Calendar className="w-5 h-5" />
        <span>{prescription.date}</span>
      </div>
      <div className="flex items-center gap-3 text-sm text-gray-500">
        <Hash className="w-5 h-5" />
        <span>{prescription.id}</span>
      </div>
      <div className="flex items-center justify-between md:justify-end gap-4">
        <span
          className={`text-xs font-bold px-3 py-1 rounded-full ${
            isRedeemed
              ? "bg-green-100 text-green-700"
              : "bg-yellow-100 text-yellow-700"
          }`}
        >
          {prescription.status}
        </span>
        <button className="p-2 rounded-md hover:bg-gray-200">
          <Eye className="w-5 h-5 text-gray-600" />
        </button>
      </div>
    </div>
  );
};

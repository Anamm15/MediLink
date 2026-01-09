import { ClipboardPlus, ChevronRight } from "lucide-react";
import { Prescription } from "@/types/index.type";

interface PrescriptionCardProps {
  prescription: Prescription;
  onSelect: () => void;
}

export const PrescriptionCard = ({
  prescription,
  onSelect,
}: PrescriptionCardProps) => {
  return (
    <button
      onClick={onSelect}
      className="w-full text-left flex items-center gap-4 p-4 bg-white border border-gray-200 rounded-lg shadow-sm hover:shadow-md hover:border-cyan-400 transition-all"
    >
      <div className="flex-shrink-0 p-3 bg-cyan-50 rounded-full">
        <ClipboardPlus className="w-6 h-6 text-cyan-600" />
      </div>
      <div className="flex-grow">
        <p className="text-sm text-gray-500">
          Resep dari{" "}
          <span className="font-semibold text-gray-700">
            {prescription.doctorName}
          </span>
        </p>
        <p className="text-xs text-gray-400">
          {prescription.doctorSpecialty} • {prescription.date}
        </p>
      </div>
      <div className="flex items-center gap-2 text-sm font-semibold text-cyan-600">
        Lihat Detail
        <ChevronRight className="w-4 h-4" />
      </div>
    </button>
  );
};

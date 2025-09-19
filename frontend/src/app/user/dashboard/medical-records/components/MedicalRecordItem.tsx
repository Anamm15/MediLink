"use client";
import { useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import {
  Stethoscope,
  FileText,
  Beaker,
  Upload,
  ChevronDown,
  Download,
} from "lucide-react";
import React from "react";

// Tipe data untuk satu entri rekam medis
type MedicalRecord = any;

const recordConfig = {
  consultation: { icon: <Stethoscope />, color: "blue", title: "Konsultasi" },
  prescription: {
    icon: <FileText />,
    color: "green",
    title: "Resep Diterbitkan",
  },
  lab_result: { icon: <Beaker />, color: "purple", title: "Hasil Lab" },
  user_upload: { icon: <Upload />, color: "gray", title: "Dokumen Unggahan" },
};

export const MedicalRecordItem = ({ record }: { record: MedicalRecord }) => {
  const [isOpen, setIsOpen] = useState(false);
  const config = recordConfig[record.type as keyof typeof recordConfig];
  const colorClasses = {
    icon: `text-${config.color}-600`,
    bg: `bg-${config.color}-100`,
    ring: `ring-${config.color}-500`,
    border: `border-${config.color}-200`,
  };

  return (
    <div className="relative pl-12">
      {/* Ikon Linimasa */}
      <div
        className={`absolute left-0 top-0 flex h-12 w-12 items-center justify-center rounded-full ${colorClasses.bg} ring-4 ring-white`}
      >
        {React.cloneElement(config.icon, {
          className: `w-6 h-6 ${colorClasses.icon}`,
        })}
      </div>

      {/* Kartu Konten */}
      <motion.div
        layout
        className={`ml-4 bg-white rounded-xl shadow-sm border ${
          isOpen ? `ring-2 ${colorClasses.ring}` : colorClasses.border
        }`}
      >
        <button
          onClick={() => setIsOpen(!isOpen)}
          className="w-full flex items-center justify-between p-4 text-left"
        >
          <div>
            <p className="font-semibold text-gray-800">
              {config.title}: {record.title}
            </p>
            <p className="text-sm text-gray-500">{record.date}</p>
          </div>
          <ChevronDown
            className={`w-5 h-5 text-gray-400 transition-transform ${
              isOpen ? "rotate-180" : ""
            }`}
          />
        </button>

        {/* Detail yang Dapat Diperluas */}
        <AnimatePresence>
          {isOpen && (
            <motion.div
              initial={{ opacity: 0, height: 0 }}
              animate={{ opacity: 1, height: "auto" }}
              exit={{ opacity: 0, height: 0 }}
              transition={{ duration: 0.3, ease: "easeInOut" }}
              className="overflow-hidden"
            >
              <div className="px-4 pb-4 border-t border-dashed">
                <div className="prose prose-sm mt-4 text-gray-600">
                  <p>
                    <strong>Dokter:</strong> {record.details.doctor}
                  </p>
                  <p>
                    <strong>Ringkasan:</strong>
                  </p>
                  <blockquote>{record.details.summary}</blockquote>
                  {record.attachment && (
                    <button className="flex items-center gap-2 mt-4 px-3 py-1.5 text-xs font-semibold text-gray-600 bg-gray-100 rounded-md hover:bg-gray-200">
                      <Download className="w-3 h-3" /> {record.attachment}
                    </button>
                  )}
                </div>
              </div>
            </motion.div>
          )}
        </AnimatePresence>
      </motion.div>
    </div>
  );
};

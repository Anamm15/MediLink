"use client";

import React, { useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import {
  Calendar,
  ChevronDown,
  User,
  Clock,
  CheckCircle,
  Phone,
} from "lucide-react";
import { PrescriptionResponse } from "@/types/prescription.type";
import { formatIDDate } from "@/helpers/datetime";

interface CardProps {
  prescription: PrescriptionResponse;
}

export default function PrescriptionCard({ prescription }: CardProps) {
  const [isExpanded, setIsExpanded] = useState(false);
  const { patient, is_redeemed, medicines, created_at, id } = prescription;

  return (
    <motion.div
      layout
      onClick={() => setIsExpanded(!isExpanded)}
      className={`bg-white rounded-xl shadow-sm border overflow-hidden cursor-pointer hover:shadow-md transition-shadow group ${
        isExpanded
          ? "ring-2 ring-cyan-500 border-transparent shadow-md"
          : "border-gray-200 hover:border-cyan-200"
      }`}
    >
      <motion.div layout className="p-4 flex items-center justify-between">
        <div className="flex items-center gap-3">
          <div className="w-10 h-10 rounded-full bg-blue-100 flex items-center justify-center text-blue-600">
            <User className="w-5 h-5" />
          </div>
          <div>
            <h3 className="font-semibold text-gray-800">
              {patient?.name || "Pasien"}
            </h3>
            <div className="flex items-center gap-2 text-xs text-gray-500">
              <Phone className="w-3 h-3" /> {patient?.phone_number || "-"}
            </div>
          </div>
        </div>

        <div className="flex items-center gap-4">
          <div className="text-right hidden sm:block">
            <p className="text-xs text-gray-400">Prescription ID</p>
            <p className="text-xs font-mono text-gray-600">
              #{id.substring(0, 8)}
            </p>
          </div>
          <motion.div
            animate={{ rotate: isExpanded ? 180 : 0 }}
            transition={{ duration: 0.2 }}
          >
            <ChevronDown className="w-5 h-5 text-gray-400" />
          </motion.div>
        </div>
      </motion.div>

      <AnimatePresence>
        {isExpanded && (
          <motion.div
            initial={{ height: 0, opacity: 0 }}
            animate={{ height: "auto", opacity: 1 }}
            exit={{ height: 0, opacity: 0 }}
            transition={{ duration: 0.3 }}
            className="bg-gray-50 border-t border-gray-100"
          >
            <div className="p-4 grid gap-4">
              <div className="flex justify-between items-center bg-white p-3 rounded-lg border border-gray-200">
                <div className="flex items-center gap-2 text-sm text-gray-600">
                  <Calendar className="w-4 h-4" />
                  {formatIDDate(created_at)}
                </div>
                <div
                  className={`flex items-center gap-1 text-xs font-bold px-2 py-1 rounded ${
                    is_redeemed
                      ? "bg-green-100 text-green-700"
                      : "bg-yellow-100 text-yellow-700"
                  }`}
                >
                  {is_redeemed ? (
                    <CheckCircle className="w-3 h-3" />
                  ) : (
                    <Clock className="w-3 h-3" />
                  )}
                  {is_redeemed ? "Redeemed" : "Pending"}
                </div>
              </div>

              <div>
                <h4 className="text-xs font-semibold text-gray-500 mb-2">
                  Prescription Items:
                </h4>
                <div className="flex flex-wrap gap-2">
                  {medicines.map((m, i) => (
                    <span
                      key={i}
                      className="text-xs bg-white border border-gray-200 px-2 py-1 rounded-md text-gray-700"
                    >
                      {m.name}{" "}
                      <span className="font-bold text-gray-400">
                        x{m.quantity}
                      </span>
                    </span>
                  ))}
                </div>
              </div>
            </div>
          </motion.div>
        )}
      </AnimatePresence>
    </motion.div>
  );
}

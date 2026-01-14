"use client";

import React, { useState } from "react";
import Image from "next/image";
import { motion, AnimatePresence } from "framer-motion";
import {
  Calendar,
  Pill,
  ShoppingCart,
  Eye,
  ChevronDown,
  Stethoscope,
} from "lucide-react";
import { PrescriptionResponse } from "@/types/prescription.type";
import { DEFAULT_PROFILE } from "@/helpers/constant";
import { formatIDDate } from "@/helpers/datetime";
import { formatCurrency } from "@/helpers/currency";

interface CardProps {
  prescription: PrescriptionResponse;
}

export default function PrescriptionCard({ prescription }: CardProps) {
  const [isExpanded, setIsExpanded] = useState(false);
  const { doctor, is_redeemed, medicines, created_at, notes } = prescription;
  const isActive = !is_redeemed;

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
        <div className="flex items-center gap-4">
          <div className="relative">
            <Image
              src={doctor?.avatar_url || DEFAULT_PROFILE}
              alt={doctor?.name || "Doctor"}
              width={50}
              height={50}
              className="rounded-full object-cover border border-gray-100"
            />
            <div className="absolute -bottom-1 -right-1 bg-white rounded-full p-1 border shadow-sm">
              <Stethoscope className="w-3 h-3 text-cyan-600" />
            </div>
          </div>
          <div>
            <h3 className="font-semibold text-gray-800 group-hover:text-cyan-600 transition-colors">
              {doctor?.name}
            </h3>
            <p className="text-xs text-gray-500">{doctor?.specialization}</p>
            <div className="flex items-center gap-2 mt-1 md:hidden">
              <span
                className={`text-[10px] px-2 py-0.5 rounded-full font-medium ${
                  isActive
                    ? "bg-yellow-100 text-yellow-700"
                    : "bg-green-100 text-green-700"
                }`}
              >
                {isActive ? "Not Redeemed" : "Redeemed"}
              </span>
            </div>
          </div>
        </div>

        <div className="flex flex-col items-end gap-2">
          <div className="hidden md:flex items-center gap-2 text-sm text-gray-500">
            <Calendar className="w-4 h-4" />
            <span>{formatIDDate(created_at)}</span>
          </div>
          <div className="flex items-center gap-3">
            <span
              className={`hidden md:inline-block text-xs font-bold px-3 py-1 rounded-full ${
                isActive
                  ? "bg-yellow-100 text-yellow-700"
                  : "bg-green-100 text-green-700"
              }`}
            >
              {isActive ? "Not Redeemed" : "Redeemed"}
            </span>
            <motion.div
              animate={{ rotate: isExpanded ? 180 : 0 }}
              transition={{ duration: 0.2 }}
            >
              <ChevronDown className="w-5 h-5 text-gray-400" />
            </motion.div>
          </div>
        </div>
      </motion.div>

      <AnimatePresence>
        {isExpanded && (
          <motion.div
            initial={{ height: 0, opacity: 0 }}
            animate={{ height: "auto", opacity: 1 }}
            exit={{ height: 0, opacity: 0 }}
            transition={{ duration: 0.3, ease: "easeInOut" }}
            className="overflow-hidden bg-slate-50/50"
          >
            <div className="p-4 border-t border-gray-100 space-y-4">
              {/* List Obat */}
              <div>
                <h4 className="text-xs font-bold text-gray-500 uppercase tracking-wider mb-2 flex items-center gap-2">
                  <Pill className="w-4 h-4" /> Medicines List
                </h4>
                <div className="bg-white rounded-lg border border-gray-200 divide-y divide-gray-100">
                  {medicines.map((med, index) => (
                    <div
                      key={index}
                      className="p-3 flex justify-between items-center text-sm"
                    >
                      <div>
                        <p className="font-medium text-gray-800">{med.name}</p>
                        <p className="text-xs text-gray-500">{med.category}</p>
                      </div>
                      <div className="text-right">
                        <span className="block font-semibold text-gray-700">
                          x{med.quantity}
                        </span>
                        <span className="text-xs text-gray-400">
                          {formatCurrency(med.base_price * med.quantity)}
                        </span>
                      </div>
                    </div>
                  ))}
                </div>
              </div>

              {notes && (
                <div className="bg-blue-50 p-3 rounded-lg text-sm text-blue-800 border border-blue-100">
                  <span className="font-semibold">Notes:</span> {notes}
                </div>
              )}

              {/* Action Buttons */}
              <div className="pt-2">
                {isActive ? (
                  <button
                    onClick={(e) => {
                      e.stopPropagation();
                    }}
                    className="w-full flex items-center justify-center gap-2 px-4 py-2.5 bg-cyan-500 hover:bg-cyan-600 text-white font-semibold rounded-lg transition-colors shadow-sm"
                  >
                    <ShoppingCart className="w-4 h-4" /> Redeem Now
                  </button>
                ) : (
                  <button
                    onClick={(e) => {
                      e.stopPropagation();
                    }}
                    className="w-full flex items-center justify-center gap-2 px-4 py-2.5 bg-white border border-gray-300 hover:bg-gray-50 text-gray-700 font-semibold rounded-lg transition-colors"
                  >
                    <Eye className="w-4 h-4" /> View Transaction Details
                  </button>
                )}
              </div>
            </div>
          </motion.div>
        )}
      </AnimatePresence>
    </motion.div>
  );
}

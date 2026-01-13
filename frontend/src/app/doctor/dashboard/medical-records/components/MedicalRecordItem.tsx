"use client";

import { useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import {
  ChevronDown,
  Stethoscope,
  FileText,
  Activity,
  Search,
  ClipboardCheck,
  Pill,
  ExternalLink,
  Calendar,
} from "lucide-react";
import Link from "next/link";
import React from "react";
import { MedicalRecordResponse } from "@/types/medical_record.type";

export const MedicalRecordItem = ({
  record,
}: {
  record: MedicalRecordResponse;
}) => {
  const [isOpen, setIsOpen] = useState(false);

  const soapFields = [
    {
      key: "subjective",
      label: "Subjective (Complaint)",
      content: record.subjective,
      icon: <Activity className="w-4 h-4 text-orange-500" />,
      bg: "bg-orange-50",
      border: "border-orange-100",
      placeholder: "No subjective complaints recorded.",
    },
    {
      key: "objective",
      label: "Objective (Examination)",
      content: record.objective,
      icon: <Search className="w-4 h-4 text-blue-500" />,
      bg: "bg-blue-50",
      border: "border-blue-100",
      placeholder: "No physical examination data available.",
    },
    {
      key: "assessment",
      label: "Assessment (Diagnosis)",
      content: record.assessment,
      icon: <ClipboardCheck className="w-4 h-4 text-purple-500" />,
      bg: "bg-purple-50",
      border: "border-purple-100",
      placeholder: "Diagnosis not determined yet.",
    },
    {
      key: "plan",
      label: "Plan (Plan & Medication)",
      content: record.plan,
      icon: <Pill className="w-4 h-4 text-green-500" />,
      bg: "bg-green-50",
      border: "border-green-100",
      placeholder: "No specific follow-up plan.",
    },
  ];

  return (
    <div className="relative pl-12 group">
      {/* Timeline Vertical Line (Hidden on last item) */}
      <div className="absolute left-6 top-10 bottom-0 w-0.5 bg-gray-100 group-last:hidden"></div>

      {/* Main Timeline Icon */}
      <div className="absolute left-0 top-0 flex h-12 w-12 items-center justify-center rounded-full bg-cyan-50 ring-4 ring-white border border-cyan-100 z-10">
        <FileText className="w-5 h-5 text-cyan-600" />
      </div>

      {/* Content Card */}
      <motion.div
        layout
        className={`ml-4 bg-white rounded-2xl shadow-sm border transition-all duration-300 ${
          isOpen
            ? "ring-2 ring-cyan-500 border-transparent shadow-md"
            : "border-gray-200 hover:border-cyan-200"
        }`}
      >
        {/* Card Header (Click to Expand) */}
        <button
          onClick={() => setIsOpen(!isOpen)}
          className="w-full flex items-center justify-between p-5 text-left focus:outline-none"
        >
          <div>
            <div className="flex items-center gap-2 mb-1">
              <span className="text-[10px] font-bold uppercase tracking-wider text-cyan-700 bg-cyan-100 px-2 py-0.5 rounded">
                Medical Record
              </span>
              <span className="text-xs text-gray-400 font-mono">
                #{record.id.slice(0, 8)}
              </span>
            </div>
            <h3 className="font-bold text-gray-800 text-lg">{record.title}</h3>
            <p className="text-sm text-gray-500 mt-1 flex items-center gap-1.5">
              <Calendar className="w-3.5 h-3.5" />
              {record.date}
            </p>
          </div>

          <div
            className={`p-2 rounded-full transition-colors ${
              isOpen ? "bg-cyan-500 text-white" : "bg-gray-50 text-gray-400"
            }`}
          >
            <ChevronDown
              className={`w-5 h-5 transition-transform duration-300 ${
                isOpen ? "rotate-180" : ""
              }`}
            />
          </div>
        </button>

        {/* SOAP Details (Expandable) */}
        <AnimatePresence>
          {isOpen && (
            <motion.div
              initial={{ opacity: 0, height: 0 }}
              animate={{ opacity: 1, height: "auto" }}
              exit={{ opacity: 0, height: 0 }}
              transition={{ duration: 0.3, ease: "easeInOut" }}
              className="overflow-hidden bg-slate-50/50 rounded-b-2xl"
            >
              <div className="p-5 border-t border-gray-100 space-y-4">
                {/* SOAP Grid */}
                <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                  {soapFields.map((field) => (
                    <div
                      key={field.key}
                      className={`p-4 rounded-xl border ${field.border} bg-white shadow-sm`}
                    >
                      <div className="flex items-center gap-2 mb-2">
                        {field.icon}
                        <span className="text-xs font-black uppercase text-gray-500 tracking-wider">
                          {field.label}
                        </span>
                      </div>
                      <p
                        className={`text-sm leading-relaxed ${
                          field.content
                            ? "text-gray-700 font-medium"
                            : "text-gray-400 italic"
                        }`}
                      >
                        {field.content || field.placeholder}
                      </p>
                    </div>
                  ))}
                </div>

                {/* Footer: Metadata & Link */}
                <div className="flex flex-col sm:flex-row items-center justify-between gap-4 mt-2 pt-4 border-t border-gray-200/60">
                  <div className="text-[11px] text-gray-400">
                    <p>
                      Created:{" "}
                      {new Date(record.created_at).toLocaleString("en-US")}
                    </p>
                    <p>Appointment ID: {record.appointment_id}</p>
                  </div>

                  <Link
                    href={`/doctor/dashboard/appointments/${record.appointment_id}`}
                    className="flex items-center gap-2 text-xs font-bold text-cyan-700 hover:text-cyan-800 hover:underline transition-all"
                  >
                    <Stethoscope className="w-3.5 h-3.5" />
                    View Appointment Details{" "}
                    <ExternalLink className="w-3 h-3" />
                  </Link>
                </div>
              </div>
            </motion.div>
          )}
        </AnimatePresence>
      </motion.div>
    </div>
  );
};

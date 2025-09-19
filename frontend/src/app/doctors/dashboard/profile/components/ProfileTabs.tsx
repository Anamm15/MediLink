"use client";
import { useState } from "react";
import { SectionCard } from "@/components/cards/SectionCard";
import { PracticeSchedule } from "./PracticeSchedule";
import { motion, AnimatePresence } from "framer-motion"; // <-- Import motion dan AnimatePresence

const tabs = [
  { id: "info", label: "Informasi Detail" },
  { id: "schedule", label: "Kelola Jadwal Praktik" },
];

export const ProfileTabs = ({ doctorData }: { doctorData: any }) => {
  const [activeTab, setActiveTab] = useState(tabs[0].id);

  const tabContentVariants = {
    initial: { opacity: 0, y: 20 },
    animate: { opacity: 1, y: 0 },
    exit: { opacity: 0, y: -20 },
  };

  return (
    <div className="flex flex-col gap-6">
      {/* Tombol Tab dengan Indikator Animasi */}
      <div className="bg-white p-1 rounded-xl shadow-sm border flex items-center gap-2">
        {tabs.map((tab) => (
          <button
            key={tab.id}
            onClick={() => setActiveTab(tab.id)}
            className={`w-full py-2.5 text-sm font-semibold rounded-lg relative transition-colors ${
              activeTab === tab.id
                ? "text-cyan-600"
                : "text-gray-500 hover:bg-gray-100"
            }`}
          >
            {tab.label}
          </button>
        ))}
      </div>

      {/* Konten Tab dengan Animasi Transisi */}
      <div>
        <AnimatePresence mode="wait">
          {" "}
          {/* 'mode="wait"' memastikan animasi exit selesai sebelum animasi enter dimulai */}
          <motion.div
            key={activeTab} // <-- Kunci unik ini memberitahu AnimatePresence kapan konten berubah
            variants={tabContentVariants}
            initial="initial"
            animate="animate"
            exit="exit"
            transition={{ duration: 0.3, ease: "easeInOut" }}
          >
            {activeTab === "info" && (
              <div className="space-y-6">
                <SectionCard title="Tentang Saya">
                  <p className="text-gray-600 text-sm leading-relaxed">
                    {doctorData.bio}
                  </p>
                </SectionCard>
                <SectionCard title="Pendidikan">
                  <ul className="space-y-2 text-sm">
                    {doctorData.education.map((edu: any, index: number) => (
                      <li key={index}>
                        <span className="font-semibold">{edu.degree}</span> -{" "}
                        {edu.university} ({edu.year})
                      </li>
                    ))}
                  </ul>
                </SectionCard>
              </div>
            )}
            {activeTab === "schedule" && <PracticeSchedule />}
          </motion.div>
        </AnimatePresence>
      </div>
    </div>
  );
};

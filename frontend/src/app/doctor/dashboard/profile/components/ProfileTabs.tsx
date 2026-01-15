"use client";
import { useState } from "react";
import { SectionCard } from "@/components/cards/SectionCard";
import { PracticeSchedule } from "./PracticeSchedule";
import { motion, AnimatePresence } from "framer-motion";
import { DoctorProfileResponse } from "@/types/doctor.type";

const tabs = [
  { id: "info", label: "Detail Information" },
  { id: "schedule", label: "Schedule Management" },
];

export const ProfileTabs = ({
  doctorData,
}: {
  doctorData: DoctorProfileResponse;
}) => {
  const [activeTab, setActiveTab] = useState(tabs[0].id);

  const tabContentVariants = {
    initial: { opacity: 0, y: 20 },
    animate: { opacity: 1, y: 0 },
    exit: { opacity: 0, y: -20 },
  };

  return (
    <div className="flex flex-col gap-6">
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

      <div>
        <AnimatePresence mode="wait">
          {" "}
          <motion.div
            key={activeTab}
            variants={tabContentVariants}
            initial="initial"
            animate="animate"
            exit="exit"
            transition={{ duration: 0.3, ease: "easeInOut" }}
          >
            {activeTab === "info" && (
              <div className="space-y-6">
                <SectionCard title="About Me">
                  <p className="text-gray-600 text-sm leading-relaxed">
                    {doctorData.bio}
                  </p>
                </SectionCard>
                <SectionCard title="Education">
                  <ul className="space-y-2 text-sm">
                    {doctorData.education?.map((edu, index) => (
                      <li key={index}>
                        <span className="font-semibold">{edu.degree}</span> -{" "}
                        {edu.institution} ({edu.year})
                      </li>
                    ))}
                  </ul>
                </SectionCard>
                <SectionCard title="Clinic">
                  <ul className="space-y-2 text-sm">
                    {doctorData.clinic?.map((clinic, index) => (
                      <li key={index}>
                        <div className="font-semibold">{clinic.name}</div>
                        <div>
                          {clinic.address}, {clinic.city}
                        </div>
                      </li>
                    ))}
                  </ul>
                </SectionCard>
              </div>
            )}
            {activeTab === "schedule" && doctorData.id && (
              <PracticeSchedule doctor_id={doctorData.id} />
            )}
          </motion.div>
        </AnimatePresence>
      </div>
    </div>
  );
};

"use client";
import { useState } from "react";
import { SectionCard } from "@/components/cards/SectionCard";
import { Trash2, PlusCircle } from "lucide-react";

// Initial dummy data for schedule
const initialSchedule = [
  { day: "Senin", isActive: true, slots: ["09:00 - 12:00", "14:00 - 17:00"] },
  { day: "Selasa", isActive: true, slots: ["09:00 - 12:00"] },
  { day: "Rabu", isActive: false, slots: [] },
  { day: "Kamis", isActive: true, slots: ["18:00 - 21:00"] },
  { day: "Jumat", isActive: true, slots: ["09:00 - 11:00"] },
  { day: "Sabtu", isActive: false, slots: [] },
  { day: "Minggu", isActive: false, slots: [] },
];

export const PracticeSchedule = () => {
  const [schedule, setSchedule] = useState(initialSchedule);

  const toggleDay = (day: string) => {
    setSchedule(
      schedule.map((d) => (d.day === day ? { ...d, isActive: !d.isActive } : d))
    );
  };

  // Placeholder functions for adding/removing slots
  const addSlot = (day: string) => alert(`Tambah slot untuk hari ${day}`);
  const removeSlot = (day: string, slot: string) =>
    alert(`Hapus slot ${slot} dari hari ${day}`);

  return (
    <SectionCard title="Kelola Jadwal Praktik Mingguan">
      <div className="space-y-4">
        {schedule.map((d) => (
          <div key={d.day} className="p-3 rounded-lg border">
            <div className="flex items-center justify-between">
              <span
                className={`font-semibold ${
                  d.isActive ? "text-gray-800" : "text-gray-400"
                }`}
              >
                {d.day}
              </span>
              <div className="flex items-center gap-2">
                <span
                  className={`text-xs font-medium ${
                    d.isActive ? "text-green-600" : "text-gray-500"
                  }`}
                >
                  {d.isActive ? "Aktif" : "Tidak Aktif"}
                </span>
                {/* Toggle Switch */}
                <button
                  onClick={() => toggleDay(d.day)}
                  className={`relative inline-flex h-6 w-11 items-center rounded-full transition-colors ${
                    d.isActive ? "bg-cyan-500" : "bg-gray-200"
                  }`}
                >
                  <span
                    className={`inline-block h-4 w-4 transform rounded-full bg-white transition-transform ${
                      d.isActive ? "translate-x-6" : "translate-x-1"
                    }`}
                  />
                </button>
              </div>
            </div>
            {d.isActive && (
              <div className="mt-3 pt-3 border-t border-dashed">
                <p className="text-sm text-gray-500 mb-2">
                  Slot Waktu Tersedia:
                </p>
                <div className="flex flex-wrap gap-2">
                  {d.slots.map((slot) => (
                    <div
                      key={slot}
                      className="flex items-center gap-1 bg-cyan-50 text-cyan-700 text-xs font-semibold px-2 py-1 rounded-full"
                    >
                      {slot}
                      <button
                        onClick={() => removeSlot(d.day, slot)}
                        className="hover:text-red-500"
                      >
                        <Trash2 className="w-3 h-3" />
                      </button>
                    </div>
                  ))}
                  <button
                    onClick={() => addSlot(d.day)}
                    className="flex items-center gap-1 text-cyan-600 hover:text-cyan-800 text-xs font-semibold"
                  >
                    <PlusCircle className="w-4 h-4" /> Tambah Slot
                  </button>
                </div>
              </div>
            )}
          </div>
        ))}
        <button className="mt-4 w-full font-semibold py-2.5 rounded-lg bg-cyan-500 text-white hover:bg-cyan-600">
          Simpan Perubahan Jadwal
        </button>
      </div>
    </SectionCard>
  );
};

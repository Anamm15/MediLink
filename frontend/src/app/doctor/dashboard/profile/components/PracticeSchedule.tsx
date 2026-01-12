"use client";

import { useEffect, useState } from "react";
import { SectionCard } from "@/components/cards/SectionCard";
import {
  Trash2,
  PlusCircle,
  Clock,
  Banknote,
  Users,
  Edit,
  MapPin,
} from "lucide-react";
import {
  useDeleteSchedule,
  useSchedulesQuery,
  useUpdateStatusSchedule,
} from "../hooks/useSchedule";
import { formatSchedule, formattedSchedule, Slot } from "@/types/schedule.type";
import ScheduleModal from "./ScheduleModal";
import DeleteAlert from "@/components/ui/DeleteAlert";

export const PracticeSchedule = ({ doctor_id }: { doctor_id: string }) => {
  const { data } = useSchedulesQuery(doctor_id);
  const [schedules, setSchedules] = useState<formattedSchedule[]>([]);
  const [deleteAlert, setDeleteAlert] = useState(false);
  const [isScheduleModalOpen, setIsScheduleModalOpen] = useState(false);
  const [selectedDay, setSelectedDay] = useState<string>("");
  const [selectedSlotId, setSelectedSlotId] = useState<string>("");
  const [isEdit, setIsEdit] = useState(false);
  const [selectedSlot, setSelectedSlot] = useState<Slot | null>(null);
  const { mutateAsync: updateStatus } = useUpdateStatusSchedule(
    setSchedules,
    selectedSlotId
  );
  const { mutateAsync: deleteSchedule } = useDeleteSchedule(
    schedules,
    setSchedules,
    selectedSlotId
  );

  useEffect(() => {
    if (data) setSchedules(formatSchedule(data));
  }, [data]);

  const handleToggleSlot = async (
    id: string,
    day: string,
    is_active: boolean
  ) => {
    setSelectedSlotId(id);
    await updateStatus({
      id: id,
      is_active: !is_active,
    });
  };

  const handleAddSlot = (day: string) => {
    setSelectedDay(day);
    setIsScheduleModalOpen(true);
    setSelectedSlotId("");
  };

  const handleUpdateSlot = (slot: Slot) => {
    setSelectedSlot(slot);
    setIsScheduleModalOpen(true);
    setIsEdit(true);
  };

  const handleDeleteSlot = async (id: string) => {
    await deleteSchedule(id);
    setSelectedSlotId("");
    setDeleteAlert(false);
  };

  return (
    <SectionCard title="Manage Your Weekly Schedule">
      <div className="space-y-6">
        {schedules.map((schedule) => (
          <div key={schedule.day} className="group">
            <div className="flex items-center justify-between mb-3">
              <h4 className="font-bold text-gray-800 flex items-center gap-2">
                <span className="w-2 h-2 rounded-full bg-cyan-500"></span>
                {schedule.day}
              </h4>
              <button
                onClick={() => handleAddSlot(schedule.day)}
                className="text-xs font-bold text-cyan-600 hover:text-cyan-700 flex items-center gap-1 bg-cyan-50 px-2 py-1 rounded-md transition-colors"
              >
                <PlusCircle className="w-3.5 h-3.5" /> Add Slot
              </button>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-3">
              {schedule.slots.length > 0 ? (
                schedule.slots.map((slot) => (
                  <div
                    key={slot.id}
                    className={`relative p-4 rounded-xl border transition-all duration-300 ${
                      slot.isActive
                        ? "bg-white border-gray-200 shadow-sm"
                        : "bg-gray-50 border-gray-100 opacity-70"
                    }`}
                  >
                    {/* Baris Atas: Waktu & Toggle */}
                    <div className="flex justify-between items-start mb-3">
                      <div className="flex items-center gap-2 text-gray-800">
                        <Clock
                          className={`w-4 h-4 ${
                            slot.isActive ? "text-cyan-500" : "text-gray-400"
                          }`}
                        />
                        <span className="font-bold text-sm">
                          {slot.startTime} - {slot.endTime}
                        </span>
                      </div>

                      <div className="flex items-center gap-3">
                        {/* Toggle Switch Per Slot */}
                        <button
                          onClick={() =>
                            handleToggleSlot(
                              slot.id,
                              schedule.day,
                              slot.isActive
                            )
                          }
                          className={`relative inline-flex h-5 w-9 items-center rounded-full transition-colors focus:outline-none ${
                            slot.isActive ? "bg-cyan-500" : "bg-gray-400"
                          }`}
                        >
                          <span
                            className={`inline-block h-3 w-3 transform rounded-full bg-white transition-transform ${
                              slot.isActive ? "translate-x-5" : "translate-x-1"
                            }`}
                          />
                        </button>

                        <button
                          onClick={() => handleUpdateSlot(slot)}
                          className="text-gray-500 hover:text-blue-600 transition-colors"
                        >
                          <Edit className="w-4 h-4" />
                        </button>

                        <button
                          onClick={() => {
                            setDeleteAlert(true);
                            setSelectedSlotId(slot.id);
                          }}
                          className="text-gray-500 hover:text-red-600 transition-colors"
                        >
                          <Trash2 className="w-4 h-4" />
                        </button>
                      </div>
                    </div>

                    {/* Info Fee & Quota */}
                    <div className="flex items-center gap-4 pt-3 border-t border-gray-50">
                      <div className="flex items-center gap-1.5 text-xs text-gray-500">
                        <Banknote className="w-3.5 h-3.5" />
                        <span>
                          Rp {slot.consultation_fee.toLocaleString("id-ID")}
                        </span>
                      </div>
                      <div className="flex items-center gap-1.5 text-xs text-gray-500">
                        <Users className="w-3.5 h-3.5" />
                        <span>Quota: {slot.max_quota}</span>
                      </div>
                      <div className="flex items-center gap-1.5 text-xs text-gray-500">
                        <MapPin className="w-3.5 h-3.5" />
                        <span>Type: {slot.type}</span>
                      </div>
                    </div>
                  </div>
                ))
              ) : (
                <div className="col-span-2 py-4 px-4 border border-dashed rounded-xl text-center">
                  <p className="text-xs text-gray-400 italic">
                    There is no practice schedule for today
                  </p>
                </div>
              )}
            </div>

            {/* Divider Antar Hari */}
            <div className="mt-6 border-b border-gray-100 last:hidden"></div>
          </div>
        ))}
      </div>

      {deleteAlert && (
        <DeleteAlert
          title="Delete Schedule"
          onClose={() => setDeleteAlert(false)}
          onConfirm={() => handleDeleteSlot(selectedSlotId)}
        />
      )}

      {isScheduleModalOpen && (
        <ScheduleModal
          day={selectedDay}
          isEdit={isEdit}
          setIsEdit={setIsEdit}
          slot={selectedSlot}
          setIsModalOpen={setIsScheduleModalOpen}
        />
      )}
    </SectionCard>
  );
};

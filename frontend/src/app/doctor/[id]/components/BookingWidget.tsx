"use client";

import { useState, useMemo, useEffect } from "react";
import { Calendar, Clock, AlertCircle } from "lucide-react";
import { DoctorScheduleResponse } from "@/types/schedule.type";
import { useParams } from "next/navigation";
import Link from "next/link";
import {
  formatIDDate,
  getCurrentDate,
  getCurrentTime,
  getDayOfDate,
  getToday,
} from "@/helpers/datetime";
import { useAvailableSchedulesQuery } from "../hooks/useAvailableSchedule";
import { toCapitalize } from "@/helpers/string";
import Calender from "@/components/ui/Calender";

export const BookingWidget = () => {
  const [selectedSchedule, setSelectedSchedule] =
    useState<DoctorScheduleResponse | null>(null);
  const [currentDate, setCurrentDate] = useState("");
  const [currentTime, setCurrentTime] = useState("");
  const { id } = useParams();
  const [selectedDate, setSelectedDate] = useState("");
  const [selectedDay, setSelectedDay] = useState("");
  const [isCalenderOpen, setIsCalenderOpen] = useState(false);
  const { data: schedules = [], refetch } = useAvailableSchedulesQuery(
    id as string,
    selectedDate,
    selectedDay
  );

  useEffect(() => {
    setCurrentTime(getCurrentTime());
    setCurrentDate(getCurrentDate());
    const { date, day } = getToday();
    setSelectedDate(date);
    setSelectedDay(day);
  }, []);

  useEffect(() => {
    const day = getDayOfDate(selectedDate);
    setSelectedDay(day);
    refetch();
  }, [selectedDate]);

  const isToday = true;

  const activeSchedules = useMemo(() => {
    if (!schedules) return [];
    return schedules.filter((s) => s.is_active);
  }, [schedules]);

  const isTimePassed = (startTime: string) => {
    if (!isToday || !currentTime) return false;
    if (currentDate !== selectedDate) return false;
    return startTime.slice(0, 5) < currentTime;
  };

  const displayFee = useMemo(() => {
    if (selectedSchedule) return selectedSchedule.consultation_fee;
    if (activeSchedules.length > 0) {
      return Math.min(...activeSchedules.map((s) => s.consultation_fee));
    }
    return 0;
  }, [selectedSchedule, activeSchedules]);

  return (
    <div className="p-6 bg-white rounded-2xl border border-gray-200 shadow-xl sticky top-24">
      <div className="mb-6 p-4 bg-cyan-50 rounded-xl border border-cyan-100">
        <p className="text-sm font-medium text-cyan-700 mb-1">
          {selectedSchedule ? "Consultation Fee" : "Start From"}
        </p>
        <div className="flex items-baseline gap-1">
          <span className="text-2xl font-bold text-cyan-600">
            Rp {displayFee.toLocaleString("id-ID")}
          </span>
        </div>
      </div>

      <div className="mt-6">
        <h3 className="font-bold text-gray-800 mb-3 flex items-center gap-2">
          <Calendar className="w-5 h-5 text-gray-400" />
          Select Day
        </h3>

        <div
          onClick={() => setIsCalenderOpen(true)}
          className="relative p-3 bg-gray-50 rounded-lg border border-gray-100 text-center cursor-pointer"
        >
          <p className="text-sm font-semibold text-gray-700">
            {toCapitalize(selectedDay)}, {formatIDDate(selectedDate)}{" "}
          </p>
        </div>
      </div>

      <div className="mt-8">
        <h3 className="font-bold text-gray-800 mb-3 flex items-center gap-2">
          <Clock className="w-5 h-5 text-gray-400" />
          Select Time (WIB)
        </h3>

        <div className="grid grid-cols-2 gap-2">
          {activeSchedules.map((slot) => {
            const passed = isTimePassed(slot.start_time);

            return (
              <button
                key={slot.id}
                disabled={passed}
                onClick={() => setSelectedSchedule(slot)}
                className={`flex flex-col items-center px-3 py-2 rounded-xl border-2 transition-all relative overflow-hidden ${
                  selectedSchedule?.id === slot.id
                    ? "bg-cyan-500 text-white border-cyan-500 shadow-md shadow-cyan-100"
                    : passed
                    ? "bg-gray-50 text-gray-300 border-gray-100 cursor-not-allowed"
                    : "bg-white text-gray-700 border-gray-100 hover:border-cyan-200 hover:bg-cyan-50"
                }`}
              >
                <span
                  className={`text-sm font-bold ${
                    passed ? "line-through" : ""
                  }`}
                >
                  {slot.start_time.slice(0, 5)} - {slot.end_time.slice(0, 5)}
                </span>

                <span
                  className={`text-[10px] mt-1 ${
                    selectedSchedule?.id === slot.id
                      ? "text-cyan-100"
                      : "text-gray-400"
                  }`}
                >
                  {passed ? "Time Passed" : `Available: ${slot.max_quota}`}
                </span>

                {passed && (
                  <div
                    className="absolute inset-0 bg-gray-50/10 opacity-50 pointer-events-none"
                    style={{
                      backgroundImage:
                        "linear-gradient(45deg, #ccc 25%, transparent 25%, transparent 50%, #ccc 50%, #ccc 75%, transparent 75%, transparent)",
                      backgroundSize: "10px 10px",
                    }}
                  ></div>
                )}
              </button>
            );
          })}
        </div>
      </div>

      <div className="mt-4 flex items-start gap-2 p-3 bg-amber-50 rounded-lg border border-amber-100">
        <AlertCircle className="w-4 h-4 text-amber-600 mt-0.5 flex-shrink-0" />
        <p className="text-[10px] text-amber-700 leading-relaxed">
          Time slots that have already passed for today cannot be re-selected.
          Please select a future time or another day.
        </p>
      </div>

      <Link
        href={
          selectedSchedule
            ? `/booking/order?schedule_id=${selectedSchedule.id}&doctor_id=${id}&date=${currentDate}&time=${selectedSchedule.start_time}&type=${selectedSchedule.type}`
            : "#"
        }
        onClick={(e) => !selectedSchedule && e.preventDefault()}
        tabIndex={!selectedSchedule ? -1 : 0}
        aria-disabled={!selectedSchedule}
        className={`w-full mt-6 flex justify-center items-center font-bold py-3 rounded-xl transition-all shadow-lg shadow-slate-200 
    ${
      !selectedSchedule
        ? "bg-gray-200 text-gray-400 cursor-not-allowed pointer-events-none shadow-none"
        : "bg-slate-800 text-white hover:bg-slate-700 active:scale-[0.98]"
    }`}
      >
        Proceed to Payment
      </Link>

      {isCalenderOpen && (
        <Calender
          currentDate={selectedDate}
          setCurrentDate={setSelectedDate}
          setIsCalendarOpen={setIsCalenderOpen}
          threshold={6}
        />
      )}
    </div>
  );
};

"use client";
import { useState } from "react";
import { Calendar, Clock } from "lucide-react";

interface BookingWidgetProps {
  fee: number;
}

const availableTimes = ["09:00", "10:00", "11:00", "14:00", "15:00", "16:00"];

const TimeButton = ({
  time,
  isSelected,
  onSelect,
}: {
  time: string;
  isSelected: boolean;
  onSelect: (time: string) => void;
}) => {
  return (
    <button
      key={time}
      onClick={() => onSelect(time)}
      className={`p-2 rounded-md text-sm font-semibold border-2 transition-colors ${
        isSelected
          ? "bg-cyan-500 text-white border-cyan-500"
          : "bg-cyan-50 text-cyan-800 border-cyan-200 hover:bg-cyan-100"
      }`}
    >
      {time}
    </button>
  );
};

export const BookingWidget = ({ fee }: BookingWidgetProps) => {
  const [selectedTime, setSelectedTime] = useState<string | null>(null);

  return (
    <div className="p-6 bg-white rounded-xl border border-gray-200 shadow-lg">
      {/* Fee Section */}
      <div className="flex justify-between items-center mb-4">
        <p className="text-lg font-semibold text-gray-700">Biaya Konsultasi</p>
        <p className="text-2xl font-bold text-cyan-600">
          Rp {fee.toLocaleString("id-ID")}
        </p>
      </div>

      {/* Calendar Section */}
      <div className="mt-6">
        <h3 className="font-semibold text-gray-800 mb-3 flex items-center">
          <Calendar className="w-5 h-5 mr-2 text-gray-500" />
          Pilih Jadwal
        </h3>
        {/* Placeholder untuk kalender */}
        <p className="text-center p-4 bg-gray-50 rounded-md text-gray-700 font-medium">
          Rabu, 20 September 2025
        </p>
      </div>

      {/* Time Section */}
      <div className="mt-6">
        <h3 className="font-semibold text-gray-800 mb-3 flex items-center">
          <Clock className="w-5 h-5 mr-2 text-gray-500" />
          Pilih Waktu (WIB)
        </h3>
        <div className="grid grid-cols-3 gap-2">
          {availableTimes.map((time) => (
            <TimeButton
              key={time}
              time={time}
              isSelected={selectedTime === time}
              onSelect={setSelectedTime}
            />
          ))}
        </div>
      </div>

      <button
        disabled={!selectedTime}
        className="w-full mt-8 bg-slate-800 text-white font-bold py-3 rounded-lg hover:bg-slate-700 transition-colors duration-200 disabled:bg-gray-300 disabled:cursor-not-allowed"
      >
        Booking Janji Temu
      </button>
    </div>
  );
};

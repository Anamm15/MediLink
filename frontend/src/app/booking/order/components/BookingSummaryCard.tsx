import Image from "next/image";
import { Calendar, Clock, Video } from "lucide-react";
import { DoctorProfileResponse } from "@/types/doctor.type";

interface SummaryProps {
  doctor: DoctorProfileResponse;
  date: string;
  time: string;
  type: string;
}

export const BookingSummaryCard = ({
  doctor,
  date,
  time,
  type,
}: SummaryProps) => {
  return (
    <div className="p-6 bg-white rounded-xl border border-gray-200 shadow-lg">
      <h2 className="text-xl font-bold text-gray-800 pb-4 border-b border-gray-200">
        Booking Summary
      </h2>
      <div className="flex items-center gap-4 py-5">
        <Image
          src={"https://i.pravatar.cc/150?u=adinda"}
          alt={doctor.name}
          width={64}
          height={64}
          className="rounded-full object-cover"
        />
        <div>
          <h3 className="font-bold text-gray-800">{doctor.name}</h3>
          <p className="text-sm text-gray-500">{doctor.specialization}</p>
        </div>
      </div>
      <div className="space-y-3 border-t border-dashed pt-5 text-sm">
        <div className="flex justify-between items-center">
          <span className="flex items-center text-gray-500">
            <Calendar className="w-4 h-4 mr-2" /> Date
          </span>
          <span className="font-semibold text-gray-800">{date}</span>
        </div>
        <div className="flex justify-between items-center">
          <span className="flex items-center text-gray-500">
            <Clock className="w-4 h-4 mr-2" /> Time
          </span>
          <span className="font-semibold text-gray-800">{time} (WIB)</span>
        </div>
        <div className="flex justify-between items-center">
          <span className="flex items-center text-gray-500">
            <Video className="w-4 h-4 mr-2" /> Type
          </span>
          <span className="font-semibold text-gray-800">{type}</span>
        </div>
      </div>
    </div>
  );
};

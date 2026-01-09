import Image from "next/image";
import { Calendar, Clock, Video } from "lucide-react";
import { DoctorDetail } from "@/types/index.type";

interface SummaryProps {
  doctor: DoctorDetail;
  booking: {
    date: string;
    time: string;
    type: "Video Call" | "Chat" | "Onsite";
  };
}

export const BookingSummaryCard = ({ doctor, booking }: SummaryProps) => {
  return (
    <div className="p-6 bg-white rounded-xl border border-gray-200 shadow-lg">
      <h2 className="text-xl font-bold text-gray-800 pb-4 border-b border-gray-200">
        Ringkasan Booking
      </h2>
      <div className="flex items-center gap-4 py-5">
        <Image
          src={doctor.avatarUrl}
          alt={doctor.name}
          width={64}
          height={64}
          className="rounded-full object-cover"
        />
        <div>
          <h3 className="font-bold text-gray-800">{doctor.name}</h3>
          <p className="text-sm text-gray-500">{doctor.specialty}</p>
        </div>
      </div>
      <div className="space-y-3 border-t border-dashed pt-5 text-sm">
        <div className="flex justify-between items-center">
          <span className="flex items-center text-gray-500">
            <Calendar className="w-4 h-4 mr-2" /> Tanggal
          </span>
          <span className="font-semibold text-gray-800">{booking.date}</span>
        </div>
        <div className="flex justify-between items-center">
          <span className="flex items-center text-gray-500">
            <Clock className="w-4 h-4 mr-2" /> Waktu
          </span>
          <span className="font-semibold text-gray-800">
            {booking.time} (WIB)
          </span>
        </div>
        <div className="flex justify-between items-center">
          <span className="flex items-center text-gray-500">
            <Video className="w-4 h-4 mr-2" /> Tipe
          </span>
          <span className="font-semibold text-gray-800">{booking.type}</span>
        </div>
      </div>
    </div>
  );
};

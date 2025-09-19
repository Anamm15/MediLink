import Image from "next/image";
import { Star, Calendar } from "lucide-react";
import { Doctor } from "@/types";

interface DoctorCardProps {
  doctor: Doctor;
}

export const DoctorCard = ({ doctor }: DoctorCardProps) => {
  return (
    <div className="bg-white rounded-xl border border-gray-200 shadow-sm hover:shadow-xl hover:-translate-y-1 transition-all duration-300 ease-in-out p-5 flex flex-col">
      <div className="flex items-start space-x-4">
        <div className="relative flex-shrink-0">
          <Image
            src={doctor.avatarUrl}
            alt={doctor.name}
            width={80}
            height={80}
            className="rounded-full object-cover border-2 border-gray-100"
          />
          {doctor.isOnline && (
            <span className="absolute bottom-0 right-0 block h-4 w-4 rounded-full bg-green-400 border-2 border-white ring-2 ring-green-400">
              <span className="sr-only">Online</span>
            </span>
          )}
        </div>
        <div className="flex-grow">
          <h3 className="text-lg font-bold text-gray-800">{doctor.name}</h3>
          <p className="text-sm font-semibold text-cyan-600">
            {doctor.specialty}
          </p>
          <div className="flex items-center mt-2 space-x-1">
            <Star className="w-4 h-4 text-yellow-500 fill-current" />
            <span className="text-sm text-gray-600 font-semibold">
              {doctor.rating.toFixed(1)}
            </span>
            <span className="text-sm text-gray-400">
              ({doctor.reviews} ulasan)
            </span>
          </div>
        </div>
      </div>

      <div className="border-t border-gray-100 my-4"></div>

      <div className="flex items-center text-sm text-gray-500 mt-auto">
        <Calendar className="w-4 h-4 mr-2 text-gray-400" />
        <span>
          Tersedia:{" "}
          <span className="font-semibold text-gray-700">
            {doctor.nextAvailable}
          </span>
        </span>
      </div>

      <button className="mt-5 w-full bg-cyan-500 text-white font-semibold py-2.5 rounded-lg hover:bg-cyan-600 transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-cyan-500">
        Lihat Profil & Jadwal
      </button>
    </div>
  );
};

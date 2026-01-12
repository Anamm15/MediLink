import Image from "next/image";
import { Star, Briefcase, Users } from "lucide-react";
import { DoctorProfileResponse } from "@/types/doctor.type";

interface HeaderProps {
  doctor: DoctorProfileResponse;
}

export const DoctorProfileHeader = ({ doctor }: HeaderProps) => {
  return (
    <div className="flex flex-col sm:flex-row items-start gap-6 p-6 bg-white rounded-xl border border-gray-200 shadow-sm">
      <div className="relative flex-shrink-0">
        <Image
          src={doctor.avatar_url}
          alt={doctor.name}
          width={128}
          height={128}
          className="rounded-full object-cover border-4 border-white shadow-md"
        />
      </div>
      <div className="flex-1">
        <h1 className="text-3xl font-bold text-gray-900">{doctor.name}</h1>
        <p className="mt-1 text-lg font-semibold text-cyan-600">
          {doctor.specialization}
        </p>
        <div className="mt-4 flex flex-wrap items-center gap-x-6 gap-y-2 text-gray-600">
          <div className="flex items-center gap-1.5">
            <Star className="w-5 h-5 text-yellow-500 fill-current" />
            <span className="font-semibold">
              {doctor.rating_total.toFixed(1)}
            </span>
            <span className="text-sm text-gray-500">
              ({doctor.review_count} reviews)
            </span>
          </div>
          <div className="flex items-center gap-1.5">
            <Briefcase className="w-5 h-5 text-gray-400" />
            <span className="font-semibold">
              {doctor.experience_years} years
            </span>
            <span className="text-sm text-gray-500">experience</span>
          </div>
          <div className="flex items-center gap-1.5">
            <Users className="w-5 h-5 text-gray-400" />
            <span className="font-semibold">{doctor.review_count}+</span>
            <span className="text-sm text-gray-500">patient</span>
          </div>
        </div>
      </div>
    </div>
  );
};

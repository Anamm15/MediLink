import Image from "next/image";
import { Star, Briefcase, Users } from "lucide-react";
import { DoctorDetail } from "@/types/index.type";

interface HeaderProps {
  doctor: DoctorDetail;
}

export const DoctorProfileHeader = ({ doctor }: HeaderProps) => {
  return (
    <div className="flex flex-col sm:flex-row items-start gap-6 p-6 bg-white rounded-xl border border-gray-200 shadow-sm">
      <div className="relative flex-shrink-0">
        <Image
          src={doctor.avatarUrl}
          alt={doctor.name}
          width={128}
          height={128}
          className="rounded-full object-cover border-4 border-white shadow-md"
        />
        {doctor.isOnline && (
          <span className="absolute bottom-2 right-2 block h-5 w-5 rounded-full bg-green-400 border-2 border-white ring-2 ring-green-400" />
        )}
      </div>
      <div className="flex-1">
        <h1 className="text-3xl font-bold text-gray-900">{doctor.name}</h1>
        <p className="mt-1 text-lg font-semibold text-cyan-600">
          {doctor.specialty}
        </p>
        <div className="mt-4 flex flex-wrap items-center gap-x-6 gap-y-2 text-gray-600">
          <div className="flex items-center gap-1.5">
            <Star className="w-5 h-5 text-yellow-500 fill-current" />
            <span className="font-semibold">{doctor.rating.toFixed(1)}</span>
            <span className="text-sm text-gray-500">
              ({doctor.reviews} ulasan)
            </span>
          </div>
          <div className="flex items-center gap-1.5">
            <Briefcase className="w-5 h-5 text-gray-400" />
            <span className="font-semibold">
              {doctor.yearsOfExperience} tahun
            </span>
            <span className="text-sm text-gray-500">pengalaman</span>
          </div>
          <div className="flex items-center gap-1.5">
            <Users className="w-5 h-5 text-gray-400" />
            <span className="font-semibold">{doctor.patientCount}+</span>
            <span className="text-sm text-gray-500">pasien</span>
          </div>
        </div>
      </div>
    </div>
  );
};

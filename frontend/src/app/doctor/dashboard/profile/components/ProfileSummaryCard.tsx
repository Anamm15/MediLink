import Image from "next/image";
import {
  Star,
  Users,
  Briefcase,
  Mail,
  Phone,
  Edit,
  MapPin,
  ShieldCheck,
} from "lucide-react";
import { DoctorProfileResponse } from "@/types/doctor.type";
import { DEFAULT_PROFILE } from "@/helpers/constant";

type ProfileSummaryCardProps = {
  doctorData: DoctorProfileResponse;
};

export const ProfileSummaryCard = ({ doctorData }: ProfileSummaryCardProps) => {
  return (
    <div className="bg-white p-6 rounded-xl shadow-sm border border-gray-200">
      <div className="text-center">
        <Image
          src={doctorData.avatar_url || DEFAULT_PROFILE}
          alt={doctorData.name}
          width={120}
          height={120}
          className="rounded-full ring-4 ring-cyan-100 mx-auto"
        />
        <h2 className="mt-4 text-2xl font-bold text-gray-800">
          {doctorData.name}
        </h2>
        <p className="mt-1 text-cyan-600 font-semibold">
          {doctorData.specialization}
        </p>
      </div>

      <div className="grid grid-cols-3 gap-4 w-full text-sm my-6">
        <div className="text-center p-2 rounded-lg bg-slate-50">
          <p className="font-bold text-lg text-gray-800">
            {doctorData.rating_total}
          </p>
          <p className="text-gray-500 text-xs">Rating</p>
        </div>
        <div className="text-center p-2 rounded-lg bg-slate-50">
          <p className="font-bold text-lg text-gray-800">
            {doctorData.review_count}
          </p>
          <p className="text-gray-500 text-xs">Patient</p>
        </div>
        <div className="text-center p-2 rounded-lg bg-slate-50">
          <p className="font-bold text-lg text-gray-800">
            {doctorData.experience_years} year(s)
          </p>
          <p className="text-gray-500 text-xs">Experience</p>
        </div>
      </div>

      <div className="w-full border-t border-gray-200"></div>

      {/* Bagian Bawah: Informasi Detail */}
      <div className="mt-6 space-y-4 text-left">
        <h3 className="text-base font-semibold text-gray-700">
          Contact & Practice Information
        </h3>

        <div className="flex items-start gap-3">
          <Mail className="w-5 h-5 text-gray-400 mt-0.5 flex-shrink-0" />
          <p className="text-sm font-medium text-gray-600 break-all">
            {doctorData.email}
          </p>
        </div>

        <div className="flex items-start gap-3">
          <Phone className="w-5 h-5 text-gray-400 mt-0.5 flex-shrink-0" />
          <p className="text-sm font-medium text-gray-600">
            {doctorData.phone_number}
          </p>
        </div>

        {doctorData.clinic &&
          doctorData.clinic.map((clinic) => (
            <div key={clinic.id} className="flex items-start gap-3">
              <MapPin className="w-5 h-5 text-gray-400 mt-0.5 flex-shrink-0" />
              <div>
                <p className="text-sm font-medium text-gray-800">
                  {clinic.name}
                </p>
                <p className="text-xs text-gray-500">{clinic.address}</p>
              </div>
            </div>
          ))}

        <div className="flex items-start gap-3">
          <ShieldCheck className="w-5 h-5 text-gray-400 mt-0.5 flex-shrink-0" />
          <div>
            <p className="text-sm font-medium text-gray-800">STR Number</p>
            <p className="text-xs text-gray-500">{doctorData.license_number}</p>
          </div>
        </div>
      </div>

      <button className="mt-8 w-full flex items-center justify-center gap-2 py-2.5 px-4 text-sm font-semibold rounded-lg bg-slate-800 text-white hover:bg-slate-700">
        <Edit className="w-4 h-4" /> Edit Detail Profil
      </button>
    </div>
  );
};

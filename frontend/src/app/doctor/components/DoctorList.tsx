import { DoctorCard } from "@/components/cards/DoctorCard";
import { DoctorProfileResponse } from "@/types/doctor.type";

interface DoctorListProps {
  doctors: DoctorProfileResponse[];
}

export const DoctorList = ({ doctors }: DoctorListProps) => {
  if (doctors.length === 0) {
    return <p className="text-center text-gray-500 mt-10">Doctor not found</p>;
  }

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      {doctors.map((doctor) => (
        <DoctorCard key={doctor.id} doctor={doctor} />
      ))}
    </div>
  );
};

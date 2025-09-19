import { DoctorCard } from "@/components/cards/DoctorCard";
import { Doctor } from "@/types";

interface DoctorListProps {
  doctors: Doctor[];
}

export const DoctorList = ({ doctors }: DoctorListProps) => {
  if (doctors.length === 0) {
    return (
      <p className="text-center text-gray-500 mt-10">Dokter tidak ditemukan.</p>
    );
  }

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      {doctors.map((doctor) => (
        <DoctorCard key={doctor.id} doctor={doctor} />
      ))}
    </div>
  );
};

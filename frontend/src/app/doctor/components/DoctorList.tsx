import { DoctorCard } from "@/components/cards/DoctorCard";
import { DefaultPagination } from "@/components/ui/pagination/DefaultPagination";
import { DoctorSearchResponse } from "@/types/doctor.type";

interface DoctorListProps {
  doctors: DoctorSearchResponse;
  page: number;
  setPage: React.Dispatch<React.SetStateAction<number>>;
}

export const DoctorList = ({ doctors, page, setPage }: DoctorListProps) => {
  const { metadata, data } = doctors;
  const { total_pages } = metadata;

  if (data.length === 0) {
    return <p className="text-center text-gray-500 mt-10">Doctor not found</p>;
  }

  return (
    <>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {data.map((doctor) => (
          <DoctorCard key={doctor.id} doctor={doctor} />
        ))}
      </div>

      {total_pages > 1 && (
        <DefaultPagination
          page={page}
          totalPages={total_pages}
          onPageChange={setPage}
          siblingCount={3}
        />
      )}
    </>
  );
};

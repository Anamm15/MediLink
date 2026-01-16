import { DEFAULT_LIMIT_QUERY, DEFAULT_PAGE_QUERY } from "@/helpers/constant";
import { useState } from "react";
import { usePatientAppointmentsQuery } from "../hooks/usePatientAppointment";
import { UserAppointmentCard } from "./UserAppointmentCard";
import { DefaultPagination } from "@/components/ui/pagination/DefaultPagination";
import { Spinner } from "@/components/ui/Spinner";

interface HistoryTabProps {
  patientId: string;
}

export default function HistoryTab({ patientId }: HistoryTabProps) {
  const [page, setPage] = useState(DEFAULT_PAGE_QUERY);
  const { data, isLoading } = usePatientAppointmentsQuery(
    patientId,
    page,
    DEFAULT_LIMIT_QUERY,
    "past"
  );
  const appointments = data?.data ?? [];
  const metadata = data?.metadata;

  if (isLoading) {
    return (
      <div className="flex h-full w-full justify-center items-center mt-40">
        <Spinner />
      </div>
    );
  }

  return (
    <div className="space-y-4">
      {appointments.map((app) => (
        <UserAppointmentCard key={app.id} appointment={app} />
      ))}

      {appointments.length === 0 && (
        <div className="text-center py-12">
          <p className="text-gray-500">
            You don not have any past appointment.
          </p>
        </div>
      )}

      {metadata && metadata.total_pages > 1 && (
        <DefaultPagination
          page={page}
          onPageChange={setPage}
          totalPages={metadata.total_pages}
          siblingCount={3}
        />
      )}
    </div>
  );
}

import { DEFAULT_LIMIT_QUERY, DEFAULT_PAGE_QUERY } from "@/helpers/constant";
import { useState } from "react";
import { DefaultPagination } from "@/components/ui/pagination/DefaultPagination";
import { Spinner } from "@/components/ui/Spinner";
import { useDoctorAppointmentsQuery } from "../hooks/useDoctorAppointment";
import { AppointmentResponse } from "@/types/appointment.type";
import { motion } from "framer-motion";
import { formatIDDate } from "@/helpers/datetime";
import { AppointmentCard } from "@/components/cards/AppointmentCard";
import { groupAppointmentsByDate } from "../page";

interface UpcomingTabProps {
  doctorId: string;
}

export default function UpcomingTab({ doctorId }: UpcomingTabProps) {
  const [page, setPage] = useState(DEFAULT_PAGE_QUERY);
  const { data, isLoading } = useDoctorAppointmentsQuery(
    doctorId,
    page,
    DEFAULT_LIMIT_QUERY,
    "upcoming"
  );
  const appointments = data?.data ?? [];
  const groupedAppointments = groupAppointmentsByDate(appointments);
  const metadata = data?.metadata;

  if (isLoading) {
    return (
      <div className="flex h-full w-full justify-center items-center mt-40">
        <Spinner />
      </div>
    );
  }

  return (
    <div className="space-y-8">
      {Object.entries(groupedAppointments).map(([date, apps], index) => (
        <motion.div
          key={date}
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5, delay: index * 0.1 }}
        >
          <h2 className="text-lg font-bold text-gray-800 mb-4 pb-2 border-b-2 border-cyan-500 inline-block">
            {formatIDDate(date)}
          </h2>
          <div className="space-y-4">
            {apps.map((app, appIndex) => (
              <motion.div
                key={app.id}
                initial={{ opacity: 0, x: -20 }}
                animate={{ opacity: 1, x: 0 }}
                transition={{
                  duration: 0.5,
                  delay: appIndex * 0.1 + index * 0.1,
                }}
              >
                <AppointmentCard
                  appointment={app}
                  isUpcoming={false}
                  isActionable={true}
                />
              </motion.div>
            ))}
          </div>
        </motion.div>
      ))}

      {appointments.length === 0 && (
        <div className="text-center py-12">
          <p className="text-gray-500">
            You don not have any upcoming appointment.
          </p>
        </div>
      )}

      {metadata && metadata.total_pages > 1 && (
        <DefaultPagination
          page={page}
          onPageChange={setPage}
          totalPages={metadata.total_pages}
          siblingCount={2}
        />
      )}
    </div>
  );
}

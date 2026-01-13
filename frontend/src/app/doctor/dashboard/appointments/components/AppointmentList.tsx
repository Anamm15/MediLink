"use client";
import { motion } from "framer-motion";
import { AppointmentCard } from "@/components/cards/AppointmentCard";
import { AppointmentDetailResponse } from "@/types/appointment.type";

const groupAppointmentsByDate = (appointments: AppointmentDetailResponse[]) => {
  const groups = appointments.reduce((acc, app) => {
    (acc[app.appointment_date] = acc[app.appointment_date] || []).push(app);
    return acc;
  }, {} as Record<string, AppointmentDetailResponse[]>);
  return groups;
};

export const AppointmentList = ({
  appointments,
}: {
  appointments: AppointmentDetailResponse[];
}) => {
  const groupedAppointments = groupAppointmentsByDate(appointments);

  if (appointments.length === 0) {
    return (
      <p className="text-center text-gray-500 mt-10">No appointments yet.</p>
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
            {date}
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
    </div>
  );
};

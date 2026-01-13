"use client";

import AppointmentDetail from "@/components/dashboard/shared/AppointmentDetail";
import { useAppointment } from "@/hooks/useAppointment";
import { useParams } from "next/navigation";

export default function AppointmentDetailPage() {
  const { id } = useParams();
  const { data: appointment } = useAppointment(id as string);
  if (!appointment) return null;
  return (
    <AppointmentDetail
      data={appointment}
      urlBack="/user/dashboard/appointments"
    />
  );
}

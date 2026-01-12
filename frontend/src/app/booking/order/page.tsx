"use client";

import { BookingSummaryCard } from "./components/BookingSummaryCard";
import { PaymentFlow } from "./components/PaymentFlow";
import { Navbar } from "@/components/layout/Navbar";
import { useDoctorQuery, useScheduleQuery } from "@/hooks/useDoctor";
import { useSearchParams } from "next/navigation";
import { CreateBookingRequest } from "@/types/appointment.type";

export default function OrderPage() {
  const searchParams = useSearchParams();
  const doctor_id = searchParams.get("doctor_id");
  const schedule_id = searchParams.get("schedule_id");
  const date = searchParams.get("date");
  const time = searchParams.get("time");
  const type = searchParams.get("type");
  const payload: CreateBookingRequest = {
    doctor_id: doctor_id as string,
    schedule_id: schedule_id as string,
    appointment_date: date as string,
  };

  const { data: doctor } = useDoctorQuery(doctor_id as string);
  const { data: schedule } = useScheduleQuery(schedule_id as string);

  return (
    <>
      <Navbar />
      <main className="bg-slate-50 min-h-screen">
        <div className="container mx-auto px-4 py-8 max-w-5xl md:py-12">
          <header className="mb-8">
            <h1 className="text-3xl md:text-4xl font-bold text-gray-800">
              Confirmation & Payment
            </h1>
            <p className="mt-2 text-md text-gray-500">
              Please double-check your appointment details before proceeding to
              payment.
            </p>
          </header>

          <div className="grid grid-cols-1 lg:grid-cols-5 gap-8 lg:gap-8 items-start">
            <div className="lg:col-span-2 gap-6">
              {doctor && date && time && type && (
                <BookingSummaryCard
                  doctor={doctor}
                  date={date}
                  time={time}
                  type={type}
                />
              )}
            </div>
            <div className="lg:col-span-3 gap-6">
              {schedule && (
                <PaymentFlow
                  consultationFee={schedule?.consultation_fee}
                  payload={payload}
                />
              )}
            </div>
          </div>
        </div>
      </main>
    </>
  );
}

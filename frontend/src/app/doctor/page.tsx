"use client";

import { Navbar } from "@/components/layout/Navbar";
import { DoctorList } from "./components/DoctorList";
import { SearchFilter } from "@/components/ui/SearchFilter";
import { useSearchDoctor } from "@/hooks/useDoctor";
import { useState } from "react";
import { DOCTOR_LIMIT_QUERY } from "@/helpers/constant";

export default function DoctorsPage() {
  const [doctorNameFiltered, setDoctorNameFiltered] = useState<string>("");
  const [page, setPage] = useState(1);

  const { data: doctors } = useSearchDoctor(
    doctorNameFiltered,
    page,
    DOCTOR_LIMIT_QUERY
  );

  return (
    <>
      <Navbar />
      <main className="bg-slate-50 min-h-screen">
        <div className="container mx-auto px-4 py-8 md:py-12">
          <header className="mb-8 text-center md:text-left">
            <h1 className="text-3xl md:text-4xl font-bold text-gray-800">
              Find Your Doctor
            </h1>
            <p className="mt-2 text-md text-gray-500">
              Book an appointment with the professional doctor of your choice
              easily.
            </p>
          </header>

          <SearchFilter setDoctorNameFiltered={setDoctorNameFiltered} />
          {doctors && (
            <DoctorList doctors={doctors} page={page} setPage={setPage} />
          )}
        </div>
      </main>
    </>
  );
}

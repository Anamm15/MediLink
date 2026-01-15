"use client";

import { DoctorProfileHeader } from "./components/DoctorProfileHeader";
import { InfoSection } from "./components/InfoSection";
import { BookingWidget } from "./components/BookingWidget";
import { Navbar } from "@/components/layout/Navbar";
import { useDoctorQuery } from "@/hooks/useDoctor";
import { useParams } from "next/navigation";
import { CheckCircle } from "lucide-react";

export default function DoctorDetailPage() {
  const params = useParams();
  const { data: doctor } = useDoctorQuery(params.id as string);

  return (
    <>
      <Navbar />
      <main className="bg-slate-50 min-h-screen">
        <div className="container mx-auto px-4 py-8 md:py-12">
          <div className="grid grid-cols-1 lg:grid-cols-3 gap-8 lg:gap-12 items-start">
            {doctor && (
              <div className="lg:col-span-2 flex flex-col gap-6">
                <DoctorProfileHeader doctor={doctor} />
                <InfoSection title="Tentang Dokter">
                  <p>{doctor.bio}</p>
                </InfoSection>
                <InfoSection title="Pendidikan">
                  <ul>
                    {doctor.education?.map((edu, index) => (
                      <li key={index} className="flex items-start gap-3 mb-2">
                        <CheckCircle className="w-5 h-5 text-cyan-500 mt-1 flex-shrink-0" />
                        <div>
                          <span className="font-semibold">{edu.degree}</span> -{" "}
                          {edu.institution} ({edu.year})
                        </div>
                      </li>
                    ))}
                  </ul>
                </InfoSection>
              </div>
            )}

            {/* Kolom Kanan: Widget Booking (Sticky) */}
            <div className="lg:sticky top-32">
              <BookingWidget />
            </div>
          </div>
        </div>
      </main>
    </>
  );
}

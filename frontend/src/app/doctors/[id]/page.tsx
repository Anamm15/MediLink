import { DoctorProfileHeader } from "./components/DoctorProfileHeader";
import { InfoSection } from "./components/InfoSection";
import { BookingWidget } from "./components/BookingWidget";
import { DoctorDetail } from "@/types";
import { CheckCircle } from "lucide-react";

// Placeholder data untuk satu dokter. Nantinya akan di-fetch berdasarkan `params.id`
const mockDoctorDetail: DoctorDetail = {
  id: "1",
  name: "Dr. Adinda Melati, Sp.A",
  specialty: "Dokter Anak",
  avatarUrl: "https://i.pravatar.cc/150?u=adinda",
  rating: 4.9,
  reviews: 128,
  isOnline: true,
  nextAvailable: "Besok, 10:00 WIB",
  yearsOfExperience: 12,
  patientCount: 5000,
  bio: "Dr. Adinda Melati adalah seorang Dokter Spesialis Anak lulusan Universitas Indonesia. Beliau memiliki pengalaman lebih dari 12 tahun dalam menangani berbagai masalah kesehatan anak, mulai dari bayi baru lahir hingga remaja. Beliau aktif mengikuti seminar dan workshop untuk terus memperbarui ilmunya demi memberikan pelayanan terbaik.",
  consultationFee: 150000,
  clinic: {
    name: "Klinik Sehat Ceria",
    address: "Jl. Merdeka No. 123, Jakarta Pusat",
  },
  education: [
    {
      degree: "Spesialis Anak (Sp.A)",
      university: "Universitas Indonesia",
      year: "2012",
    },
    {
      degree: "Dokter Umum",
      university: "Universitas Gadjah Mada",
      year: "2008",
    },
  ],
  experience: [
    {
      position: "Dokter Anak",
      hospital: "RSUPN Dr. Cipto Mangunkusumo",
      period: "2013 - Sekarang",
    },
    {
      position: "Dokter Umum",
      hospital: "RSUD Tangerang",
      period: "2009 - 2012",
    },
  ],
  reviewsList: [], // Bisa diisi nanti
};

// params akan berisi { id: '...' } dari URL
export default function DoctorDetailPage({
  params,
}: {
  params: { id: string };
}) {
  const doctor = mockDoctorDetail; // Di aplikasi nyata: const doctor = await fetchDoctorById(params.id);

  return (
    <main className="bg-slate-50 min-h-screen">
      <div className="container mx-auto px-4 py-8 md:py-12">
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8 lg:gap-12 items-start">
          {/* Kolom Kiri: Informasi Dokter */}
          <div className="lg:col-span-2 flex flex-col gap-6">
            <DoctorProfileHeader doctor={doctor} />
            <InfoSection title="Tentang Dokter">
              <p>{doctor.bio}</p>
            </InfoSection>
            <InfoSection title="Pendidikan">
              <ul>
                {doctor.education.map((edu, index) => (
                  <li key={index} className="flex items-start gap-3 mb-2">
                    <CheckCircle className="w-5 h-5 text-cyan-500 mt-1 flex-shrink-0" />
                    <div>
                      <span className="font-semibold">{edu.degree}</span> -{" "}
                      {edu.university} ({edu.year})
                    </div>
                  </li>
                ))}
              </ul>
            </InfoSection>
            <InfoSection title="Pengalaman Kerja">
              <ul>
                {doctor.experience.map((exp, index) => (
                  <li key={index} className="flex items-start gap-3 mb-2">
                    <CheckCircle className="w-5 h-5 text-cyan-500 mt-1 flex-shrink-0" />
                    <div>
                      <span className="font-semibold">{exp.position}</span> di{" "}
                      {exp.hospital} ({exp.period})
                    </div>
                  </li>
                ))}
              </ul>
            </InfoSection>
          </div>

          {/* Kolom Kanan: Widget Booking (Sticky) */}
          <div className="lg:sticky top-24">
            <BookingWidget fee={doctor.consultationFee} />
          </div>
        </div>
      </div>
    </main>
  );
}

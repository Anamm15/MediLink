import { Navbar } from "@/components/layout/Navbar";
import { DoctorList } from "./components/DoctorList";
import { SearchFilter } from "@/components/ui/SearchFilter";
import { Doctor } from "@/types";

// Data Placeholder (nantinya ini akan datang dari API)
const mockDoctors: Doctor[] = [
  {
    id: "1",
    name: "Dr. Adinda Melati, Sp.A",
    specialty: "Dokter Anak",
    avatarUrl: "https://i.pravatar.cc/150?u=adinda",
    rating: 4.9,
    reviews: 128,
    isOnline: true,
    nextAvailable: "Besok, 10:00 WIB",
  },
  {
    id: "2",
    name: "Dr. Budi Santoso, Sp.PD",
    specialty: "Penyakit Dalam",
    avatarUrl: "https://i.pravatar.cc/150?u=budi",
    rating: 4.8,
    reviews: 210,
    isOnline: false,
    nextAvailable: "Jumat, 14:30 WIB",
  },
  {
    id: "3",
    name: "Dr. Citra Ayu, Sp.KK",
    specialty: "Kulit & Kelamin",
    avatarUrl: "https://i.pravatar.cc/150?u=citra",
    rating: 5.0,
    reviews: 95,
    isOnline: true,
    nextAvailable: "Hari ini, 19:00 WIB",
  },
  {
    id: "4",
    name: "Dr. Dian Permana, Sp.OG",
    specialty: "Kandungan & Ginekologi",
    avatarUrl: "https://i.pravatar.cc/150?u=dian",
    rating: 4.9,
    reviews: 154,
    isOnline: true,
    nextAvailable: "Besok, 09:00 WIB",
  },
  {
    id: "5",
    name: "Dr. Eko Prasetyo, Sp.M",
    specialty: "Dokter Mata",
    avatarUrl: "https://i.pravatar.cc/150?u=eko",
    rating: 4.7,
    reviews: 88,
    isOnline: false,
    nextAvailable: "Lusa, 11:00 WIB",
  },
  {
    id: "6",
    name: "Dr. Fira Lestari, Sp.THT-KL",
    specialty: "THT",
    avatarUrl: "https://i.pravatar.cc/150?u=fira",
    rating: 4.8,
    reviews: 112,
    isOnline: true,
    nextAvailable: "Hari ini, 20:00 WIB",
  },
];

export default function DoctorsPage() {
  return (
    <>
      <Navbar />
      <main className="bg-slate-50 min-h-screen">
        <div className="container mx-auto px-4 py-8 md:py-12">
          <header className="mb-8 text-center md:text-left">
            <h1 className="text-3xl md:text-4xl font-bold text-gray-800">
              Temukan Dokter Anda
            </h1>
            <p className="mt-2 text-md text-gray-500">
              Booking janji temu dengan dokter profesional pilihan Anda dengan
              mudah.
            </p>
          </header>

          <SearchFilter />
          <DoctorList doctors={mockDoctors} />
        </div>
      </main>
    </>
  );
}

import { ProfileSummaryCard } from "./components/ProfileSummaryCard";
import { ProfileTabs } from "./components/ProfileTabs";

// --- DATA DUMMY UNTUK DOKTER ---
const mockDoctorData = {
  name: "Dr. Adinda Melati, Sp.A",
  specialty: "Dokter Anak",
  strNumber: "1234567890123456",
  avatarUrl: "https://i.pravatar.cc/150?u=adinda",
  rating: 4.9,
  patientCount: "5K+",
  experience: 12,
  // --- TAMBAHKAN DATA BARU DI SINI ---
  email: "adinda.melati@healthapp.com",
  phone: "0812-3456-7890",
  clinic: {
    name: "Klinik Sehat Ceria",
    address: "Jl. Merdeka No. 123, Jakarta Pusat",
  },
  // --- AKHIR DATA BARU ---
  bio: "Saya adalah seorang Dokter Spesialis Anak dengan pengalaman lebih dari 12 tahun. Saya berdedikasi untuk memberikan perawatan kesehatan terbaik bagi anak-anak, mulai dari bayi baru lahir hingga remaja. Saya percaya bahwa pendekatan yang ramah dan komunikatif adalah kunci dalam membangun kepercayaan dengan pasien dan orang tua.",
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
};
export default function DoctorProfilePage() {
  return (
    <div className="grid grid-cols-1 lg:grid-cols-3 gap-6 items-start">
      {/* Kolom Kiri */}
      <div className="lg:col-span-1 lg:sticky top-24">
        <ProfileSummaryCard doctorData={mockDoctorData} />
      </div>

      {/* Kolom Kanan */}
      <div className="lg:col-span-2">
        <ProfileTabs doctorData={mockDoctorData} />
      </div>
    </div>
  );
}

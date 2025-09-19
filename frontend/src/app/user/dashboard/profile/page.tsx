import { ProfileHeader } from "./components/ProfileHeader";
import { EditableSectionCard } from "../components/EditableSectionCard";
// --- DATA DUMMY UNTUK PENGGUNA ---
const mockUserData = {
  name: "Budi Setiawan",
  email: "budi.setiawan@email.com",
  avatarUrl: "https://i.pravatar.cc/150?u=budi",
  dob: "15 Agustus 1988",
  gender: "Pria",
  phone: "0812-3456-7890",
  address: "Jl. Pahlawan No. 45, Surabaya, Jawa Timur",
  medicalInfo: {
    bloodType: "O+",
    allergies: "Debu, Makanan Laut",
    chronicConditions: "Asma",
  },
  emergencyContact: {
    name: "Sarah Setiawan",
    relation: "Istri",
    phone: "0812-9876-5432",
  },
};

// Komponen kecil untuk menampilkan item data agar rapi
const InfoItem = ({ label, value }: { label: string; value: string }) => (
  <div>
    <p className="text-sm text-gray-500">{label}</p>
    <p className="font-semibold text-gray-800">{value}</p>
  </div>
);

export default function UserProfilePage() {
  const { personalInfo, medicalInfo, emergencyContact } = {
    personalInfo: {
      "Nama Lengkap": mockUserData.name,
      "Tanggal Lahir": mockUserData.dob,
      "Jenis Kelamin": mockUserData.gender,
      "Nomor Telepon": mockUserData.phone,
      Alamat: mockUserData.address,
    },
    medicalInfo: {
      "Golongan Darah": mockUserData.medicalInfo.bloodType,
      Alergi: mockUserData.medicalInfo.allergies,
      "Penyakit Kronis": mockUserData.medicalInfo.chronicConditions,
    },
    emergencyContact: {
      Nama: mockUserData.emergencyContact.name,
      Hubungan: mockUserData.emergencyContact.relation,
      "Nomor Telepon": mockUserData.emergencyContact.phone,
    },
  };

  return (
    <div className="space-y-6">
      <header>
        <h1 className="text-3xl font-bold text-gray-800">Profil Saya</h1>
        <p className="mt-1 text-gray-500">
          Kelola informasi akun dan data pribadi Anda di sini.
        </p>
      </header>

      <ProfileHeader userData={mockUserData} />

      <EditableSectionCard title="Data Diri">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          {Object.entries(personalInfo).map(([label, value]) => (
            <InfoItem key={label} label={label} value={value} />
          ))}
        </div>
      </EditableSectionCard>

      <EditableSectionCard title="Informasi Medis Penting">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          {Object.entries(medicalInfo).map(([label, value]) => (
            <InfoItem key={label} label={label} value={value} />
          ))}
        </div>
      </EditableSectionCard>

      <EditableSectionCard title="Kontak Darurat">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          {Object.entries(emergencyContact).map(([label, value]) => (
            <InfoItem key={label} label={label} value={value} />
          ))}
        </div>
      </EditableSectionCard>
    </div>
  );
}

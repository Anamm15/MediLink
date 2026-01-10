"use client";

import { useEffect, useState } from "react";
import Image from "next/image";
import {
  KeyRound,
  Camera,
  User as UserIcon,
  ShieldCheck,
  ShieldAlert,
  Edit3,
  MailWarning,
  Stethoscope,
  Plus,
} from "lucide-react";
import { EditableSectionCard } from "../components/EditableSectionCard";
import { ProfileHeader } from "./components/ProfileHeader";
import { AccountInformation } from "./components/Account";
import DetailPatient from "./components/Patient";
import {
  TypographyH3,
  TypographyH4,
  TypographyP,
} from "@/components/ui/Typography";
import { UserProfileResponse } from "@/types/user.type";
import { useUserQuery } from "./hooks/useUser";

const mockInitialData: UserProfileResponse = {
  user: {
    id: "user-01",
    name: "Budi Setiawan",
    email: "budi.setiawan@email.com",
    role: "Pasien",
    status: "active",
    phone_number: "0812-3456-7890",
    avatar_url: "https://i.pravatar.cc/150?u=budi",
    is_verified: true,
    created_at: "2022-01-01",
    updated_at: "2022-01-01",
  },
  // Default: null untuk user baru
  patient: null,
  // patient: {
  //   id: "patient-01",
  //   identity_number: "1234567890",
  //   weight_kg: 70,
  //   height_cm: 170,
  //   birth_date: "15 Agustus 1988",
  //   gender: "Pria",
  //   blood_type: "O+",
  //   allergies: "Debu, Makanan Laut",
  //   history_chronic_diseases: "Asma",
  //   emergency_contact: "0812-9876-5432",
  //   insurance_number: "213702312",
  //   insurance_provider: "BPJS Kesehatan",
  //   occupation: "PNS",
  // },
};

export default function UserProfilePage() {
  const { data: userData, isLoading } = useUserQuery();
  const [data, setData] = useState<UserProfileResponse>();

  useEffect(() => {
    setData(userData);
  }, [userData]);

  return (
    <div className="max-w-5xl mx-auto space-y-8 pb-10">
      <header className="flex flex-col md:flex-row md:items-end justify-between gap-4">
        <div>
          <TypographyH3> My Profile </TypographyH3>
          <TypographyP className="text-gray-500">
            Manage your personal identity information and medical records.
          </TypographyP>
        </div>
      </header>

      {!isLoading && data && (
        <>
          <ProfileHeader user={data.user} />

          <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
            <AccountInformation data={data} setData={setData} />
            <DetailPatient data={data} setData={setData} />
          </div>
        </>
      )}
    </div>
  );
}

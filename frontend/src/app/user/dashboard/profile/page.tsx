"use client";

import { useEffect, useState } from "react";
import { ProfileHeader } from "./components/ProfileHeader";
import { AccountInformation } from "./components/Account";
import DetailPatient from "./components/Patient";
import { TypographyH3, TypographyP } from "@/components/ui/Typography";
import { UserProfileResponse } from "@/types/user.type";
import { useUserQuery } from "./hooks/useUser";

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

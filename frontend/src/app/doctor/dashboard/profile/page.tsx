"use client";

import { ProfileSummaryCard } from "./components/ProfileSummaryCard";
import { ProfileTabs } from "./components/ProfileTabs";
import { useDoctorQuery } from "./hooks/useDoctorProfile";

export default function DoctorProfilePage() {
  const { data: doctor } = useDoctorQuery();

  return (
    <div className="grid grid-cols-1 lg:grid-cols-3 gap-6 items-start">
      <div className="lg:col-span-1 lg:sticky top-24">
        {doctor && <ProfileSummaryCard doctorData={doctor} />}
      </div>

      <div className="lg:col-span-2">
        {doctor && <ProfileTabs doctorData={doctor} />}
      </div>
    </div>
  );
}

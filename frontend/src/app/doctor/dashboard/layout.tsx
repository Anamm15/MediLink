import { DashboardHeader } from "@/components/dashboard/DashboardHeader";
import { DoctorSidebar } from "@/components/dashboard/DoctorSidebar";

export default function DoctorDashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="flex min-h-screen w-full bg-gray-50">
      <div className="sticky h-screen top-0 left-0">
        <DoctorSidebar />
      </div>
      <div className="flex flex-1 flex-col">
        <DashboardHeader />
        <main className="flex-1 p-6">{children}</main>
      </div>
    </div>
  );
}

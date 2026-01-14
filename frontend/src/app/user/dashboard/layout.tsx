import { DashboardHeader } from "@/components/dashboard/DashboardHeader";
import { UserSidebar } from "@/components/dashboard/UserSidebar";

export default function UserDashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="flex min-h-screen w-full bg-gray-50">
      <div className="sticky h-screen top-0 left-0">
        <UserSidebar />
      </div>
      <div className="flex flex-1 flex-col">
        <DashboardHeader />
        <main className="flex-1 p-6">{children}</main>
      </div>
    </div>
  );
}

import Link from "next/link";
import { SidebarMenuItem } from "./SidebarMenuItem";
import {
  User,
  Settings,
  ClipboardPlus,
  CalendarDays,
  History,
  LogOut,
  ShieldCheck,
  HeartPulse,
} from "lucide-react";

const doctorMenuItems = [
  {
    href: "/doctor/dashboard/profile",
    text: "Profile",
    icon: <User className="h-4 w-4" />,
  },
  {
    href: "/doctor/dashboard/appointments",
    text: "Appointments",
    icon: <CalendarDays className="h-4 w-4" />,
  },
  {
    href: "/doctor/dashboard/medical-records",
    text: "Medical Records",
    icon: <HeartPulse className="h-4 w-4" />,
  },
  {
    href: "/doctor/dashboard/prescriptions",
    text: "Prescriptions",
    icon: <ClipboardPlus className="h-4 w-4" />,
  },
  {
    href: "/doctor/dashboard/setting",
    text: "Setting",
    icon: <Settings className="h-4 w-4" />,
  },
];

export const DoctorSidebar = () => {
  return (
    <aside className="hidden w-64 sticky flex-col border-r bg-white md:flex h-screen top-0">
      <div className="flex h-16 items-center border-b px-6">
        <Link
          href="/"
          className="flex items-center gap-2 font-semibold text-cyan-600"
        >
          <ShieldCheck className="h-6 w-6" />
          <span>HealthApp (Doctor)</span>
        </Link>
      </div>
      <div className="flex-1 overflow-auto py-4">
        <nav className="grid items-start px-4 text-sm font-medium">
          <ul className="space-y-1">
            {doctorMenuItems.map((item) => (
              <SidebarMenuItem key={item.href} {...item} />
            ))}
          </ul>
        </nav>
      </div>
      <div className="p-4 border-t">
        <SidebarMenuItem
          href="/logout"
          text="Logout"
          icon={<LogOut className="h-4 w-4" />}
        />
      </div>
    </aside>
  );
};

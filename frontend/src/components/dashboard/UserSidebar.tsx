import Link from "next/link";
import { SidebarMenuItem } from "./SidebarMenuItem";
import {
  User,
  Settings,
  ClipboardPlus,
  CalendarDays,
  HeartPulse,
  CreditCard,
  LogOut,
  ShieldCheck,
  Bell,
} from "lucide-react";

const userMenuItems = [
  {
    href: "/user/dashboard/profile",
    text: "Profile",
    icon: <User className="h-4 w-4" />,
  },
  {
    href: "/user/dashboard/appointments",
    text: "Appointment",
    icon: <CalendarDays className="h-4 w-4" />,
  },
  {
    href: "/user/dashboard/prescriptions",
    text: "Prescriptions",
    icon: <ClipboardPlus className="h-4 w-4" />,
  },
  {
    href: "/user/dashboard/medical-records",
    text: "Medical Records",
    icon: <HeartPulse className="h-4 w-4" />,
  },
  {
    href: "/user/dashboard/payment-history",
    text: "Payment History",
    icon: <CreditCard className="h-4 w-4" />,
  },
  {
    href: "/user/dashboard/notifications",
    text: "Notification",
    icon: <Bell className="h-4 w-4" />,
  },
  {
    href: "/user/dashboard/setting",
    text: "Setting",
    icon: <Settings className="h-4 w-4" />,
  },
];

export const UserSidebar = () => {
  return (
    <aside className="hidden w-64 flex-col border-r bg-white md:flex">
      <div className="flex h-16 items-center border-b px-6">
        <Link
          href="/"
          className="flex items-center gap-2 font-semibold text-cyan-600"
        >
          <ShieldCheck className="h-6 w-6" />
          <span>HealthApp</span>
        </Link>
      </div>
      <div className="flex-1 overflow-auto py-4">
        <nav className="grid items-start px-4 text-sm font-medium">
          <ul className="space-y-1">
            {userMenuItems.map((item) => (
              <SidebarMenuItem key={item.href} {...item} />
            ))}
          </ul>
        </nav>
      </div>
      <div className="mt-auto p-4 border-t">
        <SidebarMenuItem
          href="/logout"
          text="Logout"
          icon={<LogOut className="h-4 w-4" />}
        />
      </div>
    </aside>
  );
};

"use client";

import { DEFAULT_PROFILE } from "@/helpers/constant";
import { useSession } from "@/hooks/useAuth";
import { Bell } from "lucide-react";
import Image from "next/image";

export const DashboardHeader = () => {
  const { data: user } = useSession();

  return (
    <header className="sticky top-0 z-10 flex h-16 items-center gap-4 border-b bg-white px-6">
      <h1 className="text-xl font-semibold text-gray-800">Dashboard</h1>
      <div className="ml-auto flex items-center gap-4">
        <button className="p-2 rounded-full hover:bg-gray-100 relative">
          <Bell className="h-5 w-5 text-gray-500" />
          <span className="absolute top-1 right-1 block h-2 w-2 rounded-full bg-red-500 ring-2 ring-white"></span>
        </button>
        <div className="flex items-center gap-3">
          <Image
            src={user?.avatar_url || DEFAULT_PROFILE}
            alt="User Avatar"
            width={36}
            height={36}
            className="rounded-full"
          />
          <span className="text-sm font-medium text-gray-700 hidden sm:block">
            {user?.name}
          </span>
        </div>
      </div>
    </header>
  );
};

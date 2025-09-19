"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import React from "react";

interface SidebarMenuItemProps {
  href: string;
  text: string;
  icon: React.ReactNode;
}

export const SidebarMenuItem = ({ href, text, icon }: SidebarMenuItemProps) => {
  const pathname = usePathname();
  const isActive = pathname === href;

  return (
    <li>
      <Link
        href={href}
        className={`flex items-center gap-3 rounded-lg px-3 py-2 transition-all ${
          isActive
            ? "bg-cyan-100 text-cyan-700 font-semibold"
            : "text-gray-500 hover:bg-gray-100 hover:text-gray-900"
        }`}
      >
        {icon}
        {text}
      </Link>
    </li>
  );
};

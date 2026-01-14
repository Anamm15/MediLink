"use client";

import { useState, useEffect } from "react";
import Link from "next/link";
import Image from "next/image";
import {
  ShieldCheck,
  Menu,
  X,
  LogOut,
  LayoutDashboard,
  ChevronDown,
} from "lucide-react";
import { usePathname } from "next/navigation";
import { useLogout, useSession } from "@/hooks/useAuth";
import { DEFAULT_PROFILE } from "@/helpers/constant";

const menuItems = [
  { text: "Home", href: "/" },
  { text: "Find Doctor", href: "/doctor" },
  { text: "Pharmacy", href: "/pharmacy" },
  { text: "Articles", href: "/articles" },
];

const mapRoleToDashboard = (role: string): string => {
  if (role === "doctor") {
    return "/doctor/dashboard/profile";
  } else if (role === "pharmacy") {
    return "/pharmacy/dashboard/profile";
  } else if (role === "user" || role === "patient") {
    return "/user/dashboard/profile";
  } else if (role === "admin") {
    return "/admin/dashboard/profile";
  }
  return "/";
};

export const Navbar = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const [isScrolled, setIsScrolled] = useState(false);
  const pathname = usePathname();
  const { data: user } = useSession();
  const { mutate: logout } = useLogout();
  const [dashboardLinkHref, setDashboardLinkHref] = useState("/");

  useEffect(() => {
    const handleScroll = () => {
      setIsScrolled(window.scrollY > 10);
    };
    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  useEffect(() => {
    if (isMenuOpen) {
      setIsMenuOpen(false);
    }
  }, [pathname]);

  useEffect(() => {
    if (!user) return;
    setDashboardLinkHref(mapRoleToDashboard(user?.role || ""));
  }, [user]);

  return (
    <header
      className={`sticky top-0 z-50 transition-all duration-300 ${
        isScrolled ? "bg-white/80 shadow-md backdrop-blur-sm" : "bg-white"
      }`}
    >
      <div className="container mx-auto flex h-20 items-center justify-between px-4">
        <Link
          href="/"
          className="flex items-center gap-2 font-bold text-cyan-600"
        >
          <ShieldCheck className="h-7 w-7" />
          <span className="text-xl font-bold">MediLink</span>
        </Link>

        <nav className="hidden items-center gap-8 md:flex">
          {menuItems.map((item) => (
            <Link
              key={item.text}
              href={item.href}
              className={`font-semibold transition-colors hover:text-cyan-600 ${
                pathname === item.href ? "text-cyan-600" : "text-gray-600"
              }`}
            >
              {item.text}
            </Link>
          ))}
        </nav>

        <div className="hidden items-center gap-3 md:flex">
          {user ? (
            <div className="group relative">
              <Link
                href={dashboardLinkHref}
                className="flex items-center gap-2 cursor-pointer"
              >
                <Image
                  src={user.avatar_url || DEFAULT_PROFILE}
                  alt={user.name}
                  width={36}
                  height={36}
                  className="rounded-full"
                />
                <span className="font-medium text-gray-700">{user.name}</span>
                <ChevronDown className="h-4 w-4 text-gray-700" />
              </Link>
              <div className="absolute right-0 top-full mt-2 w-48 origin-top-right rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all duration-300">
                <div className="py-1">
                  <Link
                    href={dashboardLinkHref}
                    className="flex w-full items-center gap-2 px-4 py-2 text-gray-700 hover:bg-gray-100"
                  >
                    <LayoutDashboard className="h-4 w-4" /> Dashboard
                  </Link>
                  <button
                    onClick={() => logout()}
                    className="flex w-full items-center gap-2 px-4 py-2 text-red-600 hover:bg-gray-100"
                  >
                    <LogOut className="h-4 w-4" /> Logout
                  </button>
                </div>
              </div>
            </div>
          ) : (
            <>
              <Link
                href="/login"
                className="px-4 py-2 font-semibold text-gray-600 hover:text-cyan-600"
              >
                Sign In
              </Link>
              <Link
                href="/register"
                className="rounded-full bg-cyan-500 px-5 py-2 font-semibold text-white shadow-sm hover:bg-cyan-600"
              >
                Sign Up
              </Link>
            </>
          )}
        </div>

        {/* Hamburger Button */}
        <div className="md:hidden">
          <button onClick={() => setIsMenuOpen(!isMenuOpen)}>
            {isMenuOpen ? (
              <X className="h-6 w-6" />
            ) : (
              <Menu className="h-6 w-6" />
            )}
          </button>
        </div>
      </div>

      <div
        className={`absolute top-20 left-0 w-full bg-white shadow-lg md:hidden transition-transform duration-300 ease-in-out ${
          isMenuOpen ? "translate-x-0" : "-translate-x-full"
        }`}
      >
        <nav className="flex flex-col space-y-4 p-6">
          {menuItems.map((item) => (
            <Link
              key={item.text}
              href={item.href}
              className={`rounded-md px-4 py-2 font-semibold transition-colors hover:bg-cyan-50 ${
                pathname === item.href
                  ? "bg-cyan-50 text-cyan-600"
                  : "text-gray-700"
              }`}
            >
              {item.text}
            </Link>
          ))}
          <div className="border-t border-gray-200 pt-6">
            {user ? (
              <div className="space-y-4">
                <Link
                  href={dashboardLinkHref}
                  className="flex items-center gap-3"
                >
                  <Image
                    src={user.avatar_url || DEFAULT_PROFILE}
                    alt={user.name}
                    width={40}
                    height={40}
                    className="rounded-full"
                  />
                  <div>
                    <p className="font-semibold text-gray-800">{user.name}</p>
                    <p className="text-xs text-gray-500">View Dashboard</p>
                  </div>
                </Link>
                <button
                  onClick={() => logout()}
                  className="w-full rounded-md bg-red-50 py-2.5 text-sm font-semibold text-red-600"
                >
                  Logout
                </button>
              </div>
            ) : (
              <div className="flex flex-col space-y-3">
                <Link
                  href="/login"
                  className="w-full rounded-md border border-gray-300 py-2.5 text-center text-sm font-semibold text-gray-700"
                >
                  Sign In
                </Link>
                <Link
                  href="/register"
                  className="w-full rounded-md bg-cyan-500 py-2.5 text-center text-sm font-semibold text-white"
                >
                  Sign Up
                </Link>
              </div>
            )}
          </div>
        </nav>
      </div>
    </header>
  );
};

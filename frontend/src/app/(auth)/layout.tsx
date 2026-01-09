"use client";

import Link from "next/link";
import { motion } from "framer-motion";
import { ShieldCheck } from "lucide-react";
import {
  TypographyH4,
  TypographyLead,
  TypographyP,
  TypographySmall,
} from "@/components/ui/Typography";

export default function AuthLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <main className="min-h-screen flex items-center justify-center bg-slate-50 px-4">
      {/* Animation */}
      <motion.div
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.5 }}
        className="w-full max-w-md"
      >
        <div className="bg-white p-8 rounded-2xl shadow-xl border border-gray-100">
          {/* Header */}
          <div className="text-center mb-8">
            <Link
              href="/"
              className="inline-flex items-center gap-2 text-cyan-600"
            >
              <ShieldCheck className="h-10 w-10" />
              <span className="">
                <TypographyLead className="font-extrabold tracking-tight text-cyan-600">
                  HealthApp
                </TypographyLead>
              </span>
            </Link>
            <TypographyH4>Welcome Back</TypographyH4>
          </div>

          {/* Form Login */}
          {children}
        </div>

        {/* Support Link */}
        <TypographySmall className="text-center mt-5 text-gray-400">
          Need Help?{" "}
          <Link href="/support" className="underline">
            Contact Support
          </Link>
        </TypographySmall>
      </motion.div>
    </main>
  );
}

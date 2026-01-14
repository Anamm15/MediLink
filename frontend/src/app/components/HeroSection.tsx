"use client";
import { motion } from "framer-motion";
import Image from "next/image";
import Link from "next/link";
import { ArrowRight } from "lucide-react";

export const HeroSection = () => {
  return (
    <section className="bg-slate-50">
      <div className="container mx-auto flex justify-center items-center px-4 py-20 md:py-32 gap-12">
        <motion.div
          initial={{ opacity: 0, x: -50 }}
          animate={{ opacity: 1, x: 0 }}
          transition={{ duration: 0.8 }}
          className="max-w-[700px]"
        >
          <h1 className="text-4xl md:text-5xl font-bold text-gray-800 leading-tight">
            Digital Health Solutions, Right at Your Fingertips.
          </h1>
          <p className="mt-4 text-lg text-gray-500">
            Get professional doctor consultations, digital prescriptions, and
            pharmacy delivery services anytime, anywhere.
          </p>
          <div className="mt-8 flex flex-wrap gap-4">
            <Link
              href="/doctor"
              className="flex items-center gap-2 rounded-full bg-cyan-500 px-8 py-3.5 text-base font-semibold text-white shadow-lg hover:bg-cyan-600 transition-transform hover:scale-105"
            >
              Find a Doctor <ArrowRight className="w-5 h-5" />
            </Link>
            <Link
              href="/#how-it-works"
              className="rounded-full px-8 py-3.5 text-base font-semibold text-gray-700 hover:bg-gray-100"
            >
              How It Works
            </Link>
          </div>
        </motion.div>

        <motion.div
          initial={{ opacity: 0, scale: 0.8 }}
          animate={{ opacity: 1, scale: 1 }}
          transition={{ duration: 0.8, delay: 0.2 }}
          className="relative"
        >
          <Image
            src="/images/home/hero.avif"
            alt="hero"
            width={600}
            height={600}
            className="rounded-xl object-contain"
          />
        </motion.div>
      </div>
    </section>
  );
};

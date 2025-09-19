"use client";
import { motion } from "framer-motion";
import Link from "next/link";

const Wave = () => (
  <div className="absolute top-0 left-0 w-full overflow-hidden leading-[0]">
    <svg
      data-name="Layer 1"
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 1200 120"
      preserveAspectRatio="none"
      className="relative block h-[150px] w-full"
    >
      <path
        d="M321.39,56.44c58-10.79,114.16-30.13,172-41.86,82.39-16.72,168.19-17.73,250.45-.39C823.78,31,906.67,72,985.66,92.83c70.05,18.48,146.53,26.09,214.34,3V0H0V27.35A600.21,600.21,0,0,0,321.39,56.44Z"
        className="fill-slate-50"
      ></path>
    </svg>
  </div>
);

export const CtaSection = () => {
  return (
    <section className="relative bg-cyan-600 pt-32 pb-20 text-white">
      <Wave />
      <div className="container mx-auto px-4 text-center relative z-10">
        <motion.h2
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true, amount: 0.5 }}
          transition={{ duration: 0.5 }}
          className="text-3xl md:text-4xl font-bold"
        >
          Siap Mengambil Kendali Kesehatan Anda?
        </motion.h2>
        <motion.p
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true, amount: 0.5 }}
          transition={{ duration: 0.5, delay: 0.2 }}
          className="mt-4 max-w-xl mx-auto text-cyan-100"
        >
          Bergabunglah dengan ribuan pengguna lainnya dan nikmati kemudahan
          akses layanan kesehatan profesional di mana pun Anda berada.
        </motion.p>
        <motion.div
          initial={{ opacity: 0, scale: 0.8 }}
          whileInView={{ opacity: 1, scale: 1 }}
          viewport={{ once: true, amount: 0.5 }}
          transition={{ duration: 0.5, delay: 0.4 }}
          className="mt-8"
        >
          <Link
            href="/register"
            className="rounded-full bg-white px-10 py-4 text-base font-semibold text-cyan-600 shadow-lg hover:bg-slate-100 transition-transform hover:scale-105 inline-block"
          >
            Daftar Gratis Sekarang
          </Link>
        </motion.div>
      </div>
    </section>
  );
};

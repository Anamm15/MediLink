"use client";
import { motion } from "framer-motion";
import Image from "next/image";
import { CheckCircle } from "lucide-react";

const features = [
  {
    title: "Konsultasi Video & Chat Tanpa Batas",
    description:
      "Bicaralah dengan dokter spesialis melalui video call atau chat. Dapatkan diagnosis awal dan saran medis profesional dari kenyamanan rumah Anda.",
    points: [
      "Pilih dokter sesuai spesialisasi",
      "Jadwalkan sesi secara fleksibel",
      "Privasi dan keamanan terjamin",
    ],
    imageUrl: "/feature-consultation.png", // Ganti dengan gambar relevan
  },
  {
    title: "Resep Digital yang Sah dan Cepat",
    description:
      "Setelah konsultasi, dokter dapat menerbitkan resep digital yang langsung tersimpan di akun Anda. Sah, aman, dan tanpa kertas.",
    points: [
      "Resep langsung tersedia",
      "Terintegrasi dengan apotek",
      "Hindari salah baca resep",
    ],
    imageUrl: "/feature-prescription.png", // Ganti dengan gambar relevan
  },
  {
    title: "Apotek Antar, Obat Sampai di Hari yang Sama",
    description:
      "Tebus resep atau beli produk kesehatan lainnya dari apotek terpercaya kami. Pesanan Anda akan kami antar dengan cepat dan aman.",
    points: [
      "Jangkauan pengiriman luas",
      "Stok obat terjamin",
      "Lacak pesanan secara real-time",
    ],
    imageUrl: "/feature-pharmacy.png", // Ganti dengan gambar relevan
  },
];

export const FeaturesSection = () => {
  return (
    <section className="py-20 bg-slate-50 overflow-hidden">
      <div className="container mx-auto px-4">
        <motion.h2
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true, amount: 0.5 }}
          transition={{ duration: 0.5 }}
          className="text-3xl font-bold text-gray-800 text-center mb-16"
        >
          Semua Kebutuhan Kesehatan dalam Satu Genggaman
        </motion.h2>

        <div className="space-y-20">
          {features.map((feature, index) => (
            <motion.div
              key={feature.title}
              initial={{ opacity: 0, x: index % 2 === 0 ? -100 : 100 }}
              whileInView={{ opacity: 1, x: 0 }}
              viewport={{ once: true, amount: 0.3 }}
              transition={{ duration: 0.8 }}
              className="grid md:grid-cols-2 gap-12 items-center"
            >
              {/* Kolom Teks */}
              <div className={index % 2 === 1 ? "md:order-last" : ""}>
                <h3 className="text-2xl font-bold text-gray-800">
                  {feature.title}
                </h3>
                <p className="mt-4 text-gray-500">{feature.description}</p>
                <ul className="mt-6 space-y-3">
                  {feature.points.map((point) => (
                    <li key={point} className="flex items-center gap-3">
                      <CheckCircle className="w-5 h-5 text-cyan-500 flex-shrink-0" />
                      <span className="text-gray-600">{point}</span>
                    </li>
                  ))}
                </ul>
              </div>

              {/* Kolom Gambar */}
              <div className="flex items-center justify-center">
                <Image
                  src={feature.imageUrl}
                  alt={feature.title}
                  width={500}
                  height={500}
                  className="rounded-xl object-contain shadow-2xl"
                />
              </div>
            </motion.div>
          ))}
        </div>
      </div>
    </section>
  );
};

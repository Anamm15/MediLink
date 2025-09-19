"use client";
import { motion } from "framer-motion";
import Image from "next/image";

const testimonials = [
  {
    name: "Sarah L.",
    role: "Ibu Rumah Tangga",
    comment:
      "Sangat membantu saat anak sakit tengah malam. Dokter responsif dan obat cepat sampai. Terima kasih HealthApp!",
    avatar: "https://i.pravatar.cc/150?u=sarah",
  },
  {
    name: "Rian D.",
    role: "Pekerja Kantoran",
    comment:
      "Tidak perlu izin kantor untuk konsultasi ringan. Efisien dan sangat profesional. Highly recommended!",
    avatar: "https://i.pravatar.cc/150?u=rian",
  },
  {
    name: "Anita P.",
    role: "Mahasiswi",
    comment:
      "Tebus resep kulit dari dokter jadi gampang banget. Nggak perlu antre di apotek lagi. Aplikasinya juga mudah dipakai.",
    avatar: "https://i.pravatar.cc/150?u=anita",
  },
];

export const TestimonialSection = () => {
  return (
    <section className="bg-slate-50 py-20">
      <div className="container mx-auto px-4">
        <motion.h2
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true, amount: 0.5 }}
          transition={{ duration: 0.5 }}
          className="text-3xl font-bold text-gray-800 text-center"
        >
          Apa Kata Mereka yang Sudah Menggunakan?
        </motion.h2>
        <div className="mt-12 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          {testimonials.map((t, i) => (
            <motion.div
              key={t.name}
              initial={{ opacity: 0, scale: 0.9 }}
              whileInView={{ opacity: 1, scale: 1 }}
              viewport={{ once: true, amount: 0.5 }}
              transition={{ duration: 0.5, delay: i * 0.1 }}
              className="flex flex-col p-6 bg-white rounded-lg shadow-md"
            >
              <p className="text-gray-600 flex-grow">"{t.comment}"</p>
              <div className="mt-4 flex items-center gap-4">
                <Image
                  src={t.avatar}
                  alt={t.name}
                  width={48}
                  height={48}
                  className="rounded-full"
                />
                <div>
                  <p className="font-semibold text-gray-800">{t.name}</p>
                  <p className="text-sm text-gray-500">{t.role}</p>
                </div>
              </div>
            </motion.div>
          ))}
        </div>
      </div>
    </section>
  );
};

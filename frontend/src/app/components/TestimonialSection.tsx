"use client";
import { motion } from "framer-motion";
import Image from "next/image";

const testimonials = [
  {
    name: "Sarah L.",
    role: "Homemaker",
    comment:
      "Extremely helpful when my child got sick late at night. The doctor was responsive and the medicine arrived quickly. Thank you, HealthApp!",
    avatar: "https://i.pravatar.cc/150?u=sarah",
  },
  {
    name: "Rian D.",
    role: "Office Worker",
    comment:
      "No need to take time off work for a minor consultation. Efficient and very professional. Highly recommended!",
    avatar: "https://i.pravatar.cc/150?u=rian",
  },
  {
    name: "Anita P.",
    role: "College Student",
    comment:
      "Redeeming a dermatology prescription from the doctor was super easy. No more waiting in line at the pharmacy. The app is also very easy to use.",
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
          What do those who have used it say?
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

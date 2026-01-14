"use client";
import { motion } from "framer-motion";
import { Users, Stethoscope, Star } from "lucide-react";

const stats = [
  {
    icon: <Users className="w-8 h-8 text-cyan-500" />,
    value: "10,000+",
    label: "Satisfied Patient",
  },
  {
    icon: <Stethoscope className="w-8 h-8 text-cyan-500" />,
    value: "500+",
    label: "Profesional Doctor",
  },
  {
    icon: <Star className="w-8 h-8 text-cyan-500" />,
    value: "4.9/5",
    label: "App Rating",
  },
];

export const StatsSection = () => {
  return (
    <section className="py-16 bg-white">
      <div className="container mx-auto px-4">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8 text-center">
          {stats.map((stat, index) => (
            <motion.div
              key={stat.label}
              initial={{ opacity: 0, y: 30 }}
              whileInView={{ opacity: 1, y: 0 }}
              viewport={{ once: true, amount: 0.5 }}
              transition={{ duration: 0.6, delay: index * 0.2 }}
              className="flex flex-col items-center"
            >
              {stat.icon}
              <p className="mt-2 text-4xl font-bold text-gray-800">
                {stat.value}
              </p>
              <p className="text-gray-500">{stat.label}</p>
            </motion.div>
          ))}
        </div>
      </div>
    </section>
  );
};

"use client";
import { motion } from "framer-motion";
import { Search, MessageSquare, Pill } from "lucide-react";

const steps = [
  {
    icon: <Search />,
    title: "1. Find a Doctor",
    description:
      "Find a specialist or general practitioner that matches your needs.",
  },
  {
    icon: <MessageSquare />,
    title: "2. Online Consultation",
    description: "Consult via chat or video call comfortably from your home.",
  },
  {
    icon: <Pill />,
    title: "3. Redeem Prescriptions & Medicines",
    description:
      "Receive digital prescriptions and have your medicines delivered directly to your location.",
  },
];

export const HowItWorksSection = () => {
  return (
    <section id="how-it-works" className="py-20">
      <div className="container mx-auto px-4 text-center">
        <motion.h2
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true, amount: 0.5 }}
          transition={{ duration: 0.5 }}
          className="text-3xl font-bold text-gray-800"
        >
          Health Made Easy in 3 Steps
        </motion.h2>
        <motion.p
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true, amount: 0.5 }}
          transition={{ duration: 0.5, delay: 0.2 }}
          className="mt-2 max-w-2xl mx-auto text-gray-500"
        >
          We simplify the healthcare process for you.
        </motion.p>

        <div className="mt-12 grid grid-cols-1 md:grid-cols-3 gap-8">
          {steps.map((step, index) => (
            <motion.div
              key={step.title}
              initial={{ opacity: 0, y: 50 }}
              whileInView={{ opacity: 1, y: 0 }}
              viewport={{ once: true, amount: 0.5 }}
              transition={{ duration: 0.5, delay: index * 0.2 }}
              className="flex flex-col items-center p-8 bg-white rounded-xl shadow-lg hover:shadow-cyan-100/50 hover:-translate-y-2 transition-all"
            >
              <div className="flex h-16 w-16 items-center justify-center rounded-full bg-cyan-100 text-cyan-600">
                {step.icon}
              </div>
              <h3 className="mt-6 text-xl font-bold text-gray-800">
                {step.title}
              </h3>
              <p className="mt-2 text-gray-500">{step.description}</p>
            </motion.div>
          ))}
        </div>
      </div>
    </section>
  );
};

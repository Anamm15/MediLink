"use client";
import { motion } from "framer-motion";
import Image from "next/image";
import { CheckCircle } from "lucide-react";

const features = [
  {
    title: "Unlimited Video & Chat Consultation",
    description:
      "Talk to specialist doctors via video call or chat. Get an initial diagnosis and professional medical advice from the comfort of your home.",
    points: [
      "Choose doctors by specialization",
      "Schedule sessions flexibly",
      "Guaranteed privacy and security",
    ],
    imageUrl: "/images/home/consultation.png",
  },
  {
    title: "Fast & Valid Digital Prescriptions",
    description:
      "After the consultation, doctors can issue digital prescriptions that are instantly saved to your account. Legal, secure, and paperless.",
    points: [
      "Prescriptions available instantly",
      "Integrated with pharmacies",
      "Avoid prescription misinterpretation",
    ],
    imageUrl: "/images/home/drug.png",
  },
  {
    title: "Same-Day Medicine Delivery from Partner Pharmacies",
    description:
      "Redeem prescriptions or purchase other health products from our trusted partner pharmacies. Your order will be delivered quickly and safely.",
    points: [
      "Wide delivery coverage",
      "Guaranteed medicine availability",
      "Real-time order tracking",
    ],
    imageUrl: "/images/home/pharmacy.png",
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
          All your health needs in one hand
        </motion.h2>

        <div className="space-y-20 max-w-6xl mx-auto">
          {features.map((feature, index) => (
            <motion.div
              key={feature.title}
              initial={{ opacity: 0, x: index % 2 === 0 ? -100 : 100 }}
              whileInView={{ opacity: 1, x: 0 }}
              viewport={{ once: true, amount: 0.3 }}
              transition={{ duration: 0.8 }}
              className="grid md:grid-cols-2 items-center"
            >
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

              <div className="flex items-center justify-center">
                <Image
                  src={feature.imageUrl}
                  alt={feature.title}
                  width={400}
                  height={400}
                  className="rounded-xl object-cover shadow-2xl transition-all duration-300 hover:scale-105"
                />
              </div>
            </motion.div>
          ))}
        </div>
      </div>
    </section>
  );
};

"use client";
import { motion } from "framer-motion";
import Link from "next/link";
import {
  Calendar,
  Pill,
  CreditCard,
  Megaphone,
  ArrowRight,
} from "lucide-react";
import React from "react";

// Tipe data untuk Notifikasi
type Notification = any;

const notificationConfig = {
  appointment_reminder: {
    icon: <Calendar className="w-5 h-5 text-blue-600" />,
    bgColor: "bg-blue-100",
  },
  prescription_ready: {
    icon: <Pill className="w-5 h-5 text-green-600" />,
    bgColor: "bg-green-100",
  },
  payment_success: {
    icon: <CreditCard className="w-5 h-5 text-purple-600" />,
    bgColor: "bg-purple-100",
  },
  promo: {
    icon: <Megaphone className="w-5 h-5 text-orange-600" />,
    bgColor: "bg-orange-100",
  },
};

export const NotificationItem = ({
  notification,
}: {
  notification: Notification;
}) => {
  const config =
    notificationConfig[notification.type as keyof typeof notificationConfig];

  return (
    <motion.div
      initial={{ opacity: 0, x: -20 }}
      animate={{ opacity: 1, x: 0 }}
      transition={{ duration: 0.4 }}
      className={`relative flex items-start gap-4 p-4 rounded-lg border transition-colors ${
        !notification.isRead
          ? "bg-cyan-50 border-cyan-200"
          : "bg-white border-gray-200 hover:bg-slate-50"
      }`}
    >
      {!notification.isRead && (
        <span className="absolute top-3 right-3 block h-2.5 w-2.5 rounded-full bg-cyan-500"></span>
      )}

      <div className={`flex-shrink-0 p-3 rounded-full ${config.bgColor}`}>
        {config.icon}
      </div>

      <div className="flex-grow">
        <p className="font-semibold text-gray-800">{notification.title}</p>
        <p className="text-sm text-gray-500">{notification.description}</p>
        {notification.action && (
          <Link
            href={notification.action.href}
            className="mt-2 inline-flex items-center gap-1 text-sm font-semibold text-cyan-600 hover:text-cyan-800"
          >
            {notification.action.text} <ArrowRight className="w-4 h-4" />
          </Link>
        )}
      </div>

      <p className="flex-shrink-0 text-xs text-gray-400">
        {notification.timestamp}
      </p>
    </motion.div>
  );
};

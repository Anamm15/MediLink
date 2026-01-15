import { AlertCircle, CheckCircle2 } from "lucide-react";
import React from "react";

export const statusConfig: Record<
  string,
  {
    label: string;
    color: string;
    icon: React.FC<React.SVGProps<SVGSVGElement>>;
  }
> = {
  confirmed: {
    label: "confirmed",
    color: "bg-cyan-100 text-cyan-700",
    icon: CheckCircle2,
  },
  pending: {
    label: "pending",
    color: "bg-amber-100 text-amber-700",
    icon: AlertCircle,
  },
  completed: {
    label: "completed",
    color: "bg-green-100 text-green-700",
    icon: CheckCircle2,
  },
  canceled: {
    label: "canceled",
    color: "bg-red-100 text-red-700",
    icon: AlertCircle,
  },
  in_progress: {
    label: "in progress",
    color: "bg-amber-100 text-amber-700",
    icon: AlertCircle,
  },
};

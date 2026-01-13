import {
  cancelAppointment,
  completeAppointment,
} from "@/services/appointment.service";
import { AppointmentDetailResponse } from "@/types/appointment.type";
import { useMutation } from "@tanstack/react-query";
import React from "react";
import { toast } from "sonner";

export function useCancelAppointment(
  appointment: AppointmentDetailResponse,
  setAppointment: React.Dispatch<
    React.SetStateAction<AppointmentDetailResponse>
  >
) {
  return useMutation({
    mutationFn: (id: string) => cancelAppointment(id),

    onMutate: () => {
      setAppointment({ ...appointment, status: "canceled" });
      const toastId = toast.loading("Canceling appointment...");
      return { toastId };
    },

    onSuccess: (data, variables, context) => {
      toast.success("Appointment canceled", {
        id: context?.toastId,
        duration: 3000,
      });
    },

    onError: (error, variables, context) => {
      toast.error("Failed to cancel appointment", {
        id: context?.toastId,
        duration: 3000,
      });
      console.log(error);
    },
  });
}

export function useCompleteAppointment(
  appointment: AppointmentDetailResponse,
  setAppointment: React.Dispatch<
    React.SetStateAction<AppointmentDetailResponse>
  >
) {
  return useMutation({
    mutationFn: (id: string) => completeAppointment(id),

    onMutate: () => {
      const toastId = toast.loading("Completing appointment...");
      return { toastId };
    },

    onSuccess: (data, variables, context) => {
      setAppointment({ ...appointment, status: "completed" });
      toast.success("Appointment completed", {
        id: context?.toastId,
        duration: 3000,
      });
    },

    onError: (error, variables, context) => {
      toast.error("Failed to complete appointment", {
        id: context?.toastId,
        duration: 3000,
      });
      console.log(error);
    },
  });
}

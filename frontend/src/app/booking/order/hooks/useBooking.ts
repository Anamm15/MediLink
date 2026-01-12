import { createAppointment } from "@/services/appointment.service";
import { useMutation } from "@tanstack/react-query";
import { toast } from "sonner";

export function useBooking() {
  return useMutation({
    mutationFn: createAppointment,

    onMutate: () => {
      const toastId = toast.loading("Verifying data...");
      return { toastId };
    },

    onSuccess: async (data, variables, context) => {
      try {
        toast.success("User onboarded successfully", {
          id: context?.toastId,
          duration: 3000,
        });
      } catch (error) {
        toast.error("Failed to refresh session, please re-login", {
          id: context?.toastId,
        });
      }
    },

    onError: (error, variables, context) => {
      toast.error("Failed to create patient", {
        id: context?.toastId,
        duration: 3000,
      });
      console.log(error);
    },
  });
}

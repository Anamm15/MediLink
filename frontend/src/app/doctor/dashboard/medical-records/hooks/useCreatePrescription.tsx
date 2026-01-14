import { createPrescription } from "@/services/prescription.service";
import { useMutation } from "@tanstack/react-query";
import { toast } from "sonner";

export function useCreatePrescription() {
  return useMutation({
    mutationFn: createPrescription,

    onMutate: () => {
      const toastId = toast.loading("Creating prescription...");
      return { toastId };
    },

    onSuccess: (data, variables, context) => {
      toast.success("Prescription created successfully", {
        id: context?.toastId,
        duration: 3000,
      });
    },

    onError: (error, variables, context) => {
      toast.error("Failed to create prescription", {
        id: context?.toastId,
        duration: 3000,
      });
      console.log(error);
    },
  });
}

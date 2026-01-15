import { createMedicalRecord } from "@/services/medical_record.service";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";

export function useCreateMedicalRecord(doctor_id: string) {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: createMedicalRecord,

    onMutate: () => {
      const toastId = toast.loading("Creating data...");
      return { toastId };
    },

    onSuccess: async (data, variables, context) => {
      queryClient.removeQueries({
        queryKey: ["doctor-medical-records", doctor_id],
      });
      toast.success("Medical record created successfully", {
        id: context?.toastId,
        duration: 3000,
      });
    },

    onError: (error, variables, context) => {
      toast.error("Failed to create medical record", {
        id: context?.toastId,
        duration: 3000,
      });
      console.log(error);
    },
  });
}

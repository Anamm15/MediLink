import { updatePatient } from "@/services/patient.service";
import { useMutation } from "@tanstack/react-query";
import { toast } from "sonner";

export function useUpdatePatient(
  setIsEditing: React.Dispatch<React.SetStateAction<boolean>>
) {
  return useMutation({
    mutationFn: updatePatient,

    onMutate: () => {
      const toastId = toast.loading("Updating...");
      return { toastId };
    },

    onSuccess: (data, variables, context) => {
      setIsEditing(false);
      toast.success("Profile updated successfully", {
        id: context?.toastId,
        duration: 3000,
      });
    },

    onError: (error, variables, context) => {
      toast.error("Failed to update profile", {
        id: context?.toastId,
        duration: 3000,
      });
      console.log(error);
    },
  });
}

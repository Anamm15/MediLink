import { register } from "@/services/auth.service";
import { useMutation } from "@tanstack/react-query";
import { toast } from "sonner";

export function useRegister() {
  return useMutation({
    mutationFn: register,

    onMutate: () => {
      const toastId = toast.loading("Loading...");
      return { toastId };
    },

    onSuccess: (data, variables, context) => {
      toast.success("Register successful", {
        id: context?.toastId,
        duration: 3000,
      });
    },

    onError: (error: any, variables, context) => {
      const message =
        error.response?.data?.message ||
        error.message ||
        "An error occurred while registering";
      toast.error(message, { id: context?.toastId });
    },
  });
}

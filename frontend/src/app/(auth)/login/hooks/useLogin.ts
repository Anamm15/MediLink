import { login } from "@/services/auth.service";
import { useMutation } from "@tanstack/react-query";
import { toast } from "sonner";

export function useLogin() {
  return useMutation({
    mutationFn: login,
    onMutate: () => {
      const toastId = toast.loading("Loading...");
      return { toastId };
    },

    onSuccess: (data, variables, context) => {
      console.log(data);

      localStorage.setItem("token", data);
      toast.success("Login successful", {
        id: context?.toastId,
        duration: 3000,
      });
    },

    onError: (error: any, variables, context) => {
      const message =
        error.response?.data?.message ||
        error.message ||
        "An error occurred while logging in";
      toast.error(message, { id: context?.toastId });
    },
  });
}

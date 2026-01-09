import { requestResetPassword, resetPassword } from "@/services/auth.service";
import { useMutation } from "@tanstack/react-query";
import { toast } from "sonner";
import { useRouter } from "next/navigation";

export function useRequestResetPassword() {
  return useMutation({
    mutationFn: requestResetPassword,

    onMutate: () => {
      const toastId = toast.loading("Loading...");
      return { toastId };
    },

    onSuccess: (data, variables, context) => {
      toast.success("OTP has been sent to your email", {
        id: context?.toastId,
        duration: 3000,
      });
    },

    onError: (error: any, variables, context) => {
      const message =
        error.response?.data?.message ||
        error.message ||
        "An error occurred while sending OTP";
      toast.error(message, { id: context?.toastId });
    },
  });
}

export function useResetPassword() {
  const router = useRouter();
  return useMutation({
    mutationFn: resetPassword,

    onMutate: () => {
      const toastId = toast.loading("Loading...");
      return { toastId };
    },

    onSuccess: (data, variables, context) => {
      toast.success("Password has been reset successfully", {
        id: context?.toastId,
        duration: 3000,
      });
      router.push("/login");
    },

    onError: (error: any, variables, context) => {
      const message =
        error.response?.data?.message ||
        error.message ||
        "An error occurred while resetting password";
      toast.error(message, { id: context?.toastId });
    },
  });
}

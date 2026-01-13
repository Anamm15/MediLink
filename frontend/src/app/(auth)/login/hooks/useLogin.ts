import { login } from "@/services/auth.service";
import { useMutation } from "@tanstack/react-query";
import { toast } from "sonner";
import { useRouter, useSearchParams } from "next/navigation";

export function useLogin() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const returnTo = searchParams.get("returnTo");

  return useMutation({
    mutationFn: login,
    onMutate: () => {
      const toastId = toast.loading("Loading...");
      return { toastId };
    },

    onSuccess: (data, variables, context) => {
      localStorage.setItem("token", data);
      toast.success("Login successful", {
        id: context?.toastId,
        duration: 3000,
      });

      if (returnTo && returnTo.startsWith("/")) {
        router.push(returnTo);
      } else {
        router.push("/");
      }
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

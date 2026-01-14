"use client";

import { logout } from "@/services/auth.service";
import { me } from "@/services/user.service";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useRouter } from "next/navigation";
import { toast } from "sonner";

export function useSession() {
  return useQuery({
    queryKey: ["me"],
    queryFn: me,
    staleTime: 24 * 60 * 60 * 1000,
  });
}

export function useLogout() {
  const queryClient = useQueryClient();
  const router = useRouter();

  return useMutation({
    mutationFn: logout,

    onMutate: () => {
      const toastId = toast.loading("Logging out...");
      return { toastId };
    },

    onSuccess: (data, variables, context) => {
      localStorage.removeItem("token");
      queryClient.clear();
      toast.success("Logout successful", {
        id: context?.toastId,
        duration: 3000,
      });
      router.push("/");
    },

    onError: (error, variables, context) => {
      toast.error("Failed to logout", {
        id: context?.toastId,
        duration: 3000,
      });
      console.log(error);
    },
  });
}

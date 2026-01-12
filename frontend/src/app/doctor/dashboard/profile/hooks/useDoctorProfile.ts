import { doctorMe, updateDoctor } from "@/services/doctor.service";
import {
  DoctorProfileResponse,
  DoctorUpdateRequest,
} from "@/types/doctor.type";
import { useQuery, useMutation } from "@tanstack/react-query";
import { toast } from "sonner";

export function useDoctorQuery() {
  return useQuery<DoctorProfileResponse>({
    queryKey: ["doctor"],
    queryFn: () => doctorMe(),
  });
}

export function useUpdateDoctor(
  setIsEditing: React.Dispatch<React.SetStateAction<boolean>>
) {
  return useMutation({
    mutationFn: updateDoctor,

    onMutate: () => {
      const toastId = toast.loading("Updating profile...");
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

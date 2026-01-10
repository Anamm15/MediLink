import { UserProfileResponse } from "@/types/user.type";
import { useMutation, useQuery } from "@tanstack/react-query";
import {
  getProfile,
  onBoardPatient,
  sendEmailVerification,
  updateProfile,
  verifyEmail,
} from "@/services/user.service";
import { toast } from "sonner";
import React from "react";
import { forceRefreshToken } from "@/helpers/auth";
import { setAuthToken } from "@/lib/api";

export function useUserQuery() {
  return useQuery<UserProfileResponse>({
    queryKey: ["profile"],
    queryFn: getProfile,
  });
}

export function useUpdateUser(
  setIsEditing: React.Dispatch<React.SetStateAction<boolean>>
) {
  return useMutation({
    mutationFn: updateProfile,

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

export function useOnBoardPatient(
  setIsModalOpen: React.Dispatch<React.SetStateAction<boolean>>,
  dataParam: UserProfileResponse,
  setData: React.Dispatch<React.SetStateAction<UserProfileResponse | undefined>>
) {
  return useMutation({
    mutationFn: onBoardPatient,

    onMutate: () => {
      const toastId = toast.loading("Verifying data...");
      return { toastId };
    },

    onSuccess: async (data, variables, context) => {
      try {
        setIsModalOpen(false);
        setData({
          ...dataParam,
          patient: data,
        });

        const newAccessToken = await forceRefreshToken();
        setAuthToken(newAccessToken);

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

export function useSendEmailVerification(
  setIsVerificationModalOpen: React.Dispatch<React.SetStateAction<boolean>>
) {
  return useMutation({
    mutationFn: sendEmailVerification,

    onMutate: () => {
      const toastId = toast.loading("Sending email...");
      return { toastId };
    },

    onSuccess: (data, variables, context) => {
      toast.success("Email has sent successfully", {
        id: context?.toastId,
        duration: 3000,
      });
      setIsVerificationModalOpen(true);
    },

    onError: (error, variables, context) => {
      toast.error("Failed to send email", {
        id: context?.toastId,
        duration: 3000,
      });
      console.log(error);
    },
  });
}

export function useVerifyEmail(
  setIsModalOpen: React.Dispatch<React.SetStateAction<boolean>>,
  dataParam: UserProfileResponse,
  setData: React.Dispatch<React.SetStateAction<UserProfileResponse | undefined>>
) {
  return useMutation({
    mutationFn: verifyEmail,

    onMutate: () => {
      const toastId = toast.loading("Verifying otp...");
      return { toastId };
    },

    onSuccess: (data, variables, context) => {
      setIsModalOpen(false);
      setData({
        ...dataParam,
        user: {
          ...dataParam.user,
          is_verified: true,
        },
      });
      toast.success("Email has verified successfully", {
        id: context?.toastId,
        duration: 3000,
      });
    },

    onError: (error, variables, context) => {
      toast.error("Failed to verify email", {
        id: context?.toastId,
        duration: 3000,
      });
      console.log(error);
    },
  });
}

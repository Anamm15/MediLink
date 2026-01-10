"use client";

import { Button } from "@/components/ui/Button";
import { Input } from "@/components/ui/form/Input";
import { Modal } from "@/components/ui/Modal";
import { useState } from "react";
import { useVerifyEmail } from "../hooks/useUser";
import { TypographyH4 } from "@/components/ui/Typography";
import { UserProfileResponse } from "@/types/user.type";

type EmailVerificationModalProps = {
  isOpen: boolean;
  setIsOpen: React.Dispatch<React.SetStateAction<boolean>>;
  data: UserProfileResponse;
  setData: React.Dispatch<
    React.SetStateAction<UserProfileResponse | undefined>
  >;
};

export default function EmailVerificationModal({
  isOpen,
  setIsOpen,
  data,
  setData,
}: EmailVerificationModalProps) {
  const [OTP, setOTP] = useState("");
  const { mutate: verifyEmail } = useVerifyEmail(setIsOpen, data, setData);

  const handleSubmit = () => {
    verifyEmail({ otp: OTP });
  };

  return (
    <Modal setIsOpen={setIsOpen} open={isOpen} className="w-96">
      <TypographyH4 className="text-center mb-5">
        Verify Your Email
      </TypographyH4>
      <Input
        label="OTP"
        placeholder="Enter OTP sent to your email"
        required
        value={OTP}
        onChange={(e) => setOTP(e.target.value)}
      />
      <Button onClick={handleSubmit} className="w-full mt-4">
        Submit
      </Button>
    </Modal>
  );
}

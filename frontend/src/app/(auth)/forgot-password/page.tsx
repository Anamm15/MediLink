"use client";

"use client";

import { useState } from "react";
import { Eye, EyeOff } from "lucide-react";
import { Input } from "@/components/ui/form/Input";
import { Button } from "@/components/ui/Button";
import { ResetPasswordRequest } from "@/types/auth.type";
import {
  useRequestResetPassword,
  useResetPassword,
} from "./hooks/useForgotPassword";

export default function RegisterPage() {
  const [data, setData] = useState<ResetPasswordRequest>({
    email: "",
    new_password: "",
    otp: "",
  });
  const [showPassword, setShowPassword] = useState(false);
  const { mutate: requestResetPassword } = useRequestResetPassword();
  const { mutate: resetPassword } = useResetPassword();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    resetPassword(data);
  };

  const handleRequestResetPassword = async (e: React.FormEvent) => {
    e.preventDefault();
    requestResetPassword(data.email);
  };

  return (
    <main>
      <form onSubmit={handleSubmit} className="space-y-5">
        <div className="relative">
          <Input
            label="Email"
            type="email"
            placeholder="yourmail@email.com"
            required
            value={data.email}
            onChange={(e) => setData({ ...data, email: e.target.value })}
          />
          <button
            onClick={handleRequestResetPassword}
            className="absolute right-3 top-1/2 text-sm text-cyan-600 font-bold cursor-pointer hover:text-cyan-200"
          >
            Send OTP
          </button>
        </div>

        <Input
          label="OTP"
          type="text"
          placeholder="+628123456789"
          required
          value={data.otp}
          onChange={(e) => setData({ ...data, otp: e.target.value })}
        />

        <div className="relative">
          <Input
            label="Password"
            type={showPassword ? "text" : "password"}
            placeholder="••••••••"
            required
            value={data.new_password}
            onChange={(e) => setData({ ...data, new_password: e.target.value })}
          />
          <button
            type="button"
            onClick={() => setShowPassword(!showPassword)}
            className="absolute right-3 top-1/2 text-gray-400 hover:text-cyan-600 transition-colors"
          >
            {!showPassword ? (
              <EyeOff className="h-5 w-5" />
            ) : (
              <Eye className="h-5 w-5" />
            )}
          </button>
        </div>

        <Button
          type="submit"
          className="w-full py-6 text-lg font-bold shadow-lg shadow-cyan-200"
        >
          Change Password
        </Button>
      </form>
    </main>
  );
}

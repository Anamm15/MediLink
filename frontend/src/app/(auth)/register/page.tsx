"use client";

import { useState } from "react";
import Link from "next/link";
import { Eye, EyeOff } from "lucide-react";
import { Input } from "@/components/ui/form/Input";
import { Button } from "@/components/ui/Button";
import { toast } from "sonner";
import { useRegister } from "./hooks/useRegister";
import { RegistrationRequest } from "@/types/auth.type";
import { TypographySmall } from "@/components/ui/Typography";

export default function RegisterPage() {
  const [data, setData] = useState<RegistrationRequest>({
    name: "",
    email: "",
    password: "",
    phone_number: "",
  });
  const [confirmPassword, setConfirmPassword] = useState("");
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const { mutate } = useRegister();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (data.password !== confirmPassword) {
      toast.message("Password and confirm password do not match");
      return;
    }

    mutate(data);
  };

  return (
    <main>
      <form onSubmit={handleSubmit} className="space-y-5">
        <Input
          label="Name"
          type="text"
          placeholder="John Doe"
          required
          value={data.name}
          onChange={(e) => setData({ ...data, name: e.target.value })}
        />

        <Input
          label="Email"
          type="email"
          placeholder="yourmail@email.com"
          required
          value={data.email}
          onChange={(e) => setData({ ...data, email: e.target.value })}
        />

        <Input
          label="Phone Number"
          type="text"
          placeholder="+628123456789"
          required
          value={data.phone_number}
          onChange={(e) => setData({ ...data, phone_number: e.target.value })}
        />

        <div className="relative">
          <Input
            label="Password"
            type={showPassword ? "text" : "password"}
            placeholder="••••••••"
            required
            value={data.password}
            onChange={(e) => setData({ ...data, password: e.target.value })}
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

        <div className="relative">
          <Input
            label="Confirm Password"
            type={showConfirmPassword ? "text" : "password"}
            placeholder="••••••••"
            required
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
          />
          <button
            type="button"
            onClick={() => setShowConfirmPassword(!showConfirmPassword)}
            className="absolute right-3 top-1/2 text-gray-400 hover:text-cyan-600 transition-colors"
          >
            {!showConfirmPassword ? (
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
          Register
        </Button>
      </form>

      {/* Divider */}
      <div className="relative my-5">
        <div className="absolute inset-0 flex items-center">
          <span className="w-full border-t border-gray-200"></span>
        </div>
        <div className="relative flex justify-center text-xs uppercase">
          <span className="bg-white px-2 text-gray-400">Or Sign Up With</span>
        </div>
      </div>

      {/* Social Login (Opsional - Menambah Kesan Premium) */}
      <div className="grid grid-cols-2 gap-4">
        <button className="flex items-center justify-center gap-2 py-2.5 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors text-sm font-semibold text-gray-700">
          <img
            src="https://www.svgrepo.com/show/355037/google.svg"
            className="h-4 w-4"
            alt="Google"
          />
          Google
        </button>
        <button className="flex items-center justify-center gap-2 py-2.5 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors text-sm font-semibold text-gray-700">
          <img
            src="https://www.svgrepo.com/show/452196/facebook-1.svg"
            className="h-5 w-5"
            alt="Facebook"
          />
          Facebook
        </button>
      </div>

      {/* Footer Card */}
      <TypographySmall className="mt-4 text-center text-gray-500">
        Already have an account?{" "}
        <Link href="/login" className="text-cyan-600 font-bold hover:underline">
          Sign In Here
        </Link>
      </TypographySmall>
    </main>
  );
}

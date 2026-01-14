"use client";

import { useEffect, useState } from "react";
import Link from "next/link";
import { Eye, EyeOff } from "lucide-react";
import { Input } from "@/components/ui/form/Input";
import { Button } from "@/components/ui/Button";
import { useLogin } from "./hooks/useLogin";
import { LoginRequest } from "@/types/auth.type";
import { TypographySmall } from "@/components/ui/Typography";

export default function LoginPage() {
  const [data, setData] = useState<LoginRequest>({
    email: "",
    password: "",
  });
  const [showPassword, setShowPassword] = useState(false);
  const { mutate: login } = useLogin();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    login(data);
  };

  return (
    <main>
      <form onSubmit={handleSubmit} className="space-y-5">
        <Input
          label="Email"
          type="email"
          placeholder="yourmail@email.com"
          required
          value={data.email}
          onChange={(e) => setData({ ...data, email: e.target.value })}
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

        <div className="flex items-center justify-between text-sm">
          <label className="flex items-center gap-2 cursor-pointer text-gray-600">
            <input
              type="checkbox"
              className="rounded border-gray-300 text-cyan-600 focus:ring-cyan-500"
            />
            Remember Me
          </label>
          <Link
            href="/forgot-password"
            className="text-cyan-600 font-semibold hover:underline"
          >
            Forgot Password?
          </Link>
        </div>

        <Button
          type="submit"
          className="w-full py-6 text-lg font-bold shadow-lg shadow-cyan-200"
        >
          Sign In Now
        </Button>
      </form>

      {/* Divider */}
      <div className="relative my-5">
        <div className="absolute inset-0 flex items-center">
          <span className="w-full border-t border-gray-200"></span>
        </div>
        <div className="relative flex justify-center text-xs uppercase">
          <span className="bg-white px-2 text-gray-400">Or Sign In With</span>
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
        Have no account yet?{" "}
        <Link
          href="/register"
          className="text-cyan-600 font-bold hover:underline"
        >
          Sign Up Here
        </Link>
      </TypographySmall>
    </main>
  );
}

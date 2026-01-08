import * as React from "react";
import { cn } from "@/lib/utils";

type ErrorTextProps = {
  error?: string | boolean;
  helperText?: string;
};

export function ErrorText({ error, helperText }: ErrorTextProps) {
  return (
    <div>
      {error && typeof error === "string" ? (
        <p className="mt-1 text-xs text-red-500 animate-pulse">{error}</p>
      ) : helperText ? (
        <p className="mt-1 text-xs text-slate-500">{helperText}</p>
      ) : null}
    </div>
  );
}

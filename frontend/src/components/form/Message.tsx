// components/ErrorMessage.tsx
"use client";

import React from "react";

interface ErrorMessageProps {
  text: string;
  type?: "error" | "success";
}

export function Message({ text, type = "error" }: ErrorMessageProps) {
  return (
    <div
      className={`text-sm mt-2.5 !leading-tight ${
        type === "error" ? "text-red-600" : "text-green-600"
      }`}
    >
      {text}
    </div>
  );
};

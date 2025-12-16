"use client";

import React from "react";

interface LabelTextProps {
  id: string | undefined;
  children: React.ReactNode;
  className?: string;
}

const LabelText: React.FC<LabelTextProps> = ({ id, children, className }) => {
  return (
    <label
      htmlFor={id ? id : ""}
      className={`text-gray-700 text-sm ${className}`}
    >
      {children}
    </label>
  );
};

export default LabelText;

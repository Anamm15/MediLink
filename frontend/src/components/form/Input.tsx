"use client";

import React from "react";
import LabelText from "./LabelText";

interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label?: string;
  className?: string;
  error?: string;
  icon?: React.ReactNode;
}

const Input: React.FC<InputProps> = ({
  label,
  type = "text",
  name,
  readOnly,
  disabled,
  className = "",
  placeholder,
  error = "",
  required = false,
  icon,
  ...rest
}) => {
  return (
    <div className="flex flex-col space-y-2">
      <LabelText id={label} className="flex items-center gap-2">
        {icon && <span>{icon}</span>}
        <span>
          {label}
          {required && <span className="text-red-500 ms-1">*</span>}
        </span>
      </LabelText>
      <input
        id={label}
        name={name}
        type={type}
        readOnly={readOnly}
        disabled={disabled}
        placeholder={placeholder}
        className={`px-3 py-2 border border-[#808080] rounded-md text-sm md:text-md
          focus:outline-1 focus:outline-primary-info-active focus:ring-inset 
          hover:ring-1 hover:ring-inset hover:ring-[#000] 
          placeholder:text-sm placeholder:text-[#9AA2B1] focus:placeholder:text-[#092540] 
          ${className}`}
        {...rest}
      />
      {error && <span className="text-red-500 text-xs">{error}</span>}
    </div>
  );
};

export default Input;

"use client";

import React from "react";
import LabelText from "./LabelText";

type SelectOptionProps = {
  label: string;
  helperText?: string;
  hideError?: boolean;
  validation?: boolean;
  className?: string;
  readOnly?: boolean;
  error?: string;
  required?: boolean;
  placeholder?: string;
  onChange?: (e: React.ChangeEvent<HTMLSelectElement>) => void;
  icon?: React.ReactNode;
  value?: string | number | undefined;
  options: { value: string | number | undefined; label: string }[];
};

const SelectOption = ({
  label,
  className = "",
  readOnly = false,
  placeholder = "",
  icon,
  error = "",
  onChange,
  options,
  value,
  required = false,
  ...rest
}: SelectOptionProps) => {
  return (
    <div className="flex flex-col space-y-2">
      <LabelText id={label} className="flex items-center gap-2">
        {icon && <span>{icon}</span>}
        <span>
          {label}
          {required && <span className="text-red-500 ms-1">*</span>}
        </span>
      </LabelText>
      <div className="relative">
        <select
          id={label}
          name={label}
          disabled={readOnly}
          className={`appearance-none w-full px-3 py-2 border border-[#808080] rounded-md 
            focus:outline-1 focus:outline-primary-info-active focus:ring-inset 
            hover:ring-1 hover:ring-inset hover:ring-[#000] text-sm md:text-md
            placeholder:text-sm placeholder:text-[#9AA2B1] focus:placeholder:text-[#092540] 
            pr-10 ${className}`}
          aria-label={label}
          value={value}
          onChange={onChange}
          {...rest}
        >
          {placeholder && (
            <option value="" disabled hidden>
              {placeholder}
            </option>
          )}
          {options.map((option) => (
            <option key={String(option.value)} value={String(option.value)}>
              {String(option.label)}
            </option>
          ))}
        </select>

        <span className="pointer-events-none absolute inset-y-0 right-3 flex items-center z-10">
          <svg
            className="w-5 h-5 text-gray-600"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M19 9l-7 7-7-7"
            />
          </svg>
        </span>
      </div>
      {error && <span className="text-red-500 text-xs">{error}</span>}
    </div>
  );
};

export default SelectOption;

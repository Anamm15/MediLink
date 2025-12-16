"use client";

import React from "react";
import LabelText from "./LabelText";

interface CheckboxProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label: string;
  className?: string;
}

const Checkbox: React.FC<CheckboxProps> = ({
  label,
  id,
  name,
  readOnly,
  disabled,
  className = "",
  placeholder,
  defaultChecked,
  onChange,
  ...rest
}) => {
  return (
    <div className="mb-4 flex items-center space-x-2">
      <label className="relative flex items-center cursor-pointer select-none">
        <input
          type="checkbox"
          id={id}
          name={name}
          readOnly={readOnly}
          disabled={disabled}
          placeholder={placeholder}
          defaultChecked={defaultChecked}
          onChange={onChange}
          className="peer hidden"
          {...rest}
        />

        <span
          className={`h-5 w-5 rounded-md border border-gray-400 flex items-center justify-center 
            transition-all duration-200 peer-checked:bg-blue-600 peer-checked:border-blue-600
            peer-disabled:bg-gray-200 peer-disabled:border-gray-300
            ${className}`}
        >
          <svg
            className="h-3 w-3 text-white opacity-0 scale-75 transition-all duration-200 peer-checked:opacity-100 peer-checked:scale-100"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            strokeWidth="3"
          >
            <path strokeLinecap="round" strokeLinejoin="round" d="M5 13l4 4L19 7" />
          </svg>
        </span>

        <span className="ml-2">
          <LabelText id={id || name || ""}>{label}</LabelText>
        </span>
      </label>
    </div>
  );
};

export default Checkbox;

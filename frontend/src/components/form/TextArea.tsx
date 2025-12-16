"use client";

import React from "react";
import LabelText from "./LabelText";

interface TextAreaProps
  extends React.TextareaHTMLAttributes<HTMLTextAreaElement> {
  label: string;
  className?: string;
  error?: string;
}

const TextArea: React.FC<TextAreaProps> = ({
  label,
  name,
  readOnly,
  disabled,
  className = "",
  placeholder,
  required = false,
  error = "",
  ...rest
}) => {
  return (
    <div className="mb-4 flex flex-col space-y-2">
      <LabelText id={label}>
        {label}
        {required && <span className="text-red-500 ms-1">*</span>}
      </LabelText>
      <textarea
        id={label}
        name={name}
        readOnly={readOnly}
        disabled={disabled}
        placeholder={placeholder}
        className={`px-3 py-2 border border-[#808080] rounded-md
          focus:outline-1 focus:outline-primary-info-active focus:ring-inset 
          hover:ring-1 hover:ring-inset hover:ring-[#000] resize-none
          placeholder:text-sm min-h-32
          placeholder:text-[#9AA2B1] focus:placeholder:text-[#092540] 
          ${className}`}
        {...rest}
      />
      {error && <span className="text-red-500 text-xs">{error}</span>}
    </div>
  );
};

export default TextArea;

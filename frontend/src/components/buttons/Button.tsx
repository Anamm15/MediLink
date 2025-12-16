"use client";

import React from "react";

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  className?: string;
  children: React.ReactNode;
}

const Button: React.FC<ButtonProps> = ({
  className = "",
  onClick,
  children,
  ...rest
}) => {
  return (
    <button
      className={`shadow-md px-4 py-2 cursor-pointer bg-accent rounded-lg font-semibold text-black transition-all duration-300
        hover:bg-accent-hover ${className}`}
      onClick={onClick}
      type="button"
      {...rest}
    >
      {children}
    </button>
  );
};

export default Button;

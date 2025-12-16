"use client";

import React from "react";
import { CiCircleChevDown } from "react-icons/ci";

interface ArrowButtonProps {
  className?: string;
  id: string;
}

const ArrowButton: React.FC<ArrowButtonProps> = ({ className = "", id }) => {
  const handleClick = () => {
    const element = document.getElementById(id);
    if (element) {
      let offset: number;
      if (window.innerWidth < 1024) {
        offset = element.getBoundingClientRect().top + window.scrollY - 64;
      } else {
        offset = element.getBoundingClientRect().top + window.scrollY - 96;
      }
      window.scrollTo({
        top: offset,
        behavior: "smooth",
      });
    }
  };

  return (
    <div className="mt-6 flex items-center justify-center">
      <p className={`z-10 max-lg:text-[32px] ${className}`}>
        <button onClick={handleClick} type="button">
          <CiCircleChevDown className="animate-bounce" />
        </button>
      </p>
    </div>
  );
};

export default ArrowButton;

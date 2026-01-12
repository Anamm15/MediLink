import React from "react";

interface InfoSectionProps {
  title: string;
  children: React.ReactNode;
}

export const InfoSection = ({ title, children }: InfoSectionProps) => {
  return (
    <div className="bg-white p-6 rounded-xl border border-gray-200 shadow-sm">
      <h2 className="text-xl font-bold text-gray-800 pb-3 border-b border-gray-200 mb-4">
        {title}
      </h2>
      <div className="prose prose-sm max-w-none text-gray-600">{children}</div>
    </div>
  );
};

import React from "react";

interface SectionCardProps {
  title: string;
  children: React.ReactNode;
  action?: React.ReactNode; // Tombol aksi opsional seperti "Edit"
}

export const SectionCard = ({ title, children, action }: SectionCardProps) => {
  return (
    <div className="bg-white rounded-xl shadow-sm border border-gray-200">
      <div className="flex items-center justify-between p-4 border-b">
        <h3 className="text-lg font-semibold text-gray-800">{title}</h3>
        {action}
      </div>
      <div className="p-4">{children}</div>
    </div>
  );
};

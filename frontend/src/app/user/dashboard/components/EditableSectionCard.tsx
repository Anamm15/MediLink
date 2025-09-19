import React from "react";
import { Edit } from "lucide-react";

interface CardProps {
  title: string;
  children: React.ReactNode;
}

export const EditableSectionCard = ({ title, children }: CardProps) => {
  return (
    <div className="bg-white rounded-xl shadow-sm border border-gray-200">
      <div className="flex items-center justify-between p-4 border-b">
        <h3 className="text-lg font-semibold text-gray-800">{title}</h3>
        <button className="flex items-center gap-2 px-3 py-1.5 text-xs font-semibold text-gray-600 bg-gray-100 rounded-md hover:bg-gray-200">
          <Edit className="w-3 h-3" /> Edit
        </button>
      </div>
      <div className="p-6">{children}</div>
    </div>
  );
};

import { Input } from "@/components/ui/form/Input";
import React from "react";

type EditableFieldProps = {
  value: string | number | null | undefined;
  isEditing: boolean;
  onChange: (value: string) => void;
  placeholder?: string;
};

export default function EditableField({
  value,
  isEditing,
  onChange,
  placeholder,
}: EditableFieldProps) {
  return isEditing ? (
    <Input
      type="text"
      placeholder={placeholder}
      value={value ?? ""}
      onChange={(e) => onChange(e.target.value)}
    />
  ) : (
    <p className="text-gray-800 font-semibold">{value ? value : "-"}</p>
  );
}

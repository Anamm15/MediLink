import { Search } from "lucide-react";
import { Input } from "./form/Input";
import { Select } from "./form/Select";
import React, { useState } from "react";

type SearchFilterProps = {
  setDoctorNameFiltered: React.Dispatch<React.SetStateAction<string>>;
};

export const SearchFilter = ({ setDoctorNameFiltered }: SearchFilterProps) => {
  const [name, setName] = useState("");

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault();
    setDoctorNameFiltered(name);
  };

  const options = [
    {
      value: "general",
      label: "Dokter Umum",
    },
    {
      value: "dentist",
      label: "Dokter Gigi",
    },
    {
      value: "pediatrician",
      label: "Dokter Anak",
    },
  ];

  return (
    <div className="bg-white p-6 rounded-xl shadow-md mb-8">
      <form
        onSubmit={handleSearch}
        className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-4 items-end"
      >
        {/* Search by Name */}
        <div className="md:col-span-3 lg:col-span-2">
          <Input
            id="doctor-name"
            type="text"
            label="Docter Name"
            placeholder="Contoh: Dr. Budi Santoso"
            startIcon={<Search className="h-5 w-5 text-gray-400" />}
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
        </div>

        {/* Filter by Specialty */}
        <Select label="Specialty" className="">
          {options.map((option) => (
            <option key={option.value} value={option.value}>
              {option.label}
            </option>
          ))}
        </Select>

        {/* Search Button */}
        <div className="mt-4 md:mt-0">
          <button
            type="submit"
            className="w-full bg-slate-800 text-white font-semibold py-2 px-6 rounded-lg hover:bg-slate-700 transition-colors duration-200 flex items-center justify-center"
          >
            <Search className="w-4 h-4 mr-2" />
            Cari
          </button>
        </div>
      </form>
    </div>
  );
};

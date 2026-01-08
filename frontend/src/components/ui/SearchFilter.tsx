import { Search } from "lucide-react";
import { Input } from "./form/Input";
import { Select } from "./form/Select";

export const SearchFilter = () => {
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
      <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-4 items-end">
        {/* Search by Name */}
        <div className="md:col-span-3 lg:col-span-2">
          <Input
            id="doctor-name"
            type="text"
            placeholder="Contoh: Dr. Budi Santoso"
            // icon={<Search className="h-5 w-5 text-gray-400" />}
          />
        </div>

        {/* Filter by Specialty */}
        <Select label="Spesialisasi">
          {options.map((option) => (
            <option key={option.value} value={option.value}>
              {option.label}
            </option>
          ))}
        </Select>

        {/* Search Button */}
        <div className="mt-4 md:mt-0">
          <button className="w-full bg-slate-800 text-white font-semibold py-2 px-6 rounded-lg hover:bg-slate-700 transition-colors duration-200 flex items-center justify-center">
            <Search className="w-4 h-4 mr-2" />
            Cari
          </button>
        </div>
      </div>
    </div>
  );
};

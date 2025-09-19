import { Search, MapPin } from "lucide-react";

export const SearchFilter = () => {
  return (
    <div className="bg-white p-6 rounded-xl shadow-md mb-8">
      <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-4 items-end">
        {/* Search by Name */}
        <div className="md:col-span-3 lg:col-span-2">
          <label
            htmlFor="doctor-name"
            className="block text-sm font-medium text-gray-700 mb-1"
          >
            Cari Nama Dokter
          </label>
          <div className="relative">
            <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <Search className="h-5 w-5 text-gray-400" />
            </div>
            <input
              type="text"
              id="doctor-name"
              className="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-cyan-500 focus:border-cyan-500"
              placeholder="Contoh: Dr. Budi Santoso"
            />
          </div>
        </div>

        {/* Filter by Specialty */}
        <div>
          <label
            htmlFor="specialty"
            className="block text-sm font-medium text-gray-700 mb-1"
          >
            Spesialisasi
          </label>
          <select
            id="specialty"
            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-cyan-500 focus:border-cyan-500"
          >
            <option>Semua</option>
            <option>Dokter Umum</option>
            <option>Dokter Gigi</option>
            <option>Dokter Anak</option>
          </select>
        </div>

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

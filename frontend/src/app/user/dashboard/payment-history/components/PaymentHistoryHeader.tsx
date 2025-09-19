import { Search } from "lucide-react";

export const PaymentHistoryHeader = () => {
  return (
    <div className="bg-white p-4 rounded-xl shadow-sm border border-gray-200">
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        {/* Ringkasan */}
        <div className="p-4 bg-slate-50 rounded-lg">
          <p className="text-sm text-gray-500">Total Pengeluaran Bulan Ini</p>
          <p className="text-2xl font-bold text-gray-800">Rp 475.000</p>
        </div>

        {/* Pencarian */}
        <div className="relative md:col-span-2">
          <input
            type="text"
            placeholder="Cari berdasarkan deskripsi..."
            className="w-full h-full p-3 pl-10 border rounded-lg"
          />
          <Search className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-gray-400" />
        </div>

        {/* Filter (bisa ditambahkan di sini) */}
      </div>
    </div>
  );
};

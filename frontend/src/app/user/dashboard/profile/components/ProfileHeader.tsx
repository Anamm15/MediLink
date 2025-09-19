import Image from "next/image";
import { KeyRound, Camera } from "lucide-react";

// Terima data user sebagai props
export const ProfileHeader = ({ userData }: { userData: any }) => {
  return (
    <div className="bg-white p-6 rounded-xl shadow-sm border border-gray-200 flex flex-col sm:flex-row items-center gap-6">
      <div className="relative">
        <Image
          src={userData.avatarUrl}
          alt={userData.name}
          width={96}
          height={96}
          className="rounded-full"
        />
        <button className="absolute bottom-0 right-0 p-1.5 bg-gray-200 rounded-full hover:bg-gray-300">
          <Camera className="w-4 h-4 text-gray-700" />
        </button>
      </div>
      <div className="flex-grow text-center sm:text-left">
        <h2 className="text-2xl font-bold text-gray-800">{userData.name}</h2>
        <p className="text-gray-500">{userData.email}</p>
      </div>
      <div className="flex-shrink-0">
        <button className="flex items-center gap-2 px-4 py-2 text-sm font-semibold text-gray-700 bg-gray-100 rounded-lg hover:bg-gray-200">
          <KeyRound className="w-4 h-4" /> Ubah Password
        </button>
      </div>
    </div>
  );
};
